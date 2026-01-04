package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.Guild(i.GuildID)
	if err != nil {
		fmt.Println("Failed to get guild", err)
	}
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	if data.Name != "orbit" {
		fmt.Println("The command must start from /orbit")
		return
	}

	switch data.Options[0].Name {
	case "status":
		latency := s.HeartbeatLatency()

		statusMsgembed := &discordgo.MessageEmbed{
			Title: fmt.Sprintf("🔹 **%s - Connection Status**", guild.Name),
			Description: fmt.Sprintf(
				"**Heartbeat Latency:** `%d ms`",
				latency.Milliseconds(),
			),
			Color: 0x1ABC9C,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: guild.IconURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Need help?",
					Value:  "Use `/orbit help` to see all commands",
					Inline: false,
				},
			},
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{statusMsgembed},
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label:    "Run /orbit help",
								Style:    discordgo.PrimaryButton,
								CustomID: "orbit_help",
							},
						},
					},
				},
			},
		})
	}

}
