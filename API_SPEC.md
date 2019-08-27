# API Specification
## Model
### User
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>created_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>updated_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>deleted_at</td><td>string | null</td><td>RFC 3339</td></tr>
<tr><td>name</td><td>string</td><td></td></tr>
<tr><td>email</td><td>string</td><td></td></tr>
</table>

### Trophy
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>created_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>updated_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>deleted_at</td><td>string | null</td><td>RFC 3339</td></tr>
<tr><td>title</td><td>string</td><td></td></tr>
<tr><td>description</td><td>string</td><td></td></tr>
<tr><td>user_id</td><td>number</td><td></td></tr>
</table>

## API
### GET /users/
#### Request
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>email</td><td>string</td><td></td></tr>
</table>

#### Response
Array of `User`

#### Example
```sh
$ curl localhost:5000/users/?email=mail@munieru.jp
```

```json
[
  {
    "id": 1,
    "created_at": "2019-08-27T13:04:13+09:00",
    "updated_at": "2019-08-27T13:04:13+09:00",
    "name": "ムニエル",
    "email": "mail@munieru.jp"
  },
  {
    "id": 2,
    "created_at": "2019-08-27T13:04:13+09:00",
    "updated_at": "2019-08-27T13:04:13+09:00",
    "name": "田中太郎",
    "email": "tanaka@example.com"
  }
]
```

### GET /users/:user_id
#### Request
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>:user_id</td><td>number</td><td></td></tr>
</table>

#### Response
`User`

#### Example
```sh
$ curl localhost:5000/users/1
```

```json
{
  "id": 1,
  "created_at": "2019-08-27T11:38:02+09:00",
  "updated_at": "2019-08-27T11:38:02+09:00",
  "name": "ムニエル",
  "email": "mail@munieru.jp"
}
```


### GET /trophies/
#### Request
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>user_id</td><td>number</td><td></td></tr>
</table>

#### Response
Array of `Trophy`

#### Example
```sh
$ curl localhost:5000/trophies/?user_id=1
```

```json
[
  {
    "id": 1,
    "created_at": "2019-08-27T13:04:13+09:00",
    "updated_at": "2019-08-27T13:04:13+09:00",
    "title": "誕生",
    "description": "この世に生を受けた。",
    "user_id": 1
  },
  {
    "id": 2,
    "created_at": "2019-08-27T13:04:13+09:00",
    "updated_at": "2019-08-27T13:04:13+09:00",
    "title": "My Trophy作成",
    "description": "My Trophyを作った。",
    "user_id": 1
  }
]
```
