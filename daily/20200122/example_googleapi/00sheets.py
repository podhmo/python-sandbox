import os
import dotenv
import google_auth_oauthlib
import requests
from google.oauth2.credentials import Credentials
from google.auth.transport.requests import Request, AuthorizedSession


def get_credentials() -> Credentials:
    dotenv.load_dotenv()
    client_id = os.environ["CLIENT_ID"]
    client_secret = os.environ["CLIENT_SECRET"]

    scopes = [
        "https://www.googleapis.com/auth/drive",
    ]

    credentials = google_auth_oauthlib.get_user_credentials(
        scopes, client_id, client_secret
    )
    return credentials


def get_session(credentials: Credentials) -> AuthorizedSession:
    return AuthorizedSession(credentials=credentials)


def iterate_spreadsheets(session: AuthorizedSession):
    page_token = ""
    url = "https://www.googleapis.com/drive/v3/files"

    params = {
        "q": "mimeType='application/vnd.google-apps.spreadsheet'",
        "pageSize": 1000,
        "supportsTeamDrives": True,
        "includeTeamDriveItems": True,
        "fields": "nextPageToken,files(name,id)",
    }

    while page_token is not None:
        if page_token:
            params["pageToken"] = page_token

        res = session.request("get", url, params=params)
        assert res.status_code == 200, res.status_code
        data = res.json()
        for file in data.get("files", []):
            yield file

        page_token = data.get("nextPageToken", None)
        if page_token is None:
            break


credentials = get_credentials()

print(credentials.to_json())
print(credentials.token)
print(credentials.valid)
print(credentials.refresh(Request()))

session = get_session(credentials)
assert isinstance(session, requests.Session), session.__class__.mro()

for file in iterate_spreadsheets(session):
    print(f'Found file: {file.get("name")} ({file.get("id")})')
