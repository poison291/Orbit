package main

import (
	"fmt"
	"log"
	"os"
	"github.com/poison291/orbit/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	//Bot session created
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	sess.AddHandler(handlers.MessageHandler)

	// Bot connection => with discord
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()
	fmt.Println("The Orbit is online!")
	select {}
}
