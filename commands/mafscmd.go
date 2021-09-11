package commands

import(
    "strings"
    "unicode"
    "github.com/bwmarrin/discordgo"
    "github.com/Knetic/govaluate"
    "fmt"

)
//MafsSend func
func MafsSend(session *discordgo.Session, msg *discordgo.MessageCreate){
    
    //Checks for prefix and if the bot isnt the author of the message
    if msg.Author.ID != botID && strings.Contains(msg.Content, "y.mafs") == true{
        
        //mafexpression
        mafex := msg.Content

        //Formatting string
        mafex = strings.TrimLeft(mafex, "y.mafs ")
        mafex = spaceStringsTrimmer(mafex)
        
        //String to number calculation
        expression, _ := govaluate.NewEvaluableExpression(mafex)
        parameters := make(map[string]interface{}, 8)
        result, _ := expression.Evaluate(parameters)

        //Interface type to assetion string
        sresult := fmt.Sprintf("Result is: %v", result)
        
        //Result send
        session.ChannelMessageSend(msg.ChannelID, sresult)


    } else{
        return
    }
}

//Function to trim space between math expression
func spaceStringsTrimmer(str string) string {
    var b strings.Builder
    b.Grow(len(str))
    for _, ch := range str {
        if !unicode.IsSpace(ch) {
            b.WriteRune(ch)
        }
    }
    return b.String()
}
