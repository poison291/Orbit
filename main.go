package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/poison291/orbit/handlers"
	"github.com/poison291/orbit/bot"
)

func main(){
	bot.StartTime = time.Now()
	
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	//Bot session created
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent | discordgo.IntentsGuildMembers

	// Handlers Stuff
	sess.AddHandler(handlers.MessageHandler)
	sess.AddHandler(handlers.Welcome)
	sess.AddHandler(handlers.InteractionHandler)     
	sess.AddHandler(handlers.HelpHandler)     

	// Bot connection => with discord
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	handlers.RegisterCommands(sess, "1421139951231438948") // my server
	handlers.RegisterCommands(sess, "1002933481246052412") // rexxor server
	
	// handlers.DeleteGuildCommands(sess, "1421139951231438948") // my server
	// handlers.DeleteGuildCommands(sess, "1002933481246052412") // rexxor server
	
	defer sess.Close()
	fmt.Println("The Orbit is online!")
	select {}
}
