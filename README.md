# Open Source Pledge Potential Member List

This program runs a web server which serves information from the private “Open Source Pledge Potential Member List”
spreadsheet.

## Setup

To run this program, you need a Google account that has access to the spreadsheet.

1. Create a project in the [Google Cloud Console](https://console.cloud.google.com/).
2. Go to [“Enables APIs & services”](https://console.cloud.google.com/apis/dashboard) and give your project access to
   the Google Sheets API.
3. Create [OAuth 2.0 credentials](https://console.cloud.google.com/apis/credentials) for this API. You can create a web
   client, for example. Make sure you include a redirect URI, which can just be <https://osspledge.com>. Download the
   JSON file with the credentials, and put it in `credentials.json`.
4. Go to the [OAuth consent screen](https://console.cloud.google.com/apis/credentials/consent) settings, make sure your
   app is set to “Testing”, and add yourself as a “Test user”.
5. Run the program (see below).
6. When prompted, open the URL the program provides you with in order to give your project Google Sheets access to your
   account.
7. You'll be redirected to your redirect URL, and there will be a `code` parameter in the URL. Copy this `code`
   parameter into the console then hit `<Enter>`. A `token.json` file will then be saved.

## Running

To run locally:

```
go run .
```

To deploy, `go build -o bin/` and run `bin/osp-potential-member-list` from the root repo directory on your deployment.
