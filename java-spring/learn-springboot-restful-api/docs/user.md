# User API Spec

## Register User

- Endpoint : POST /api/users

Request Body :
```json
{
  "username": "username",
  "password": "password",
  "name": "user_name"
}
```

Response Body (Success):
```json
{
  "data": "OK"
}
```

Response Body (Failed):
```json
{
  "errors": "username must not blank"
}
```

## Login User
- Endpoint : POST /api/auth/login

Request Body :

```json
{
  "username": "username",
  "password": "password"
}
```

Response Body (Success):
```json
{
  "data": {
    "token": "TOKEN",
    "expiredAt": 12312313212
  }
}
```

Response Body (Failed):
```json
{
  "errors": "username or password is blank"
}
```

## Get User
- Endpoint : GET /api/users/current

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": {
    "username": "username",
    "name": "user_name"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "Unauthorized"
}
```

## Update User
- Endpoint : PATCH /api/users/current

Request Header:
- X-API-TOKEN: Token (Mandatory)

Request Body :

```json
{
  "name": "user_name",  // put if name want to be updated
  "password": "new_password"  // put if password want to be updated
}
```

Response Body (Success):
```json
{
  "data": {
    "username": "username",
    "name": "user_name"
  }
}
```

Response Body (Failed, 401):
```json
{
  "errors": "Unauthorized"
}
```

## Logout User
- Endpoint : DELETE /api/auth/logout

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": "OK"
}
```