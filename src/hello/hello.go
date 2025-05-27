package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func checkWebsite(site string) error {
	response, err := http.Get(site)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println("The website", site, "is OK.")
	case http.StatusNotFound:
		fmt.Println("The website", site, "is returning a 404 status.")
	default:
		fmt.Println("Returned status by website not mapped. Value returned:", response.StatusCode)
	}
	return nil
}

func readWebsitesFromTextFile() ([]string, error) {
	var websites []string

	file, err := os.Open("src/hello/websites.txt")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)

		if err == io.EOF {
			fmt.Println("File readed.")
			break
		}
	}

	return websites, nil
}

func startMonitoring() {
	fmt.Println("Monitoring")

	sites, err := readWebsitesFromTextFile()
	if err != nil {
		fmt.Printf("Error opening websites.txt file: %v", err)
	}

	for i := 0; i < timesToMonitoring; i++ {
		for _, site := range sites {
			err := checkWebsite(site)
			if err != nil {
				fmt.Printf("Error requesting site %s. Error: %v ", site, err)
			}
		}

		if i > 0 {
			time.Sleep(delayBetweenMonitorings)
		}
	}
}

func main() {
	showIntroduction()
	showOptions()

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
