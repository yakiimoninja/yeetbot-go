package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var emojis = [5]string{"<:rg:809153241421053962>", "<:rl:809153265568448583>", "<:lg:809153253664751697>", "<:hm:809153288086749254>", "<:ck:809153274971684884>"}

//MtnSend function
func MtnSend(session *discordgo.Session, msg *discordgo.MessageCreate) {

	if strings.Contains(msg.Content, "y.shoot") == true {
		//Random seed
		rand.Seed(time.Now().UTC().UnixNano())

		userName := strings.TrimLeft(msg.Content, "y.shoot")
		userName = strings.TrimSpace(userName)
        userName = "@" + userName

		session.ChannelMessageDelete(msg.ChannelID, msg.ID)
		session.ChannelMessageSend(msg.ChannelID, msg.Author.Username+"  "+randomQuakeEmoji()+"  "+ userName)

	}
}

func randomQuakeEmoji() (quakeEmoji string) {

	quakeEmoji = emojis[rand.Intn(5)]
	return quakeEmoji
}
