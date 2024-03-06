package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hugolgst/rich-go/client"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := client.Login(os.Getenv("CLIENT_ID"))
	fmt.Println("hello world")

	if err != nil {
		log.Fatal("Error while connecting. Check the client id in your .env file")
	}

	err = client.SetActivity(client.Activity{
		State:   "TEST!!!!",
		Details: "go test",
	})

	if err != nil {
		log.Fatal("Error while setting an activity.")
	}

	fmt.Scanln()
}
