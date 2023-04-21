package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/urfave/cli"
)

func start(c *cli.Context) {
	time.Sleep(3 * time.Second)
	err := envInit(c)
	if err != nil {
		return
	}

	//初始化
	// ms.drawInit(masterSQLDriver, slaveSQLDriver)
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, getErr := bot.GetUpdatesChan(u)
	if getErr != nil {
		return
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

	//graceful shutdown
	{
		shutdowObserver := make(chan os.Signal, 1)
		signal.Notify(shutdowObserver, syscall.SIGINT, syscall.SIGTERM)
		// Wait close signal here
		s := <-shutdowObserver
		receiveCloseTime := time.Now()
		logger.Debug("Receive  signal:", s, receiveCloseTime.Unix())
		//stop grpc
		//停止當前 queue

		//final chance
		time.Sleep(5 * time.Second)

		logger.Info("final shut down", time.Now().Unix(), "spent time (contain 5s sleep):", time.Since(receiveCloseTime).String())
	}
}
