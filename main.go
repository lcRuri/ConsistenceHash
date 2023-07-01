package main

import (
	"fmt"
	"hashring/consistenhash"
)

func main() {
	ring := consistenhash.New(2)

	//设置三个地址
	ring.AddNodes([]string{
		"192.168.5.2",
		"192.168.2.3",
		"192.168.3.4",
	})

	key1 := "6666666666"
	key2 := "asasaasasas"
	key3 := "qwert"

	addr1 := ring.GetNode(key1)
	fmt.Println(key1 + " is map to addr:" + addr1)

	addr2 := ring.GetNode(key2)
	fmt.Println(key2 + "key is map to addr:" + addr2)

	addr3 := ring.GetNode(key3)
	fmt.Println(key3 + "is map to addr:" + addr3)
}
