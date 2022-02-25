package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lucasnevespereira/samplebot-discord/handlers"
	"github.com/lucasnevespereira/samplebot-discord/utils"
)

func init() {
	flag.StringVar(&utils.BotToken, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bot, err := discordgo.New("Bot " + utils.BotToken)
	if err != nil {
		fmt.Println("error creating bot", err)
		return
	}

	bot.AddHandler(handlers.CreateMsg)

	// This bot intents only to receive messages.
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening bot", err)
		return
	}

	fmt.Println("SampleBot is running. Press ctrl-c to abort.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sigs

	bot.Close()

}
