## `/storage`
### setup
```bash
go get google.golang.org/api/option
go get google.golang.org/api/storage/v1
export GOOGLE_APPLICATION_CREDENTIALS="/workspaces/go-api-samples/key.json"
echo $GOOGLE_APPLICATION_CREDENTIALS
go run cmd/storage/main.go 
```

### test
```bash
curl "localhost:8080/storage?bucket=sanbox-334000_bucket&object=test.txt" -i
curl "localhost:8080/storage?bucket=sanbox-334000_bucket&object=test.html" -i
```

### note
- ResponseWriter の設定
- Hanlder の使い方
- `type` の書式
- `os.Getenv` の漏れ
- `defer rc.Body.Close()` の漏れ
