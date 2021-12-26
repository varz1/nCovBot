package bot

import (
	"github.com/varz1/nCovBot/channel"
	"log"
)

func sender() {
	for msg := range channel.MessageChannel {
		_, err := botAPI.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func TimerSender() {

}