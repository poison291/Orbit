package handlers

import (
	"github.com/bwmarrin/discordgo"
)


func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	
	if m.Content == "test" {
		s.ChannelMessageSend(m.ChannelID, "Server Message Event Driven Succesfully	")
	}
}
