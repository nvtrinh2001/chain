import json
import os
import shlex
import base64
import subprocess
from flask import Flask, request, jsonify

app = Flask(__name__)

HEADERS = {
    "content-type": "application/json",
    "access-control-allow-origin": "*",
    "access-control-allow-methods": "OPTIONS, POST",
}

runtime_version = "1.0.0"

def success(returncode, stdout, stderr, err):
    return {
        "returncode": returncode,
        "stdout": stdout,
        "stderr": stderr,
        "error": err,
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
        return jsonify({
            "statusCode": 200,
            "headers": HEADERS
        })

    env = os.environ.copy()

    MAX_EXECUTABLE = 10000
    MAX_DATA_SIZE = 10000

    if "requirement-file" not in body:
        return bad_request("Missing requirement file")
    requirement_file = base64.b64decode(body["requirement-file"])
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

    # Create a virtual environment
    venv_path = "/tmp/venv"
    subprocess.run(["python3", "-m", "venv", venv_path], check=True)
    
    requirement_path = "/tmp/requirements.txt"
    with open(requirement_path, "w") as req_f:
        req_f.write(requirement_file.decode())

    pip_path = "/tmp/venv/bin/pip3"
    installation_cmd = f"{pip_path} install -r {requirement_path}"
    subprocess.run(installation_cmd, shell=True)

    path = "/tmp/execute.py"
    with open(path, "w") as f:
        f.write(executable.decode())

    result = {}
    try:
        env = os.environ.copy()
        for key, value in body.get("env", {}).items():
            env[key] = value

        proc = subprocess.Popen(
            ["/tmp/venv/bin/python3", path] + shlex.split(body["calldata"]),
            env=env,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )

        proc.wait(timeout=timeout)
        returncode = proc.returncode
        stdout = proc.stdout.read(MAX_DATA_SIZE).decode()
        stderr = proc.stderr.read(MAX_DATA_SIZE).decode()
        result = success(returncode, stdout, stderr, "")
    except OSError:
        result = success(126, "", "", "Execution fail")
    except subprocess.TimeoutExpired:
        result = success(111, "", "", "Execution time limit exceeded")

    clean_cmd = "rm -rf /tmp/venv /tmp/requirements.txt /tmp/execute.py"
    subprocess.run(clean_cmd, shell=True)
    print(result)
    return result

if __name__ == '__main__':
    app.run(debug=True)

