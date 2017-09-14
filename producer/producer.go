package producer

import (
	"fmt"
	"time"
	"sync"
)

const (
	QUEUE_LEN = 100
)

const (
	// 每次 const 出现时，都会让 iota 初始化为 0
	ROLE_PRODUCER = iota // 0
	ROLE_CONSUMER        // 1
	_
	ROLE_UNKNOWN         // 3
)

// 作用域: 当前包
var ins *Svr

func init() {
	fmt.Println("Package producer init().")
}

var lock sync.Mutex // 开箱即用
// fmt.Println(reflect.TypeOf(lock).String())

type Svr struct {
	q chan string
}

func NewSvr() *Svr {
	defer lock.Unlock()

	lock.Lock()
	if ins != nil {
		return ins
	}

	ins := new(Svr)
	ins.q = make(chan string, QUEUE_LEN)
	return ins
}

// 消费者
func (t *Svr) produce() {
	for {
		t.q <- "RAW DATA - json"
		time.Sleep(time.Second)
	}
}

// 生产者
func (t *Svr) consume() {
	for {
		tmp := <-t.q
		fmt.Println(tmp)
	}
}

// 启动 num 个消费者
func (t *Svr) runProducers(num int) {
	for i := 0; i < num; i++ {
		go t.produce()
	}
}

// 启动 num 个生产者
func (t *Svr) runConsumer(num int) {
	for i := 0; i < num; i++ {
		go t.consume()
	}
}

func (t *Svr ) Run() {
	go t.runProducers(10)
	go t.runConsumer(10)
	select {}
}

