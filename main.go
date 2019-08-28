package main

import (
	"bufio"
	"os"
	"time"

	color "github.com/gookit/color"
)

// run the following commands before running the code
// go get  "github.com/Baozisoftware/qrcode-terminal-go"
// go get "github.com/Rhymen/go-whatsapp"
// go get "github.com/gookit/color"

//INDCode country code
const INDCode = "91"
const settlingPeriod = 3
const messagePeriod = 1
const qrCodePeriod = 10

func pressKeyForExit() {
	color.Yellow.Println("[*] Press enter to close ..")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {

	messageList, err := loadCSV()
	if err != nil {
		pressKeyForExit()
	}

	contactNumber := getContactNumber(messageList)
	message := getMessage(messageList)

	wac, err := createNewConnection()

	if err != nil {
		pressKeyForExit()
		os.Exit(1)
	}

	<-time.After(settlingPeriod * time.Second)
	sendBulkMessage(wac, message, contactNumber)

	color.Blue.Println("[+] Logging out of Whatsapp")
	wac.Logout()

}
