package config 

import(
    "github.com/bwmarrin/discordgo"
    "time"
)

//OnReady function
func OnReady (session *discordgo.Session, event *discordgo.Ready){

    var songs = [15]string{"Dunkelheit", "Jesu død", "Beholding the Daughters of the Firmament", "Decrepitude I",
    "Decrepitude II", "Circumambulation of the Transcendental Pillar of Singularity",   
    "Feeble Screams From Forests Unknown", "Ea, Lord of the Depths", "Spell of Destruction", 
    "Channelling the Power of Souls Into a New God", "War", "The Crying Orc", "A Lost Forgotten Sad Spirit",
    "My Journey to the Stars", "Dungeons of Darkness",
    }


    for i := 0; i < 15; i++{

        session.UpdateListeningStatus(songs[i])
        
        time.Sleep(300 *time.Second)
        
        if i == 14 {
            i = 0
        }
        
    }
}
