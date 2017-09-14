package main

import (
	"go-lab/go-defer"
	"fmt"
)

func main() {

	// select
	//go_select.Run();
	go_defer.Run()

	// slice
	a := make([]int, 0, 9)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	fmt.Println(a[0:3]) // [i, j)

	//  interface{}
	var itf interface{}
	var tmpNum int64 = 90
	itf = tmpNum
	v, ok := itf.(int64)
	if itf.(int64) == tmpNum {
		fmt.Println("Equal")
	}
	fmt.Println(ok, v)




	// time consume
	// time_consume.Run()

	// pprof, http://localhost:6060/debug/pprof
	//go func() {
	//	http.ListenAndServe("localhost:6060", nil)
	//}()

	//producer & consumer
	//svr := producer.NewSvr()
	//svr.Run()

	// unbuffer channel, wait goroutine to finish
	//channel.UnbufferChanWait()
	// unbuffer channel, deadlock
	//channel.UnbufferChanDeadlock()

	// buffer channel, wait all goroutines to finish
	//channel.BufferChanWaitAll()

	//channel.BufferChaneLen1()
	//sync.WaitGroup{}

	// WaitGroup
	//wait_group.WaitGoroutineFinish()

	// 协程调度 - 1.5 前阻塞
	//cooperative.BlockByFor()

	// interface 接口
	//svr := interf.NewAsyncTaskSvr()
	//svr.AddTask(new(interf.DealTask))
	//svr.AddTask(new(interf.ClaimTask))
	//svr.Run()
}

