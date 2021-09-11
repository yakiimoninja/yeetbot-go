package commands

import (
	//"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//ImgSend Commands
func ImgSend(session *discordgo.Session, msg *discordgo.MessageCreate) {

	if msg.Author.ID != botID && strings.HasPrefix(msg.Content, "y.") == true {

		switch msg.Content {
		case "y.fasol":
			/*if err != nil {
			  // handle err
			  }
			  defer file.Close()*/

			session.ChannelMessageSend(msg.ChannelID, "y.fasoliasthkes")
			file, _ := os.Open("images/fasoliastikes.jpg")
			session.ChannelFileSend(msg.ChannelID, "fasoliastikes.jpg", file)

		case "y.wtf":
			file, _ := os.Open("images/wtf.jpg")
			session.ChannelFileSend(msg.ChannelID, "wtf.jpg", file)

		default:
			return
		}

	}
}
