# my-trophy-prototype-api
API of My Trophy Prototype

## Prepare
### Add Firebase secret key file
```sh
$ mv /path/to/file firebase_secret_key.json
```

## Usage
```sh
$ export DB_HOST=localhost
$ export DB_PORT=3306
$ export DB_USER_NAME=user
$ export DB_PASSWORD=pass
$ export PORT=5000
$ export ALLOW_ORIGINS=http://localhost:3000
$ go run server.go
```

API is available on [localhost:5000](http://localhost:5000)

## API Specification
See [API_SPEC.md](API_SPEC.md)
