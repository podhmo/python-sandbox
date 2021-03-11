import os
import json
import urllib.request as r

token = os.environ.get("SLACK_API_TOKEN")
assert token, token
payload = {"channel": "programmer-log", "text": "hello"}
data = json.dumps(payload).encode("utf-8")

url = "https://slack.com/api/chat.postMessage"
method = "POST"
headers = {"Content-Type": "application/json", "Authorization": f"Bearer {token}"}

req = r.Request(url, method=method, data=data, headers=headers)
with r.urlopen(req) as res:
    print(res.read().decode("utf-8"))
