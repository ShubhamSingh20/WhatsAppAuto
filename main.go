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
const (
	INDCode          = "91"
	settlingPeriod   = 3
	messagePeriod    = 1
	qrCodePeriod     = 20
	contactNumberLen = 10
)

func pressKeyToContinue() {
	color.BgYellow.Println("[*] Press enter to continue ..")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	color.BgLightGreen.Println("[+] You will be prompted with a qrcode that " +
		"you have to scan from your phone ...")

	pressKeyToContinue()

	messageList, err := loadCSV()
	if err != nil {
		pressKeyToContinue()
		os.Exit(1)
	}

	contactNumber := getContactNumber(messageList)
	message := getMessage(messageList)
	wac, err := createNewConnection()
	defer wac.Logout()

	if err != nil {
		pressKeyToContinue()
		os.Exit(1)
	}

	<-time.After(settlingPeriod * time.Second)
	sendBulkMessage(wac, message, contactNumber)

	color.BgBlue.Println("[+] Logging out of Whatsapp")

	pressKeyToContinue()

}
