package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sv-otp/common"
	otpbiz "github.com/sv-otp/module/otp/biz"
)

type confirmAndGenOTPCommand struct {
	ctx       *commandCtx
	submitBiz otpbiz.SubmitOtpBiz
}

func NewConfirmAndGenOTPCommand(ctx *commandCtx, submitBiz otpbiz.SubmitOtpBiz) *confirmAndGenOTPCommand {
	return &confirmAndGenOTPCommand{ctx: ctx, submitBiz: submitBiz}
}

func IsConfirmAndGenOTP(update *tgbotapi.Update) bool {
	message := update.Message
	return message.Contact != nil && message.Contact.UserID == message.From.ID
}

const confirmAndGenOTPMessageTmpl = `
Cảm ơn bạn đã chia sẽ số điện thoại: <b>{{.PhoneAddrs}}</b>
Mã OTP của bạn là: <b>{{.OTP}} </b>
MÃ OTP không được được bất cứ ai, mã OTP chỉ có hiệu lực trong 3 Phút
`

func (c *confirmAndGenOTPCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	otp := common.EncodeToString(4)
	vars := make(map[string]interface{})
	vars["PhoneAddrs"] = c.ctx.update.Message.Contact.PhoneNumber
	vars["OTP"] = otp
	messResp := common.ProcessString(confirmAndGenOTPMessageTmpl, vars)

	telegramUser := &common.TelegramUser{
		PhoneNumber: c.ctx.update.Message.Contact.PhoneNumber,
		FirstName:   c.ctx.update.Message.Contact.FirstName,
		LastName:    c.ctx.update.Message.Contact.LastName,
		UserID:      c.ctx.update.Message.Contact.UserID,
		Time:        c.ctx.update.Message.Time(),
		Otp:         otp,
	}
	c.submitBiz.SubmitOtp(telegramUser)

	env := c.ctx.appCtx.GetEnv()
	urlOpen := env["openUrl"]
	var tokenUrlKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Nhập OTP trên website", urlOpen),
		),
	)

	chatID := c.ctx.update.Message.Chat.ID

	var _msg = tgbotapi.NewMessage(chatID, "")
	_msg.Text = messResp
	_msg.ParseMode = "html"
	_msg.ReplyMarkup = tokenUrlKeyboard
	*msg = _msg
	return msg
}
