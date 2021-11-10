package appctx

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sv-otp/component/webhook"
)

type AppContext interface {
	GetBot() *tgbotapi.BotAPI
	WebhookProvider() webhook.WebhookProvider
	GetEnv() map[string]string
}

type appCtx struct {
	bot             *tgbotapi.BotAPI
	webhookProvider webhook.WebhookProvider
	env             map[string]string
}

func NewAppContext(bot *tgbotapi.BotAPI, webhookProvider webhook.WebhookProvider, env map[string]string) *appCtx {
	return &appCtx{bot: bot, webhookProvider: webhookProvider, env: env}
}

func (c *appCtx) GetBot() *tgbotapi.BotAPI {
	return c.bot
}

func (c *appCtx) WebhookProvider() webhook.WebhookProvider {
	return c.webhookProvider
}

func (c *appCtx) GetEnv() map[string]string {
	return c.env
}
