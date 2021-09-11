package bot

import (
	"../commands"
	"../config"
	"github.com/bwmarrin/discordgo"
)

//Start func
func Start() {

	//Making Bot instance and error checking
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		println(err.Error())
		return
	}


	//Bot Event Handlers
	goBot.AddHandler(commands.MsgSend)
	goBot.AddHandler(commands.ImgSend)
	goBot.AddHandler(commands.MafsSend)
	goBot.AddHandler(commands.MtnSend)

    //StatusHandler
    goBot.AddHandler(config.OnReady)
	//Running Bot and error checking
	err = goBot.Open()
	if err != nil {
		println("Error establishing connection,", err.Error())
		return
	}

	println("\nBot is running!")
}
