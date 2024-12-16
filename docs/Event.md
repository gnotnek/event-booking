# Event Management Documentation
Some endpoints just accessible by admin account, the usual user just can see all the listed event

# Admin Only

## Create New Event



### Endpoint

```http
POST /api/event
```
### Example Payload

```json
{
  "name": "Music Festival 2023",
  "location": "Los Angeles, CA",
  "start_date": "2023-12-15T18:00:00Z",
  "end_date": "2023-12-17T23:00:00Z",
  "price": 150.00,
  "total_seat": 1000,
  "available_seat": 750
}
```

### Example cURL

```sh
curl -X POST http://yourhostdomain.com/api/event \
-H "Content-Type: application/json" \
-d '{
    "name": "Music Festival 2023",
    "location": "Los Angeles, CA",
    "start_date": "2023-12-15T18:00:00Z",
    "end_date": "2023-12-17T23:00:00Z",
    "price": 150.00,
    "total_seat": 1000,
    "available_seat": 750
}'
```

### Example Response

```json
{
    "message": "Event created successfully",
    "event": {
        "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
        "name": "Tech Conference 2023",
        "location": "San Francisco, CA",
        "start_date": "2023-11-01T09:00:00Z",
        "end_date": "2023-11-03T17:00:00Z",
        "price": 299.99,
        "total_seat": 500,
        "available_seat": 150,
        "CreatedAt": "2024-11-12T14:46:35.8432188+07:00",
        "UpdatedAt": "2024-11-12T14:46:35.8432188+07:00",
    },
}
```



## Get All Event



### Endpoint

```http
GET /api/event
```

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/event \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "events": [
        {
            "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
            "name": "Tech Conference 2023",
            "location": "San Francisco, CA",
            "start_date": "2023-11-01T16:00:00+07:00",
            "end_date": "2023-11-04T00:00:00+07:00",
            "price": 299.99,
            "total_seat": 500,
            "available_seat": 150,
            "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
            "UpdatedAt": "2024-11-12T14:46:35.843218+07:00",
        }
    ]
}
```



## Get Event by ID



### Endpoint

```http
GET /api/event/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Event ID |

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/event/391ced0f-26b6-4bc3-8019-d8dc805051bf \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
    "name": "Tech Conference 2023",
    "location": "San Francisco, CA",
    "start_date": "2023-11-01T16:00:00+07:00",
    "end_date": "2023-11-04T00:00:00+07:00",
    "price": 299.99,
    "total_seat": 500,
    "available_seat": 150,
    "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
    "UpdatedAt": "2024-11-12T14:46:35.843218+07:00"
}
```



## Edit Event



### Endpoint

```http
POST /api/event/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Event ID |

### Example Payload

```json
{
  "name": "Music Festival 2023",
  "location": "Los Angeles, CA",
  "start_date": "2023-12-15T18:00:00Z",
  "end_date": "2023-12-17T23:00:00Z",
  "price": 150.00,
  "total_seat": 1000,
  "available_seat": 750
}
```

### Example cURL

```sh
curl -X POST http://yourhostdomain.com/api/event/391ced0f-26b6-4bc3-8019-d8dc805051bf \
-H "Content-Type: application/json" \
-d '{
    "name": "Music Festival 2023",
    "location": "Los Angeles, CA",
    "start_date": "2023-12-15T18:00:00Z",
    "end_date": "2023-12-17T23:00:00Z",
    "price": 150.00,
    "total_seat": 1000,
    "available_seat": 999
}'
```

### Example Response

```json
{
    "message": "Event created successfully",
    "event": {
        "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
        "name": "Tech Conference 2023",
        "location": "San Francisco, CA",
        "start_date": "2023-11-01T09:00:00Z",
        "end_date": "2023-11-03T17:00:00Z",
        "price": 299.99,
        "total_seat": 500,
        "available_seat": 999,
        "CreatedAt": "2024-11-12T14:46:35.8432188+07:00",
        "UpdatedAt": "2024-11-12T14:46:35.8432188+07:00",
    },
}
```



## Delete Event



### Endpoint

```http
POST /api/event/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Event ID |

### Example cURL

```sh
curl -X POST http://yourhostdomain.com/api/event/391ced0f-26b6-4bc3-8019-d8dc805051bf \
-H "Content-Type: application/json" \
```

### Example Response

```json
{
    "message": "Event deleted successfully"
}
```



<br />

# Not Admin Access
## Get All Event



### Endpoint

```http
GET /api/event
```

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/event \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "events": [
        {
            "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
            "name": "Tech Conference 2023",
            "location": "San Francisco, CA",
            "start_date": "2023-11-01T16:00:00+07:00",
            "end_date": "2023-11-04T00:00:00+07:00",
            "price": 299.99,
            "total_seat": 500,
            "available_seat": 150,
            "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
            "UpdatedAt": "2024-11-12T14:46:35.843218+07:00",
        }
    ]
}
```



## Get Event by ID



### Endpoint

```http
GET /api/event/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Event ID |

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/event/391ced0f-26b6-4bc3-8019-d8dc805051bf \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "id": "391ced0f-26b6-4bc3-8019-d8dc805051bf",
    "name": "Tech Conference 2023",
    "location": "San Francisco, CA",
    "start_date": "2023-11-01T16:00:00+07:00",
    "end_date": "2023-11-04T00:00:00+07:00",
    "price": 299.99,
    "total_seat": 500,
    "available_seat": 150,
    "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
    "UpdatedAt": "2024-11-12T14:46:35.843218+07:00"
}
```


