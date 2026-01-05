package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func HelpHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Orbit Help",
		Description: "Here are all the commands available:",
		Color:       0x1ABC9C,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "/status", Value: "Check the bot's status", Inline: false},
			{Name: "/welcome-set", Value: "Set the welcome channel", Inline: false},
			{Name: "/welcome-disable", Value: "Disable welcome messages", Inline: false},
		},
	}

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()
		if data.Name == "help"  {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
		}

	case discordgo.InteractionMessageComponent:
		if i.MessageComponentData().CustomID == "orbit_help" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			})
		}
	}
}
