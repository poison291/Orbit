package handlers

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/bwmarrin/discordgo"
)

const welcomeChannelId string = "1455945306251460669"

func Welcome(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	rand.Seed(time.Now().UnixNano())
	randomColor := rand.Intn(0xFFFFFF + 1)

	guild, err := s.Guild(m.GuildID) // getting server data

	if err != nil {
		fmt.Println("State guild error:", err)
		return
	}

	members, err := s.GuildMembers(m.GuildID, "", 1000)

	embed := &discordgo.MessageEmbed{
		Title:       "👋 Welcome!",
		Description: "Hello " + m.User.Mention() + " to " + guild.Name + "!",
		Color:       randomColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.User.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("You are member #%d", len(members)),
		},
	}
	s.ChannelMessageSend(welcomeChannelId, "Welcome "+m.User.Mention()+" to "+guild.Name+"!")
	s.ChannelMessageSendEmbed(welcomeChannelId, embed)

}
