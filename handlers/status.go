package handlers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/poison291/orbit/bot"
)

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second) // round to nearest second
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)
}

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	uptime := time.Since(bot.StartTime)
	fmt.Println("The bot is uptime for ", uptime)

	guild, err := s.Guild(i.GuildID)
	if err != nil {
		fmt.Println("Failed to get guild", err)
	}
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	switch data.Name {
	case "status":
		latency := s.HeartbeatLatency()

		statusMsgembed := &discordgo.MessageEmbed{
			Title: fmt.Sprintf("🔹 **%s - Connection Status**", guild.Name),
			Description: fmt.Sprintf(
				"**Heartbeat Latency:** `%d ms`\n**Uptime:** `%s`",
				latency.Milliseconds(),
				formatDuration(uptime),
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
								Label:    "Run /help",
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
