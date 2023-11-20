# go-api-samples
A collection of REST API samples implemented in Go.

## `/storage`
- [Docs](./cmd/storage/README.md)

## `/users`
- [Docs](./cmd/users/README.md)

## server
- env
```bash
export GOOGLE_APPLICATION_CREDENTIALS="/workspaces/go-api-samples/key.json"
echo $GOOGLE_APPLICATION_CREDENTIALS
export DATABASE_URL="local.db"
echo $DATABASE_URL
```

- installation
```bash
sudo apt-get update
sudo apt-get install sqlite3
sqlite3 --version
```

- `/users`
```bash
vscode ➜ /workspaces/go-api-samples (feature/add-user-to-server) $ curl "localhost:8080/users?user_id=1" -i
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 16 Nov 2023 03:29:45 GMT
Content-Length: 90

{"user_id":1,"email_address":"maria@example.com","created_at":0,"deleted":0,"settings":""}
vscode ➜ /workspaces/go-api-samples (feature/add-user-to-server) $ 
```

- `/storage`
```bash
vscode ➜ /workspaces/go-api-samples (feature/add-user-to-server) curl "localhost:8080/storage?bucket=sanbox-334000_bucket&object=test.html" -i
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 16 Nov 2023 03:31:51 GMT
Content-Length: 424

{"content":"\u003c!DOCTYPE html\u003e\n\u003chtml lang=\"en\"\u003e\n\u003chead\u003e\n    \u003cmeta charset=\"UTF-8\"\u003e\n    \u003cmeta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"\u003e\n    \u003cmeta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"\u003e\n    \u003ctitle\u003eDocument\u003c/title\u003e\n\u003c/head\u003e\n\u003cbody\u003e\n    test\n\u003c/body\u003e\n\u003c/html\u003e"}
vscode ➜ /workspaces/go-api-samples (feature/add-user-to-server) $
```