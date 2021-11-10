### Setup OTP

#### First thing, you need prepare API(webhook) with POST method, the payload format as below

```json
{
  "first_name": "Nolan",
  "last_name": "Nguyen",
  "otp": "2448", // OTP
  "phone_number": "8434700000", // Phone number
  "time": "1629743463", // UTC Seconds
  "user_id": "963106161" // User ID
}
```

#### create .env file as format below

```env
TELEGRAM_BOT=token telegram
WEBHOOK=API webhook
OPEN_URL= khi nhận OTP, sẽ xuất hiện nút ấn vào link này. maybe link của website
APP_ENV=production // env
```

#### deploy

```
tạo file .env xong run lệnh:
docker-compose up -d --build
```
