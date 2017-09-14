package interf

import (
	"time"
	"fmt"
	"reflect"
)

type ITask interface {
	Serialize(string) ([]byte, error)
	UnSerialize([]byte) (string, error)
	Run()
}

// 订单任务
type DealTask struct{}

func (t DealTask) Serialize(in string) ([]byte, error) {
	return []byte(in), nil
}

func (t DealTask) UnSerialize(out []byte) (string, error) {
	return string(out), nil
}

func (t DealTask) Run()  {
	for {
		// 1.1 假设从 Channel A 获取到数据
		// inBytes <- A
		inBytes := []byte("hello world")

		// 2.1 反序列化
		content, err:=t.UnSerialize(inBytes)
		if err != nil {
			// error log
			continue
		}

		// 3.1 假设处理数据
		// 4.1 序列化
		_, err = t.Serialize(content)

		// 5.1 放到另外的 Channel B
		// B <- byteCnt
		time.Sleep(time.Second * 2)

		fmt.Println(fmt.Sprintf("Task ** %s ** deal one msg.",  reflect.TypeOf(t).Name()))
	}
}

// 理赔任务
type ClaimTask struct{}

func (t ClaimTask) Serialize(in string) ([]byte, error) {
	return []byte(in), nil
}

func (t ClaimTask) UnSerialize(out []byte) (string, error) {
	return string(out), nil
}

func (t ClaimTask) Run()  {
	for {
		fmt.Println(fmt.Sprintf("Task [ %s ] deal one msg.",  reflect.TypeOf(t).Name()))
		time.Sleep(time.Second * 1)
	}
}

