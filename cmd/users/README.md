### setting
```bash
sudo apt-get update
sudo apt-get install sqlite3
sqlite3 --version
```

```bash
vscode ➜ /workspaces/go-api-samples/cmd/users (main) $ sqlite3 local.db
SQLite version 3.34.1 2021-01-20 14:10:07
Enter ".help" for usage hints.
sqlite> CREATE TABLE IF NOT EXISTS users ( user_id INTEGER PRIMARY KEY, email_address TEXT, created_at INTEGER, deleted INTEGER, settings TEXT);
sqlite> .tables
users
sqlite> INSERT INTO users (user_id, email_address, created_at, deleted, settings) VALUES (1, 'maria@example.com', 0, 0, '');
sqlite> INSERT INTO users (user_id, email_address, created_at, deleted, settings) VALUES (999, 'admin@example.com', 0, 0, '');
sqlite> pragma table_info(users);
0|user_id|INTEGER|0||1
1|email_address|TEXT|0||0
2|created_at|INTEGER|0||0
3|deleted|INTEGER|0||0
4|settings|TEXT|0||0
sqlite> .exit
vscode ➜ /workspaces/go-api-samples/cmd/users (main) $
```

```bash
export DATABASE_URL="local.db"
echo $DATABASE_URL
```

### note
- driver blank import
