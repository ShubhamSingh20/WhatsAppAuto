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
	INDCode        = "91"
	settlingPeriod = 3
	messagePeriod  = 1
	qrCodePeriod   = 20
)

func pressKeyToContinue() {
	color.Yellow.Println("[*] Press enter to continue ..")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	color.Blue.Println("[+] You will be prompted with a qrcode that " +
		"you have to scan from your phone ...")

	pressKeyToContinue()

	messageList, err := loadCSV()
	if err != nil {
		pressKeyToContinue()
	}

	contactNumber := getContactNumber(messageList)
	message := getMessage(messageList)

	wac, err := createNewConnection()

	if err != nil {
		pressKeyToContinue()
		os.Exit(1)
	}

	<-time.After(settlingPeriod * time.Second)
	sendBulkMessage(wac, message, contactNumber)

	color.Blue.Println("[+] Logging out of Whatsapp")
	pressKeyToContinue()
	wac.Logout()

}
