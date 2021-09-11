package commands

import(
    "fmt"
    "strings"
    "github.com/bwmarrin/discordgo"
)

const botID string = "517781262623178767"

//MsgSend Commands
func MsgSend(session *discordgo.Session, msg *discordgo.MessageCreate){
    
    if msg.Author.ID == botID{
        return
    }

    if msg.Content == "y.ping"{
        session.ChannelMessageSend(msg.ChannelID, "pong")
    } 

    // TODO TTS
    if msg.Content == "23"{
        session.ChannelMessageSend(msg.ChannelID, "Steef is 23!")
    }
    
    if msg.Content == "bb"{
    }
    //TODO TTS
    if msg.Content == "y.greku"{
    }

//------------------------------Mentions--------------------------------------//

    if strings.Contains(msg.Content, "plek"){
        session.ChannelMessageSend(msg.ChannelID, "You got shot " + msg.Author.Mention())
        fmt.Println(session.RequestGuildMembers("566021217891516457", "", 0, true))
        

    }
}
