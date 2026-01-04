package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	
	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "Server Message Event Driven Successfully")
	}
	
	if m.Content == "!wtest" {
		fmt.Println(m.ChannelID)
	}
}