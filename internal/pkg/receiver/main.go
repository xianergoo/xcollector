package receiver

import (
	"fmt"
	"time"
)

var transceivers []TTransceiver

func Start() {
	//enable logging
	//InitReceiverList
	transceivers := make([]TTransceiver, 0)
	input := make(chan string)

	engine := TEngine{
		Host: "192.168.4.179",
		Port: 4660,
	}

	transceivers = append(transceivers, TTransceiver{
		TEngine: engine,
		Enabled: true,
	})

	for _, trans := range transceivers {
		if !trans.Enabled {
			continue
		}
		logger := newLoggerForTransceiver(&trans)
		trans.Logger = logger
		trans.TEngine.Logger = logger
		err := trans.Connect()
		if err != nil {
			fmt.Println("Error connecting:", err)
			continue
		}
		fmt.Println("test")
		go trans.Execute()
	}

	for {
		select {
		case data := <-input:
			// 在这里处理采集到的数据
			// ...

			// 模拟处理时间
			time.Sleep(time.Second)
			fmt.Println(data)
		case <-time.After(time.Second * 2): // 超时处理，防止阻塞主协程
			fmt.Println("Timeout, exiting CollectData")
			return
		default:
			continue
		}
	}

	// // 示例：查找命令名称
	// fmt.Println(consts.LookupCmdName(consts.GetVersionCmd)) // 输出 "Get Version"

	// // 示例：打印所有命令名称
	// for cmdCode := consts.UnknownCmd; cmdCode <= consts.SetGroupNumberCmd; cmdCode++ {
	// 	fmt.Printf("%d: %s\n", cmdCode, consts.LookupCmdName(cmdCode))
	// }
}
