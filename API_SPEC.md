# API Specification
## Table of contents
* [Overview](#heading-overview)
  * [Format](#heading-overview-format)
  * [Error Response](#heading-overview-error-response)
* [User](#heading-user)
  * [GET /users/](#heading-user-get-users)
  * [GET /users/:id](#heading-user-get-user)
* [Trophy](#heading-trophy)
  * [GET /trophies/](#heading-trophy-get-trophies)
  * [GET /trophies/:id](#heading-trophy-get-trophy)
  * [POST /trophies/](#heading-trophy-post-trophy)
  * [PUT /trophies/:id](#heading-trophy-put-trophy)
  * [DELETE /trophies/:id](#heading-trophy-delete-trophy)

<h2 id="heading-overview">Overview</h2>
<h3 id="heading-overview-format">Format</h3>
Use JSON as format of request/response.

<h3 id="heading-overview-error-response">Error Response</h3>

Error response is conforming to [RFC 7807](https://tools.ietf.org/html/rfc7807).

<h2 id="heading-user">User</h2>
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>created_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>updated_at</td><td>string</td><td>RFC 3339</td></tr>
<tr><td>deleted_at</td><td>string | null</td><td>RFC 3339</td></tr>
<tr><td>name</td><td>string</td><td></td></tr>
<tr><td>email</td><td>string</td><td></td></tr>
</table>

<h3 id="heading-user-get-users">GET /users/</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>id</td><td>number</td><td></td></tr>
<tr><td>email</td><td>string</td><td></td></tr>
</table>

#### Response
###### Code
|Code|Condition|
|---|---|
|200|Success|
|400|Fail|

##### Body
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

<h3 id="heading-user-get-user">GET /users/:id</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>:id</td><td>number</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|200|Success|
|400|Fail|

##### Body
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

<h2 id="heading-trophy">Trophy</h2>
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

<h3 id="heading-trophy-get-trophies">GET /trophies/</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>user_id</td><td>number</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|200|Success|
|400|Fail|

##### Body
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

<h3 id="heading-trophy-get-trophy">GET /trophies/:id</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>:id</td><td>number</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|200|Success|
|400|Fail|

##### Body
`Trophy`

#### Example
```sh
$ curl localhost:5000/trophies/1
```

```json
{
  "id": 1,
  "created_at": "2019-08-27T13:04:13+09:00",
  "updated_at": "2019-08-27T13:04:13+09:00",
  "title": "誕生",
  "description": "この世に生を受けた。",
  "user_id": 1
}
```

<h3 id="heading-trophy-post-trophy">POST /trophies/</h3>

#### Request
##### Body
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>title</td><td>string</td><td></td></tr>
<tr><td>description</td><td>string</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|201|Success|
|400|Fail|

#### Example
```sh
$ curl -X POST -H "Authorization:Bearer $JWT" -H 'Content-Type:application/json' -d '{"title":"My title","description":"My description"}' localhost:5000/trophies/
```

<h3 id="heading-trophy-put-trophy">PUT /trophies/:id</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>:id</td><td>number</td><td></td></tr>
</table>

##### Body
<table>
<tr><th>Property</th><th>Type</th><th>Format</th></tr>
<tr><td>title</td><td>string</td><td></td></tr>
<tr><td>description</td><td>string</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|204|Success|
|400|Fail|

#### Example
```sh
$ curl -X PUT -H "Authorization:Bearer $JWT" -H 'Content-Type:application/json' -d '{"title":"My title","description":"My description"}' localhost:5000/trophies/1
```

<h3 id="heading-trophy-delete-trophy">DELETE /trophies/:id</h3>

#### Request
##### URL
<table>
<tr><th>Parameter</th><th>Type</th><th>Format</th></tr>
<tr><td>:id</td><td>number</td><td></td></tr>
</table>

#### Response
##### Code
|Code|Condition|
|---|---|
|204|Success|
|400|Fail|

#### Example
```sh
$ curl -X DELETE -H "Authorization:Bearer $JWT" localhost:5000/trophies/1
```
