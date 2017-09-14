package go_select

import (
	"time"
	"log"
)

var exit chan bool = make(chan bool)

func Run() {

	go func() {
		select {
		case tmpExit := <-exit:
			log.Fatalln("Receive exit chan.", tmpExit)
		}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		exit <- true
	}()

	select {}
}