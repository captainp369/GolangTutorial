package main

import (
	"fmt"
	"time"
)

func finder1(mine [5]string) {
	for _, ore := range mine {
		if ore == "ore" {
			fmt.Println("finder 1 found ore")
		}
	}
}

func finder2(mine [5]string) {
	for _, ore := range mine {
		if ore == "ore" {
			fmt.Println("finder 2 found ore")
		}
	}
}

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	// go finder1(theMine)
	// go finder2(theMine)
	oreChannel := make(chan string)
	// oreChannel1 := make(chan string)
	// oreChannel2 := make(chan string)
	minedOreChan := make(chan string)

	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(theMine)

	// go func(mine1 [5]string) {
	// 	for _, item := range mine1 {
	// 		if item == "ore" {
	// 			oreChannel1 <- item
	// 		}
	// 	}
	// }(theMine)

	// go func(mine2 [5]string) {
	// 	for _, item := range mine2 {
	// 		if item == "ore" {
	// 			oreChannel2 <- item
	// 		}
	// 	}
	// }(theMine)

	go func() {
		//for i := 0; i < 3; i++ {
		for foundOre := range oreChannel {
			//foundOre := <-oreChannel
			// foundOre1 := <-oreChannel1
			// foundOre2 := <-oreChannel2
			fmt.Println("from finder: ", foundOre)
			// fmt.Println("from finder1: ", foundOre1)
			// fmt.Println("from finder2: ", foundOre2)
			minedOreChan <- "minedOre"
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan
			fmt.Println("from miner:", minedOre)
			fmt.Println("from smelter: ore is smelted")
		}
	}()

	<-time.After(time.Second * 5)

	fmt.Println("----------")
	bufferedChan := make(chan string, 3)
	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()
	//<-time.After(time.Second * 2)
	doneChan := make(chan string)
	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
		forthRead := <-bufferedChan
		fmt.Println(forthRead)
		doneChan <- "im done"
	}()
	//<-doneChan
	//<-time.After(time.Second * 2)
	fmt.Println("--------end")
	myChan := make(chan string)

	go func() {
		myChan <- "Message!"
	}()

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}
	<-time.After(time.Second * 1)
	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

}
