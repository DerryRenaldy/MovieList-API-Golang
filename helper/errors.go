package helper

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func PrintError(err error) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		return
	}
}
