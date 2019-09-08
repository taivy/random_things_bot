package main

import (
	"api"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/net/proxy"
)

const (
	TOKEN = "YOUR_TOKEN"
	PROXY = "socks5://999.999.999.999:9999"
)

func getHelp() string {
	reply :=
		`This bot can send some random things. Available commands: 
	/cat_fact - Random fact about cats
	/geek_joke - Random geek joke
	/fox_pic - Random picture of fox
	/fact - Random (useless) fact
	/xkcd - Random XKCD comics
	/tech_quote - Random quote about tech and computers
	/startup_quote - Random famous quote by people from startup ecosystem
	`
	return reply
}

func main() {
	client := &http.Client{}
	if len(PROXY) > 0 {
		tgProxyURL, err := url.Parse(PROXY)
		if err != nil {
			log.Printf("Failed to parse proxy URL:%s\n", err)
		}
		tgDialer, err := proxy.FromURL(tgProxyURL, proxy.Direct)
		if err != nil {
			log.Printf("Failed to obtain proxy dialer: %s\n", err)
		}
		tgTransport := &http.Transport{
			Dial: tgDialer.Dial,
		}
		client.Transport = tgTransport
	}

	bot, err := tgbotapi.NewBotAPIWithClient(TOKEN, client)
	if err != nil {
		fmt.Println(bot)
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		fmt.Println(msg)
		//msg.ReplyToMessageID = update.Message.MessageID
		reply := update.Message.Text
		switch update.Message.Text {
		case "/start", "/help":
			reply = getHelp()
		case "/cat_fact":
			reply = api.CatFact()
		case "/geek_joke":
			reply = api.GeekJoke()
		case "/fox_pic":
			reply = api.FoxPic()
		case "/fact":
			reply = api.Fact()
		case "/xkcd":
			reply = api.Xkcd()
		case "/tech_quote":
			reply = api.TechQuote()
		case "/startup_quote":
			reply = api.StartupQuote()
		default:
			reply = "Unknown command. Send /help to get available commands"
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}
