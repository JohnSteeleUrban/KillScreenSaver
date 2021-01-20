package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()

	if err != nil {
		log.Fatal(err)
	}

	timeArray := os.Args[1:]
	fmt.Println(timeArray)

	if len(timeArray) != 1 {
		log.Fatal("too many arguments")
	}

	waitTime, err := strconv.Atoi(timeArray[0])

	if err != nil {
		log.Fatal(err)
	}

	// For linux, need to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	for i := 0; i < waitTime; i++ {
		fmt.Printf("%d minutes\n", i)
		fmt.Printf("out of: %d\n", waitTime)

		for j := 0; j < 60; j++ {
			fmt.Printf("%d:%d\n", i, j)
			time.Sleep(1 * time.Second)
		}

		kb.SetKeys(keybd_event.VK_B, keybd_event.VK_L, keybd_event.VK_E, keybd_event.VK_ENTER)
		err = kb.Launching()
		if err != nil {
			log.Fatal(err)
		}
	}
}
