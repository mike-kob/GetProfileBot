package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	api "github.com/mike-kob/GetProfileBot/api"
)

func main() {
	bot := &tg.BotAPI{
		Token:  os.Getenv("TOKEN"),
		Client: &http.Client{},
		Buffer: 100,
	}
	url := fmt.Sprintf("https://YOUR_URL/update?t=%s", os.Getenv("TOKEN"))
	res, _ := bot.SetWebhook(tg.NewWebhook(url))
	fmt.Printf("%v", res)
	http.HandleFunc("/update", api.Handler)

	log.Println("Listing for updates")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
