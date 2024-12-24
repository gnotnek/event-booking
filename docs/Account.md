# User Account Documentation

## Register User



### Endpoint

```http
 POST /api/signup
```

### Example Body Request
```json
{
    "name" : "John Doe",
    "email" : "john@test.com",
    "password" : "12345"
}
```
### Example cURL Request
```sh
curl -X POST http://yourhostdomain.com/api/signup \
-H "Content-Type: application/json" \
-d '{
    "name": "John Doe",
    "email": "john@test.com",
    "password": "12345"
}'
```



## Register User Admin



### Endpoint

```http
 POST /api/signup
```

### Example Body Request
```json
{
    "name" : "John admin",
    "email" : "johnadmin@test.com",
    "password" : "12345",
    "role" : "admin"
}
```
### Example cURL Request
```sh
curl -X POST http://yourhostdomain.com/api/signup \
-H "Content-Type: application/json" \
-d '{
    "name" : "John admin",
    "email" : "johnadmin@test.com",
    "password" : "12345",
    "role" : "admin"
}'
```


## Send Verification Email

Send 6 digit code to user email for verification purposes.

### Endpoint
```http
POST /api/account/send-verification
```

### Example Payload

```json
{
    "email" : "johndoe@test.com"
}
```

### Example cURL Request
```sh
curl -X POST http://yourhostdomain.com/api/account/send-verification \
-H "Content-Type: application/json" \
-d '
    {
    "email" : "johndoe@test.com"
}'
```

### Example Response

```json
{
    "Message" : "Verification code sent to your email successfully"
}
```

## Validate Verification Code

### Endpoint
```http
POST /api/account/validate
```

### Example Payload
```json
{
    "code" : "123456",
    "email" : "johndoe@test.com"
}
```

### Example cURL Request
```sh
curl -X POST http://yourhostdomain.com/api/account/validate \
-H "Content-Type: application/json" \
-d '{
    "code" : "123456",
    "email" : "johndoe@test.com"
}'
```

### Example Response
```json
{
    "message":"Verification code validated successfully"
}
```

## Log In User

Before User can log in, they must verify their email first

### Endpoint

```http
POST /api/login
```

### Example Payload

```json
{
    "email" : "johnadmin@test.com",
    "password" : "12345"
}
```

### Example cURL Request
```sh
curl -X POST http://yourhostdomain.com/api/login \
-H "Content-Type: application/json" \
-d '{
    "email" : "johnadmin@test.com",
    "password" : "12345"
}'
```

