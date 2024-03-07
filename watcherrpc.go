package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	wrpcTypes "watcherrpc/app/types"
	wrpcUtils "watcherrpc/app/utils"

	"github.com/hugolgst/rich-go/client"
	_ "github.com/joho/godotenv/autoload"
)

func SetActivity(currentHighestIndex *int, processes []wrpcTypes.ProcessToWatch) {
	lastHighest := *currentHighestIndex

	wrpcUtils.ProcessCheck(currentHighestIndex, processes)

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
	if os.Args[0] == "--create-start-link" {
		err := wrpcUtils.CreateStartUpFile()
		if err != nil {
			log.Println("Warning: couldn't create a start up file. You might want to try and make it yourself, then delete the hidden .firstlaunch file.")
		} else {
			os.Remove("./.firstlaunch")
			fmt.Println("Done! WatcherRPC will now start with your system!")
		}
	}

	_, err := os.ReadFile("./.firstlaunch")
	if err == nil && os.Args[0] != "--ignore-first-launch" {
		fmt.Print("Seems like you've launched WatcherRPC for the first time.\nDo you want to launch it on system start up? (Y/n) ")
		var output string
		fmt.Scanln(&output)
		if strings.ToLower(output) != "n" {
			err = wrpcUtils.CreateStartUpFile()
			if err != nil {
				log.Println("Warning: couldn't create a start up file. You might want to try and make it yourself, then delete the hidden .firstlaunch file.")
			} else {
				os.Remove("./.firstlaunch")
				fmt.Println("Done! WatcherRPC will now start with your system!")
			}
		} else {
			os.Remove("./.firstlaunch")
			fmt.Println("Okay! If you ever want to do it tho, use --create-start-link flag.")
		}
	}

	processesListFile, err := os.ReadFile("./processList.json")
	if err != nil {
		log.Panicln("Error: the file of processes to watch does not exist. Please, create processList.json in app folder.")
	}

	var processes []wrpcTypes.ProcessToWatch

	err = json.Unmarshal(processesListFile, &processes)
	if err != nil {
		log.Panicln("Error: couldn't parse processList.json. Maybe it is not using proper JSON format?")
	}
	currentHighestIndex := len(processes) + 1

	err = client.Login(os.Getenv("CLIENT_ID"))
	fmt.Println("Hello world!")

	if err != nil {
		log.Panicln("Error while connecting. Check the client id in your .env file")
	}
	fmt.Println("Connected! You can now minimize and ignore this window.")

	SetActivity(&currentHighestIndex, processes)

	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		SetActivity(&currentHighestIndex, processes)
	}
}
