# Contact API Spec

## Create Contact
Endpoint : POST /api/contacts

Request Header:
- X-API-TOKEN: Token (Mandatory)

Request Body:
```json
{
  "first_name": "first name",
  "last_name": "last name",
  "email": "mail@mail.com",
  "phone": "08123456789"
}
```

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "first_name": "first name",
    "last_name": "last name",
    "email": "mail@mail.com",
    "phone": "08123456789"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "email format invalid, phone format invalid"
}
```

## Update Contact
Endpoint : PUT /api/contacts/{contactID}

Request Header:
- X-API-TOKEN: Token (Mandatory)

Request Body:
```json
{
  "first_name": "first name",
  "last_name": "last name",
  "email": "mail@mail.com",
  "phone": "08123456789"
}
```

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "first_name": "first name",
    "last_name": "last name",
    "email": "mail@mail.com",
    "phone": "08123456789"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "email format invalid, phone format invalid"
}
```

## Get Contact
Endpoint : GET /api/contacts/{contactID}

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": {
    "id": "uuid",
    "first_name": "first name",
    "last_name": "last name",
    "email": "mail@mail.com",
    "phone": "08123456789"
  }
}
```

Response Body (Failed):
```json
{
  "errors": "contact is not found"
}
```

## Search Contact
Endpoint : GET /api/contacts

Query Param :
- name: String, contact first name or last name, using like query, optional
- phone: String, contact phone, using like query, optional
- email: String, contact email, using like query, optional
- page: Integer, start from 0, default 0
- size: Integer, default 10

Request Header:
- X-API-TOKEN: Token (Mandatory)

Response Body (Success):
```json
{
  "data": [
    {
      "id": "uuid",
      "first_name": "first name",
      "last_name": "last name",
      "email": "mail@mail.com",
      "phone": "08123456789"
    }
  ],
  "paging": {
    "current_page": 0,
    "total_page": 10,
    "size": 10
  }
}
```

Response Body (Failed):
```json
{
  "errors": "Unauthorized"
}
```

## Remove Contact
Endpoint : DELETE /api/contacts/{contactID}

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
  "errors": "contact is not found"
}
```