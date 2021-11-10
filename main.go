package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	godotenv "github.com/joho/godotenv"
	"github.com/sv-otp/component/appctx"
	"github.com/sv-otp/component/webhook"
	subotp "github.com/sv-otp/module/otp/transport/subotp"
)

func main() {
	godotenv.Load(".env")

	botToken := os.Getenv("TELEGRAM_BOT")
	webhookUrl := os.Getenv("WEBHOOK")
	openUrl := os.Getenv("OPEN_URL")
	appEnv := os.Getenv("APP_ENV")
	webhookProvider := webhook.NewWebhookProvider(webhookUrl)
	env := map[string]string{
		"botToken":   botToken,
		"webhookUrl": webhookUrl,
		"openUrl":    openUrl,
	}

	bot, err := tgbotapi.NewBotAPI(botToken)

	if appEnv != "production" {
		bot.Debug = true
	}

	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	appCtx := appctx.NewAppContext(bot, webhookProvider, env)
	subotp.BotOtp(appCtx, bot)

}
