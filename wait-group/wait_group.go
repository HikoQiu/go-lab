package wait_group

import (
	"sync"
	"fmt"
)

func WaitGoroutineFinish()  {
	wg:=sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("do some biz ...")
		wg.Done()
	}()

	wg.Wait()
}
