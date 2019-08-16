# my-trophy-prototype-api
API of My Trophy Prototype

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
### GET /users/:user_id
#### Response
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>created_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>updated_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>deleted_at</td><td>string | null</td><td>RFC 3339</td></tr>
<tr><td>name</td><td>string</td><td></td></tr>
<tr><td>email</td><td>string</td><td></td></tr>
</table>
