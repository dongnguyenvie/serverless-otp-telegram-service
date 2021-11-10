package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sv-otp/common"
)

type WebhookProvider interface {
	OtpCallback(user *common.TelegramUser) error
}

type webhookProvider struct {
	webhookUrl string
}

func NewWebhookProvider(webhookUrl string) *webhookProvider {
	return &webhookProvider{webhookUrl: webhookUrl}
}

func (c *webhookProvider) OtpCallback(user *common.TelegramUser) error {
	go func() {
		var params = map[string]string{
			"phone_number": user.PhoneNumber,
			"last_name":    user.LastName,
			"first_name":   user.FirstName,
			"otp":          user.Otp,
			"user_id":      strconv.Itoa(user.UserID),
			"time":         strconv.FormatInt(user.Time.Unix(), 10),
		}
		postBody, _ := json.Marshal(params)
		responseBody := bytes.NewBuffer(postBody)
		request, _ := http.NewRequest("POST", c.webhookUrl, responseBody)
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("HTTP call failed:", err)
			return
		}
		defer response.Body.Close()
	}()
	return nil
}
