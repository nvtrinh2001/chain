import json
import os
import shlex
import base64
import subprocess
from flask import Flask, request, jsonify
from threading import Thread

app = Flask(__name__)

HEADERS = {
    "content-type": "application/json",
    "x-lambda": "true",
    "access-control-allow-origin": "*",
    "access-control-allow-methods": "OPTIONS, POST", }

runtime_version = "1.0.0"


def success(returncode, stdout, stderr, err):
    return {
        "returncode": returncode,
        "stdout": stdout,
        "stderr": stderr,
        "error": err,
        "version": runtime_version
    }


def bad_request(err):
    return {"error": err}


@app.route("/execute", methods=['POST'])
def execute():
    try:
        body = request.get_json()
    except (json.decoder.JSONDecodeError, KeyError) as e:
        # Hack for preflight
        return {
            "statusCode": 200,
            "headers": HEADERS,
        }

    MAX_EXECUTABLE = 8192
    MAX_DATA_SIZE = 512

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

    path = "/tmp/execute.sh"
    with open(path, "w") as f:
        f.write(executable.decode())

    os.chmod(path, 0o775)
    try:
        env = os.environ.copy()
        env["PYTHONPATH"] = os.getcwd()
        for key, value in body.get("env", {}).items():
            env[key] = value

        proc = subprocess.Popen(
            [path] + shlex.split(body["calldata"]),
            env=env,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            )

        proc.wait(timeout=(timeout / 1000))
        returncode = proc.returncode
        stdout = proc.stdout.read(MAX_DATA_SIZE).strip().decode()
        stderr = proc.stderr.read(MAX_DATA_SIZE).decode()
        success_msg = success(returncode, stdout, stderr, "")
        print(returncode, stdout, stderr)
        print(success_msg)
        app.logger.info(success_msg)
        return success_msg
    except OSError:
        return success(126, "", "", "Execution fail")
    except subprocess.TimeoutExpired:
        return success(111, "", "", "Execution time limit exceeded")


if __name__ == '__main__':
    # flask_thread = Thread(target=app.run, kwargs={"host": '0.0.0.0', "port": 7070})
    
    # Start the Flask app thread
    # flask_thread.start()

    app.run(host='0.0.0.0', port=7070, debug=True)

