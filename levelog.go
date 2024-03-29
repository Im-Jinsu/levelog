package levelog

import (
	"fmt"
	"log"
	"os"

	"github.com/im-jinsu/levelog/utils/discord"
)

// Info : Service status
func Info(msg ...interface{}) {
	log.Output(2, fmt.Sprintf("[Info] %s",
		fmt.Sprintln(msg...)))
}

// Error : Default error logging
func Error(msg ...interface{}) {
	log.Output(2, fmt.Sprintf("[ERR] %s",
		fmt.Sprintln(msg...)))
}

// ErrorAndSend : Error logging and send notice
func ErrorAndSend(msg ...interface{}) {
	msgStr := fmt.Sprintf("[ERR] %s",
		fmt.Sprintln(msg...))
	log.Output(2, msgStr)
	// send
	if !isDevelop() {
		s := new(discord.MsgBlock)
		s.Content = msgStr
		s.Username = "Logger Notice"
		if err := s.Send(); err != nil {
			log.Println(err)
		}
	}
}

// Fatal : Very severe error, the program terminated
func Fatal(msg ...interface{}) {
	log.Output(2, fmt.Sprintf("**Fatal** %s",
		fmt.Sprint(msg...)))
	os.Exit(1)
}

// Debug : Logging in develop environment only
func Debug(msg ...interface{}) {
	if isDevelop() {
		log.Output(2, fmt.Sprintf("[Debug] %s",
			fmt.Sprintln(msg...)))
	}
}

func isDevelop() bool {
	if isdev := os.Getenv("LEVELOG_ISDEV"); isdev == "false" {
		return false
	}
	return true
}
