# Booking Documentation

## Create Booking



### Endpoint

```http
POST /api/booking
```
### Example Payload

```json
{
    "user_id" : "888849e0-7a32-4554-af86-7e9796466716",
    "event_id" : "054c589d-79b2-49e3-b77f-f59acabf1350",
    "quantity" : 3
}
```

### Example cURL

```sh
curl -X POST http://yourhostdomain.com/api/booking \
-H "Content-Type: application/json" \
-d '{
    "user_id" : "888849e0-7a32-4554-af86-7e9796466716",
    "event_id" : "054c589d-79b2-49e3-b77f-f59acabf1350",
    "quantity" : 3
}'
```

### Example Response

```json
{
    "message": "Event booked successfully",
    "booking": {
        "id": "5b03dd02-34fd-43a1-9a77-c7bbd6c19979",
        "user_id": "888849e0-7a32-4554-af86-7e9796466716",
        "event_id": "054c589d-79b2-49e3-b77f-f59acabf1350",
        "quantity": 3,
        "total_price": 899.97,
        "created_at": "2024-11-13T11:39:14.1085022+07:00",
        "updated_at": "2024-11-13T11:39:14.1085022+07:00"
    },
}
```



## Get All Booking



### Endpoint

```http
GET api/booking
```

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/booking \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "bookings": [
        {
            "id": "5b03dd02-34fd-43a1-9a77-c7bbd6c19979",
            "user_id": "888849e0-7a32-4554-af86-7e9796466716",
            "event_id": "054c589d-79b2-49e3-b77f-f59acabf1350",
            "quantity": 3,
            "total_price": 899.97,
            "CreatedAt": "2024-11-13T11:39:14.108502+07:00",
            "UpdatedAt": "2024-11-13T11:39:14.108502+07:00",
            "User": {
                "id": "888849e0-7a32-4554-af86-7e9796466716",
                "name": "John Doe",
                "email": "john@test.com",
                "password": "$2a$10$dixHYCXABGchwlEDY/0X6.YxVQEnlaWC56zCO5EDfaliCzFU88s56",
                "role": "user",
                "CreatedAt": "2024-11-12T18:52:01.941889+07:00",
                "UpdatedAt": "2024-11-12T18:52:01.941889+07:00",
                "Bookings": null
            },
            "Event": {
                "id": "054c589d-79b2-49e3-b77f-f59acabf1350",
                "name": "Tech Conference 2023",
                "location": "albuquerque, NM",
                "start_date": "2023-11-01T16:00:00+07:00",
                "end_date": "2023-11-04T00:00:00+07:00",
                "price": 299.99,
                "total_seat": 500,
                "available_seat": 187,
                "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
                "UpdatedAt": "2024-11-13T12:15:44.145021+07:00",
                "Bookings": null
            }
        },
        {
            "id": "42bc47f6-e5a5-4c3f-bd4f-eab8e5da1d09",
            "user_id": "888849e0-7a32-4554-af86-7e9796466716",
            "event_id": "054c589d-79b2-49e3-b77f-f59acabf1350",
            "quantity": 4,
            "total_price": 1199.96,
            "CreatedAt": "2024-11-13T12:15:44.11612+07:00",
            "UpdatedAt": "2024-11-13T12:15:44.11612+07:00",
            "User": {
                "id": "888849e0-7a32-4554-af86-7e9796466716",
                "name": "John Doe",
                "email": "john@test.com",
                "password": "$2a$10$dixHYCXABGchwlEDY/0X6.YxVQEnlaWC56zCO5EDfaliCzFU88s56",
                "role": "user",
                "CreatedAt": "2024-11-12T18:52:01.941889+07:00",
                "UpdatedAt": "2024-11-12T18:52:01.941889+07:00",
                "Bookings": null
            },
            "Event": {
                "id": "054c589d-79b2-49e3-b77f-f59acabf1350",
                "name": "Tech Conference 2023",
                "location": "albuquerque, NM",
                "start_date": "2023-11-01T16:00:00+07:00",
                "end_date": "2023-11-04T00:00:00+07:00",
                "price": 299.99,
                "total_seat": 500,
                "available_seat": 187,
                "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
                "UpdatedAt": "2024-11-13T12:15:44.145021+07:00",
                "Bookings": null
            }
        }
    ]
}
```



## Get Booking By ID



### Endpoint

```http
GET api/booking/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Booking ID |

### Example cURl

```sh
curl -X GET http://yourhostdomain.com/api/booking/5b03dd02-34fd-43a1-9a77-c7bbd6c19979 \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "booking": {
        "id": "5b03dd02-34fd-43a1-9a77-c7bbd6c19979",
        "user_id": "888849e0-7a32-4554-af86-7e9796466716",
        "event_id": "054c589d-79b2-49e3-b77f-f59acabf1350",
        "quantity": 3,
        "total_price": 899.97,
        "CreatedAt": "2024-11-13T11:39:14.108502+07:00",
        "UpdatedAt": "2024-11-13T11:39:14.108502+07:00",
        "User": {
            "id": "888849e0-7a32-4554-af86-7e9796466716",
            "name": "John Doe",
            "email": "john@test.com",
            "password": "$2a$10$dixHYCXABGchwlEDY/0X6.YxVQEnlaWC56zCO5EDfaliCzFU88s56",
            "role": "user",
            "CreatedAt": "2024-11-12T18:52:01.941889+07:00",
            "UpdatedAt": "2024-11-12T18:52:01.941889+07:00",
            "Bookings": null
        },
        "Event": {
            "id": "054c589d-79b2-49e3-b77f-f59acabf1350",
            "name": "Tech Conference 2023",
            "location": "albuquerque, NM",
            "start_date": "2023-11-01T16:00:00+07:00",
            "end_date": "2023-11-04T00:00:00+07:00",
            "price": 299.99,
            "total_seat": 500,
            "available_seat": 191,
            "CreatedAt": "2024-11-12T14:46:35.843218+07:00",
            "UpdatedAt": "2024-11-13T11:39:14.132148+07:00",
            "Bookings": null
        }
    }
}
```



## Delete Booking



### Endpoint

```http
DELETE api/booking/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** Booking ID |


### Example cURL

```sh
curl -X DELETE http://yourhostdomain.com/api/booking/5b03dd02-34fd-43a1-9a77-c7bbd6c19979 \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "message": "Booking canceled successfully"
}
```

