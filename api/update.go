package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

func start(bot *tg.BotAPI, update *tg.Update) {
	msgText := "Hello!\n" +
		"I am GetMyProfile bot. I return information that you share with all bots, " +
		"when writing to them.\n\n" +
		"Tap /getMe to get info."

	msg := tg.NewMessage(update.Message.Chat.ID, msgText)
	_, _ = bot.Send(msg)
}

func getMe(bot *tg.BotAPI, update *tg.Update) {
	profile := update.Message.From
	msgText := fmt.Sprintf(
		"Your info available to bots:\n\n"+
			"id: %d\n"+
			"username: @%s\n"+
			"first_name: %s\n"+
			"last_name: %s\n"+
			"language_code: %s",
		profile.ID,
		profile.UserName,
		profile.FirstName,
		profile.LastName,
		profile.LanguageCode,
	)
	msg := tg.NewMessage(update.Message.Chat.ID, msgText)
	_, _ = bot.Send(msg)
}

func Handler(w http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("t") != os.Getenv("TOKEN") {
		w.WriteHeader(401)
		return
	}

	bot := &tg.BotAPI{
		Token:  os.Getenv("TOKEN"),
		Client: &http.Client{},
		Buffer: 100,
	}

	var update tg.Update
	_ = json.NewDecoder(request.Body).Decode(&update)

	switch update.Message.Text {
	case "/start":
		start(bot, &update)
	case "/getMe":
		getMe(bot, &update)
	}
}
