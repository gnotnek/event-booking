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



## Log In User



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

