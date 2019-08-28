package main

import (
	"os"
	"time"
)

// run the following commands before running the code

// go get  "github.com/Baozisoftware/qrcode-terminal-go"
// go get "github.com/Rhymen/go-whatsapp"
// go get "github.com/gookit/color"

const settlingPeriod = 3
const messagePeriod = 1

func main() {
	contactNo := []string{"", "", ""} //phone number goes here
	message := []string{              //messages go here
		"This is an automated Message using Go routines",
		"This is an automated Message using Go routines",
		"This is an automated Message using Go routines",
	}
	wac, err := createNewConnection()

	if err != nil {
		os.Exit(1)
	}

	<-time.After(settlingPeriod * time.Second)
	sendBulkMessage(wac, message, contactNo)
	wac.Logout()

}
