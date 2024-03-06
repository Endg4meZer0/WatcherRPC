package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/hugolgst/rich-go/client"
	_ "github.com/joho/godotenv/autoload"
)

type ProcessToWatch struct {
	ProcessName    string
	Details        string
	State          string
	LargeImageKey  string
	LargeImageText string
	SmallImageKey  string
	SmallImageText string
	UseTimestamp   bool
	Buttons        []ActivityButtons
}

type ActivityButtons struct {
	Text string
	Url  string
}

func SetActivity(currentHighestIndex *int, processes []ProcessToWatch) {
	lastHighest := *currentHighestIndex

	for i := range len(processes) {
		out, err := exec.Command("tasklist", "/fi", fmt.Sprintf(`IMAGENAME eq %v`, processes[i].ProcessName)).Output()
		if err != nil {
			fmt.Printf("Warning: couldn't use the command \"tasklist\" for the process with the name of %v", processes[i].ProcessName)
			continue
		}
		if string(out[:5]) == "INFO:" {
			continue
		} else {
			*currentHighestIndex = i
			break
		}
	}

	processInfo := processes[*currentHighestIndex]
	newActivity := client.Activity{
		Details:    processInfo.Details,
		State:      processInfo.State,
		LargeImage: processInfo.LargeImageKey,
		LargeText:  processInfo.LargeImageText,
		SmallImage: processInfo.SmallImageText,
		SmallText:  processInfo.SmallImageText,
		Buttons:    []*client.Button{},
	}
	for i := range len(processInfo.Buttons) {
		newActivity.Buttons = append(newActivity.Buttons, &client.Button{
			Label: processInfo.Buttons[i].Text,
			Url:   processInfo.Buttons[i].Url,
		})
	}
	if processInfo.UseTimestamp {
		if lastHighest == *currentHighestIndex {
			return
		}
		now := time.Now()
		newActivity.Timestamps = &client.Timestamps{
			Start: &now,
		}
	}

	err := client.SetActivity(newActivity)

	if err != nil {
		log.Panicln("Error: couldn't set the new activity.")
	}
}

func main() {
	processesListFile, err := os.ReadFile("./processList.json")
	if err != nil {
		log.Panicln("Error: the file of processes to watch does not exist. Please, create processList.json in app folder.")
	}

	var processes []ProcessToWatch

	err = json.Unmarshal(processesListFile, &processes)
	if err != nil {
		log.Panicln("Error: couldn't parse processList.json. Maybe it is not using proper JSON format?")
	}
	currentHighestIndex := len(processes) + 1

	err = client.Login(os.Getenv("CLIENT_ID"))
	fmt.Println("hello world")

	if err != nil {
		log.Panicln("Error while connecting. Check the client id in your .env file")
	}

	SetActivity(&currentHighestIndex, processes)

	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		SetActivity(&currentHighestIndex, processes)
	}
}
