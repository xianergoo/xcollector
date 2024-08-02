package receiver

import (
	"fmt"
)

var transceivers []TTransceiver

func Start() {
	//enable logging

	//InitReceiverList
	transceivers := make([]TTransceiver, 0)

	transceivers = append(transceivers, TTransceiver{
		TEngine: TEngine{
			Host: "192.168.4.179",
			Port: 4660,
		},
		Enabled: false,
	})

	for _, trans := range transceivers {
		if !trans.Enabled {
			continue
		}
		err := trans.Connect()
		if err != nil {
			fmt.Println("Error connecting:", err)
			continue
		}
		go trans.Execute()
	}

	// // 示例：查找命令名称
	// fmt.Println(consts.LookupCmdName(consts.GetVersionCmd)) // 输出 "Get Version"

	// // 示例：打印所有命令名称
	// for cmdCode := consts.UnknownCmd; cmdCode <= consts.SetGroupNumberCmd; cmdCode++ {
	// 	fmt.Printf("%d: %s\n", cmdCode, consts.LookupCmdName(cmdCode))
	// }
}
