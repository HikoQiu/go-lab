package channel

import "fmt"

// Unbuffer channel, wait goroutine to finish
func UnbufferChanWait() {
	wait := make(chan bool)

	go func() {
		fmt.Println("deal some biz ...")

		wait <- true
	}()

	<-wait
}

// Unbuffer channel dead lock
func UnbufferChanDeadlock() {
	c := make(chan bool)

	// 违反 CSP
	c <- true
	<-c

	fmt.Println("Unbuffer DeadLock")
}

// buffer channel, wait all goroutines to finish
// 常用: WaitGroup 处理同样的效果
func BufferChanWaitAll() {
	waits := make(chan bool, 10)
	fmt.Println("len: ", len(waits), "cap: ", cap(waits))

	for i := 0; i < cap(waits); i ++ {
		go func(index  int) {
			fmt.Println("goroutine index - ", index)
			waits<- true
		}(i)
	}

	for i := 0; i < cap(waits); i++ {
		<- waits
	}
}

// 长度为 1 的缓冲 channel
// 对比一下无缓冲 channel
func BufferChaneLen1()  {
	wait := make(chan bool, 1)

	go func() {
		fmt.Println("deal some biz ...")

		wait <- true
	}()

	<-wait
}


