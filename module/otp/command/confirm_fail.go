package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type confirmFailCommand struct {
	ctx *commandCtx
}

func NewReceiveContactFailCommand(ctx *commandCtx) *confirmFailCommand {
	return &confirmFailCommand{ctx: ctx}
}

func IsConfirmFail(update *tgbotapi.Update) bool {
	message := update.Message
	if update.CallbackQuery != nil {
		message = update.CallbackQuery.Message
	}
	return message.Contact != nil && message.Contact.UserID != message.From.ID
}

const confirmFailMessageTmpl = `
Hệ thống đã kiểm tra, bạn không sử dụng đúng SĐT. vui lòng click /start để thực hiện tại các bước
`

func (c *confirmFailCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	chatID := c.ctx.update.Message.Chat.ID
	var _msg = tgbotapi.NewMessage(chatID, "")
	_msg.Text = confirmFailMessageTmpl
	_msg.ParseMode = "html"
	_msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	*msg = _msg
	return msg
}
