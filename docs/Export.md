# Export Event with Queue

## Get All Event

### Endpoint

```http
GET /api/export/events
```

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/export/events \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "message": "Events exported successfully",
}
```

## Export Booking by User ID

### Endpoint

```http
GET /api/export/booking/:id
```

| Params | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** User ID |

### Example cURL

```sh
curl -X GET http://yourhostdomain.com/api/export/booking/888849e0-7a32-4554-af86-7e9796466716 \
-H "Content-Type: application/json"
```

### Example Response

```json
{
    "message": "Bookings exported successfully",
}
```

