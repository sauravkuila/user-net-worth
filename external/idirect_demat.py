import http.client
import json
from datetime import datetime
import hashlib
import sys

def generate_checksum(payload, secret_key):
    time_stamp = datetime.utcnow().isoformat()[:19] + '.000Z'
    data_to_hash = time_stamp + payload + secret_key
    checksum = hashlib.sha256(data_to_hash.encode("utf-8")).hexdigest()
    return checksum

def main():
    if len(sys.argv) != 4:
        print("Usage: python basereun.py <appkey> <secretkey> <sessiontoken>")
        sys.exit(1)

    appkey, secret_key, session_token = sys.argv[1], sys.argv[2], sys.argv[3]

    # conn = http.client.HTTPConnection("sk123.free.beeceptor.com")
    conn = http.client.HTTPSConnection("api.icicidirect.com")
    payload = json.dumps({})

    checksum = generate_checksum(payload, secret_key)

    headers = {
        'Content-Type': 'application/json',
        'X-Checksum': 'token ' + checksum,
        'X-Timestamp': datetime.utcnow().isoformat()[:19] + '.000Z',
        'X-AppKey': appkey,
        'X-SessionToken': session_token
    }

    conn.request("GET", "/breezeapi/api/v1/dematholdings", payload, headers)
    res = conn.getresponse()
    data = res.read()
    print(data.decode("utf-8"))

if __name__ == "__main__":
    main()
