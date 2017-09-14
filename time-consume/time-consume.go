package time_consume

import (
	"time"
	"log"
)

func Run()  {
	defer timeoutCheck("xx slow", time.Now())

	time.Sleep(5 * time.Second)
}

func timeoutCheck(tag string, start time.Time)  {
	dis:= time.Since(start).Seconds()
	if dis > 1 {
		log.Println(tag, dis, " s")
	}
}
