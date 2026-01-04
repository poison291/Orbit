package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session, guildID string) {
	if s.State.User == nil {
		fmt.Println("Bot not initialized! Skipped Command registering process")
	}
	_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, &discordgo.ApplicationCommand{
		Name:        "orbit",
		Description: "Root command for Orbit bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "status",
				Description: "Check bot status",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Name:        "welcome",
				Description: "Configure Welcome Messages",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type: discordgo.ApplicationCommandOptionSubCommand,
						Name: "disable",
						Description: "Disable the welcome message",
					},
					{
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Name:        "set",
						Description: "Set the welcome channel",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionChannel,
								Name:        "channel",
								Description: "Select the channel for welcome messages",
								Required:    true,
							},
						},
						
					},
				},
				
				
			},
		},
	})
	if err != nil {
		fmt.Println("Error registering command", err)
	} else {
		fmt.Println("✅ /orbit command succesfully registered")
	}
}
