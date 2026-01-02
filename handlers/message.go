package handlers

import (
	"strings"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const prefix = "/orbit"

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "Server Message Event Driven Succesfully	")
	}
	
	
	args := strings.Fields(m.Content[len(prefix):])
	if len(args) == 0{
		return
	}
	
	command := strings.ToLower(args[0])
	
	switch command {
		case "hey":
		s.ChannelMessageSend(m.ChannelID, "switch case passed!")
	}

}
