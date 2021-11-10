package otpbiz

import (
	"github.com/sv-otp/common"
)

type SubmitOtpBiz interface {
	SubmitOtp(user *common.TelegramUser)
}

type WebhookProvider interface {
	OtpCallback(user *common.TelegramUser) error
}

type submitOtpBiz struct {
	webhookProvider WebhookProvider
}

func NewSubmitOtpBiz(webhookProvider WebhookProvider) *submitOtpBiz {
	return &submitOtpBiz{webhookProvider: webhookProvider}
}

func (biz *submitOtpBiz) SubmitOtp(user *common.TelegramUser) {
	biz.webhookProvider.OtpCallback(user)
}
