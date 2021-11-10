package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type callbackGetOtpCommand struct {
	ctx *commandCtx
}

func NewCallbackGetOtpCommand(ctx *commandCtx) *callbackGetOtpCommand {
	return &callbackGetOtpCommand{ctx: ctx}
}

func IsCallbackGetOtpCommand(update *tgbotapi.Update) bool {
	message := update.CallbackQuery.Data
	return message == "/get_OTP"
}

const getCallbackOtpMessageTmpl = `
hệ thống cần số diện thoại của bạn để xác thực. Vui lòng ấn vào nút <b>Chia sẻ số điện thoại để tiếp tục</b>
`

var callbackSharePhoneAddrKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("Chia sẽ số điện thoại"),
	),
)

func (c *callbackGetOtpCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	chatID := c.ctx.update.CallbackQuery.Message.Chat.ID
	_msg := tgbotapi.NewMessage(chatID, "")
	_msg.Text = getCallbackOtpMessageTmpl
	_msg.ParseMode = "html"
	_msg.ReplyMarkup = callbackSharePhoneAddrKeyboard
	*msg = _msg
	return msg
}
