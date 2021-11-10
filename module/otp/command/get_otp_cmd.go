package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type getOtpCommand struct {
	ctx *commandCtx
}

func NewGetOtpCommand(ctx *commandCtx) *getOtpCommand {
	return &getOtpCommand{ctx: ctx}
}

func IsGetOtpCommand(update *tgbotapi.Update) bool {
	message := update.Message
	return message.Text == "/get_OTP"
}

const getOtpMessageTmpl = `
hệ thống cần số diện thoại của bạn để xác thực. Vui lòng ấn vào nút <b>Chia sẻ số điện thoại để tiếp tục</b>
`

var sharePhoneAddrKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("Chia sẽ số điện thoại"),
	),
)

func (c *getOtpCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	chatID := c.ctx.update.Message.Chat.ID
	_msg := tgbotapi.NewMessage(chatID, "")
	_msg.Text = getOtpMessageTmpl
	_msg.ParseMode = "html"
	_msg.ReplyMarkup = sharePhoneAddrKeyboard
	*msg = _msg
	return msg
}
