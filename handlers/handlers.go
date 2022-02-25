package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lucasnevespereira/samplebot-discord/utils"
)

func CreateMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from the bot himself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "*test" {

		// Send msg
		_, err := s.ChannelMessageSend(m.ChannelID, utils.TestMsg)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Unknown command")
	}

}
