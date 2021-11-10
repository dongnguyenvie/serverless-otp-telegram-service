package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type mirrorCommand struct {
	ctx *commandCtx
}

func NewMirrorCommand(ctx *commandCtx) *mirrorCommand {
	return &mirrorCommand{ctx: ctx}
}

func IsMirrorCommandCase(update *tgbotapi.Update) bool {
	message := update.Message
	return message.Text == "/mirror_reply"
}

func (c *mirrorCommand) Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig {
	*msg = tgbotapi.NewMessage(c.ctx.update.Message.Chat.ID, c.ctx.update.Message.Text)
	return msg
}
