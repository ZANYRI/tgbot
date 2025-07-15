package main 

import (
	"fmt"
	"log"


	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

func main() {

	config, err := ReadEnv()

	if err != nil {
		log.Fatal("Ошибка!")
	} 

	b, err := gotgbot.NewBot(config.Token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewMessage(message.Text, answer))

	webhookOpts := ext.WebhookOpts{
		ListenAddr:  "localhost:8080",
		SecretToken: config.WebhookSecret, 
	}

	err = updater.StartWebhook(b, "custom-path/"+config.Token, webhookOpts)
	if err != nil {
		panic("failed to start webhook: " + err.Error())
	}

	err = updater.SetAllBotWebhooks(config.WebhookDomain, &gotgbot.SetWebhookOpts{
		MaxConnections:     100,
		DropPendingUpdates: true,
		SecretToken:        webhookOpts.SecretToken,
	})
	if err != nil {
		panic("failed to set webhook: " + err.Error())
	}

	log.Printf("%s has been started...\n", b.User.Username)

	updater.Idle()
}

func answer(b *gotgbot.Bot, ctx *ext.Context) error {
	response, err := ReqApi(ctx.EffectiveMessage.Text)
	if err != nil {
		return fmt.Errorf("failed to call ReqApi: %w", err)
	}

	_, err = ctx.EffectiveMessage.Reply(b, response, nil)
	if err != nil {
		return fmt.Errorf("failed to send reply: %w", err)
	}

	return nil
}