# Address API Spec

## Create Address
Endpoint: POST /api/contacts/{contactID}/addresses

Request Header:
- X-API-TOKEN: Token (Mandatory)

Request Body:
```json
{
  "street": "street",
  "city": "city",
  "province": "province",
  "country": "country",
  "postal_code": "12345"
}
```

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "street": "street",
    "city": "city",
    "province": "province",
    "country": "country",
    "postal_code": "12345"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "Unauthorized"
}
```

## Update Address
Endpoint: PUT /api/contacts/{contactID}/addresses/{addressID}

Request Header:
- X-API-TOKEN: Token (Mandatory)

Request Body:
```json
{
  "street": "street",
  "city": "city",
  "province": "province",
  "country": "country",
  "postal_code": "12345"
}
```

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "street": "street",
    "city": "city",
    "province": "province",
    "country": "country",
    "postal_code": "12345"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "Address is not found"
}
```


## Get Address
Endpoint: GET /api/contacts/{contactID}/addresses/{addressID}

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "street": "street",
    "city": "city",
    "province": "province",
    "country": "country",
    "postal_code": "12345"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "Address is not found"
}
```


## Remove Address
Endpoint: DELETE /api/contacts/{contactID}/addresses/{addressID}

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": "OK"
}
```

Response Body (Failed):
```json
{
  "errors": "Address is not found"
}
```


## List Address
Endpoint: GET /api/contacts/{contactID}/addresses

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": [
    {
      "id": "uuid",
      "street": "street",
      "city": "city",
      "province": "province",
      "country": "country",
      "postal_code": "12345"
    }
  ]
}
```

Response Body (Failed):
```json
{
  "errors": "Unauthorized"
}
```
