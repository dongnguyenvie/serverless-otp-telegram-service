package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type otpCommand struct {
	ctx *commandCtx
}

func NewStartCommand(ctx *commandCtx) *otpCommand {
	return &otpCommand{ctx: ctx}
}

func IsStartCommandCase(update *tgbotapi.Update) bool {
	message := update.Message
	return message.Text == "/start"
}

var startKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Lấy mã OTP", "/get_OTP"),
	),
)

const startMessage = `Chào bạn, để sử dụng dịch vụ OTP market7777.
Bạn muốn lấy mã OTP, Vui lòng click nút bên dưới
`

func (c *otpCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	var _msg = tgbotapi.NewMessage(c.ctx.update.Message.Chat.ID, "")
	_msg.Text = startMessage
	_msg.ParseMode = "html"
	_msg.ReplyMarkup = startKeyboard
	*msg = _msg
	return msg
}
