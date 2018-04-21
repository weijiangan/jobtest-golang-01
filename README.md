# jobtest-golang-01
Simple gRPC server client implementation

# Usage:

## Config file
Add a JSON file to the same root directory of the .go file or executable
```json
{
  "user": "Postgres username",
  "password": "Postgres password",
  "database": "Postgres database"
}
```

## Client
CLI options:
```
-s [client_ip] [server_ip] tags(JSON) message
-q [client_ip] [server_ip] [tags(JSON)]
```

If you decide to opt out any [optional] arguments, replace with empty string `""`
