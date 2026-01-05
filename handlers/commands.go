package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session, guildID string) {
	if s.State.User == nil {
		fmt.Println("Bot not initialized! Skipped Command registering process")
		return
	}
	server, err := s.Guild(guildID)
	if err != nil {
		fmt.Println("Error getting guild Data", err)
	}
	s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
		Name:        "status",
		Description: "Check bot status",
	})

	s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Show available commands",
	})

	s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
		Name:        "welcome-set",
		Description: "Set the welcome channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "Select the welcome channel",
				Required:    true,
			},
		},
	})

	s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
		Name:        "welcome-disable",
		Description: "Disable welcome messages",
	})
	if err != nil {
		fmt.Println("Error registering command", err)
	} else {
		fmt.Println("✅ /orbit command succesfully registered", server.Name)
	}
}
