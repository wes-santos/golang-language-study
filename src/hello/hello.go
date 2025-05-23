package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const timesToMonitoring = 1
const delayBetweenMonitorings = 5 * time.Second

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func showIntroduction() {
	name := "Weslley"
	fmt.Println("Hello,", name+"!")
}

func showOptions() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
}

func checkWebsite(site string) {
	response, _ := http.Get(site)
	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println("The website", site, "is OK.")
	case http.StatusNotFound:
		fmt.Println("The website", site, "is returning a 404 status.")
	default:
		fmt.Println("Returned status by website not mapped. Value eturned:", response.StatusCode)
	}
}

func startMonitoring() {
	fmt.Println("Monitoring")
	sites := [2]string{"https://httpbin.org/status/404", "https://alura.com.br"}

	for i := 0; i < timesToMonitoring; i++ {
		for _, site := range sites {
			checkWebsite(site)
		}

		if i > 0 {
			time.Sleep(delayBetweenMonitorings)
		}
	}
}

func main() {
	showIntroduction()
	showOptions()

	// if command == 1 {
	// 	fmt.Println("Monitoring...")
	// } else if command == 2 {
	// 	fmt.Println("Logging...")
	// } else if command == 0 {
	// 	fmt.Println("Exiting program. Bye.")
	// } else {
	// 	fmt.Println("Command unrecognized.")
	// }

	command := readCommand()

	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Logging...")
	case 3:
		fmt.Println("Exiting program. Bye.")
		os.Exit(0)
	default:
		fmt.Println("Command unrecognized.")
		os.Exit(-1)
	}
}
