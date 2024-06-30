import json
import os
import shlex
import base64
import subprocess
from flask import Flask, request, jsonify
import time
import argparse
import ipfshttpclient
import logging

app = Flask(__name__)

HEADERS = {
    "content-type": "application/json",
    "access-control-allow-origin": "*",
    "access-control-allow-methods": "OPTIONS, POST",
}

runtime_version = "1.0.0"

# Setup logging
logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')

# IPFS Client
class IPFSClient:
    def __init__(self):
        self._hash = None
        self._client = ipfshttpclient.connect(session=True)
        logging.debug("IPFS client initialized")

    def upload(self, content):
        self._hash = self._client.add_json(content)
        logging.debug(f"Uploaded content to IPFS, hash: {self._hash}")

    def get_hash(self):
        return self._hash

    def close(self):  # Call this when you're done
        self._client.close()
        logging.debug("IPFS client connection closed")

def success(returncode, stdout, stderr, err, duration, installation_time, execution_time):
    result = {
        "returncode": returncode,
        "stdout": stdout,
        "stderr": stderr,
        "error": err,
        "duration": str(duration) + "s",
        "version": runtime_version,
        "installation-time": str(installation_time),
        "execution-time": str(execution_time),
    }
    logging.debug(f"Success result: {result}")
    return result

def bad_request(err):
    logging.error(f"Bad request: {err}")
    return {
        "statusCode": 400,
        "headers": HEADERS,
        "body": json.dumps({"error": err}),
    }

@app.route('/execute', methods=['POST'])
def lambda_handler_flask():
    try:
        body = request.json
#         logging.debug(f"Request body: {body}")
    except (json.decoder.JSONDecodeError, KeyError) as e:
        logging.error("Cannot be jsonified")
        return bad_request("Cannot be jsonified")

    env = os.environ.copy()

    MAX_EXECUTABLE = 10000
    MAX_DATA_SIZE = 10000

    if "used-external-libraries" not in body:
        return bad_request("Missing used external libraries")
    used_external_libraries = body["used-external-libraries"]

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
    calldata = body["calldata"]

    if "timeout" not in body:
        return bad_request("Missing timeout value")
    try:
        timeout = int(body["timeout"])
    except ValueError:
        return bad_request("Timeout format invalid")

    port = request.environ.get('SERVER_PORT')
    logging.debug(f"Server port: {port}")

    start_time = time.time()  # Capture the start time

    venv_path = ""
    # Create a virtual environment
    if used_external_libraries == "no":
        venv_path = f"/tmp/default-venv{port}"
    else:
        venv_path = f"/tmp/venv{port}"
        subprocess.run(["python3", "-m", "venv", venv_path], check=True)
    logging.debug(f"Virtual environment path: {venv_path}")

    requirement_path = f"/tmp/requirements{port}.txt"
    with open(requirement_path, "w") as req_f:
        req_f.write(requirement_file.decode())
    logging.debug(f"Requirement path: {requirement_path}")

    if used_external_libraries == "yes":
        pip_path = f"{venv_path}/bin/pip3"
        installation_cmd = f"{pip_path} install -r {requirement_path}"
        subprocess.run(installation_cmd, shell=True)
        logging.debug(f"Installation command: {installation_cmd}")
    end_time = time.time()
    installation_time = end_time - start_time
    logging.debug(f"Installation time: {installation_time}")

    start_time = time.time()
    path = f"/tmp/execute{port}.py"
    with open(path, "w") as f:
        f.write(executable.decode())
    logging.debug(f"Executable path: {path}")

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
#         logging.debug(f"Subprocess started with command: {[f'{venv_path}/bin/python3', path] + shlex.split(body['calldata'])]}")

        proc.wait(timeout=timeout)
        execution_time = time.time() - start_time  # Calculate the execution time
        returncode = proc.returncode
        stdout = proc.stdout.read(MAX_DATA_SIZE).decode()
        stderr = proc.stderr.read(MAX_DATA_SIZE).decode()
        logging.debug(f"Subprocess return code: {returncode}")
        logging.debug(f"Subprocess stdout: {stdout}")
        logging.debug(f"Subprocess stderr: {stderr}")

        ipfs_cli = IPFSClient()
        ipfs_cli.upload(stdout)
        content_hash = ipfs_cli.get_hash()
        ipfs_cli.close()
        logging.debug(f"Content hash from IPFS: {content_hash}")

        result = success(returncode, content_hash, stderr, "", installation_time + execution_time, installation_time, execution_time)
    except OSError:
        execution_time = time.time() - start_time  # Calculate the execution time
        logging.error("Execution failed with OSError")
        result = success(126, "", "", "Execution fail", installation_time + execution_time, installation_time, execution_time)
    except subprocess.TimeoutExpired:
        execution_time = time.time() - start_time  # Calculate the execution time
        logging.error("Execution time limit exceeded")
        result = success(111, "", "", "Execution time limit exceeded", installation_time + execution_time, installation_time, execution_time)
    finally:
        if used_external_libraries == "yes":
            clean_cmd = f"rm -rf {venv_path} {requirement_path} {path}"
            subprocess.run(clean_cmd, shell=True)
            logging.debug(f"Cleanup command: {clean_cmd}")
        logging.debug(f"Final result: {str(result)}")
        return result

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Run the Flask application')
    parser.add_argument('port', type=int, help='The port number to run the Flask application on')
    args = parser.parse_args()
    logging.debug(f"Starting Flask application on port: {args.port}")
    app.run(debug=True, port=args.port)
