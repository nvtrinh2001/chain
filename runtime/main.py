import json
import os
import shlex
import base64
import subprocess
from flask import Flask, request, jsonify
import time
import argparse

app = Flask(__name__)

HEADERS = {
    "content-type": "application/json",
    "access-control-allow-origin": "*",
    "access-control-allow-methods": "OPTIONS, POST",
}

runtime_version = "1.0.0"

def success(returncode, stdout, stderr, err, duration):
    return {
        "returncode": returncode,
        "stdout": stdout,
        "stderr": stderr,
        "error": err,
        "duration": str(duration) + "s",
        "version": runtime_version,
    }

def bad_request(err):
    return {
        "statusCode": 400,
        "headers": HEADERS,
        "body": json.dumps({"error": err}),
    }

@app.route('/execute', methods=['POST'])
def lambda_handler_flask():
    try:
        body = request.json
    except (json.decoder.JSONDecodeError, KeyError) as e:
        # Hack for preflight
        return bad_request("Cannot be jsonified")

    env = os.environ.copy()

    MAX_EXECUTABLE = 10000
    MAX_DATA_SIZE = 10000

    if "executable" not in body:
        return bad_request("Missing executable value")
    executable = base64.b64decode(body["executable"])
    if len(executable) > MAX_EXECUTABLE:
        return bad_request("Executable exceeds max size")
    if "calldata" not in body:
        return bad_request("Missing calldata value")
    if len(body["calldata"]) > MAX_DATA_SIZE:
        return bad_request("Calldata exceeds max size")
    if "timeout" not in body:
        return bad_request("Missing timeout value")
    try:
        timeout = int(body["timeout"])
    except ValueError:
        return bad_request("Timeout format invalid")

    port = request.environ.get('SERVER_PORT')

    start_time = time.time()  # Capture the start time

    # Create a virtual environment
    venv_path = f"venv{port}"
    subprocess.run(["python3", "-m", "venv", venv_path], check=True)

    path = f"/tmp/execute{port}.py"
    with open(path, "w") as f:
        f.write(executable.decode())

    result = {}
    try:
        env = os.environ.copy()
        for key, value in body.get("env", {}).items():
            env[key] = value

        proc = subprocess.Popen(
            [f"{venv_path}/bin/python3", path] + shlex.split(body["calldata"]),
            env=env,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )

        proc.wait(timeout=timeout)
        execution_time = time.time() - start_time  # Calculate the execution time
        returncode = proc.returncode
        stdout = proc.stdout.read(MAX_DATA_SIZE).decode()
        stderr = proc.stderr.read(MAX_DATA_SIZE).decode()

        result = success(returncode, stdout, stderr, "", execution_time)
    except OSError:
        execution_time = time.time() - start_time  # Calculate the execution time
        result = success(126, "", "", "Execution fail", execution_time)
    except subprocess.TimeoutExpired:
        execution_time = time.time() - start_time  # Calculate the execution time
        result = success(111, "", "", "Execution time limit exceeded", execution_time)
    finally:
        clean_cmd = f"rm -rf {venv_path} {path}"
        subprocess.run(clean_cmd, shell=True)
        print(result)
        return result

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Run the Flask application')
    parser.add_argument('port', type=int, help='The port number to run the Flask application on')
    args = parser.parse_args()
    app.run(debug=True, port=args.port)

