### The schema is valid
```json
{
  "ok": true,
  "result": [
    {
      "update_id": 129413628,
      "message": {
        "message_id": 299,
        "from": {
          "id": 963106161, ==> so sánh cái này || send tin nhăn lên
          "is_bot": false,
          "first_name": "Dong",
          "last_name": "Nguyen",
          "username": "quynhdev",
          "language_code": "vi"
        },
        "chat": {
          "id": 963106161, ==> so sánh cái này
          "first_name": "Dong",
          "last_name": "Nguyen",
          "username": "quynhdev",
          "type": "private"
        },
        "date": 1629635954,
        "contact": { // nội dung
          "phone_number": "84347884884",
          "first_name": "Dong",
          "last_name": "Nguyen",
          "user_id": 963106161 ==> so sánh cái này giống mấy cái kia là số đt chính chủ
        }
      }
    }
  ]
}

```


### The schema is invalid
```json
{
  "ok": true,
  "result": [
    {
      "update_id": 129413629,
      "message": {
        "message_id": 301,
        "from": {
          "id": 963106161,
          "is_bot": false,
          "first_name": "Dong",
          "last_name": "Nguyen",
          "username": "quynhdev",
          "language_code": "vi"
        },
        "chat": {
          "id": 963106161,
          "first_name": "Dong",
          "last_name": "Nguyen",
          "username": "quynhdev",
          "type": "private"
        },
        "date": 1629636058,
        "contact": {
          "phone_number": "84942248638",
          "first_name": "A M\u1eadp Php",
          "user_id": 1093910281
        }
      }
    }
  ]
}

```
