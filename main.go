package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	discord.AddHandler(onMessage)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}

func onMessage(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.Content == "!ping" {
		session.ChannelMessageSend(m.ChannelID, "pong!")
	}
}
