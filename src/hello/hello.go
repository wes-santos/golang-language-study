package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const timesToMonitoring = 1
const delayBetweenMonitorings = 5 * time.Second

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

func showIntroduction() {
	name := "Weslley"
	fmt.Println("Hello,", name+"!")
}

func showOptions() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
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

func checkWebsite(site string) error {
	response, err := http.Get(site)
	if err != nil {
		return err
	}
	switch response.StatusCode {
	case http.StatusOK:
		fmt.Println("The website", site, "is OK.")
		saveLogsToFile(site, true)
	case http.StatusNotFound:
		fmt.Println("The website", site, "is returning a 404 status.")
		saveLogsToFile(site, false)
	default:
		fmt.Println("Returned status by website not mapped. Value returned:", response.StatusCode)
		saveLogsToFile(site, false)
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

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func saveLogsToFile(site string, status bool) error {
	file, err := os.OpenFile(
		"src/logs.txt",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()

	if err != nil {
		return err
	}

	timeNow := time.Now().Format("2006-01-02 15:04:05")

	file.WriteString(
		fmt.Sprintf(
			"%s - %s - online: %s\n",
			timeNow,
			site,
			strconv.FormatBool(status),
		),
	)

	return nil
}
