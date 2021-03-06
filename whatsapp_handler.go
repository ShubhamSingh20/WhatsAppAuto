package main

import (
	"fmt"
	"time"

	whatsapp "github.com/Rhymen/go-whatsapp"
	color "github.com/gookit/color"
)

//bold the string..
//https://askubuntu.com/questions/528928/how-to-do-underline-bold-italic-strikethrough-color-background-and-size-i

func sendMessage(wac *whatsapp.Conn, message string, contactNo string) {

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: contactNo + "@s.whatsapp.net",
		},
		Text: message,
	}

	magID, err := wac.Send(msg)

	if err != nil {
		color.BgBlue.Println("[-] Error sending message to ", contactNo, " error => ", err)
	} else {
		color.BgBlue.Println("[+] Message Sent to ", contactNo, " with ID : "+magID)
	}

}

func sendBulkMessage(wac *whatsapp.Conn, message []string, contactNo []string) {

	for i := 0; i < len(message); i++ {
		sendMessage(wac, message[i], contactNo[i])
	}

}

func createNewConnection() (*whatsapp.Conn, error) {
	wac, err := whatsapp.NewConn(qrCodePeriod * time.Second)
	if err != nil {
		color.Error.Println("[-] Error creating connection: ", err)
		return nil, err
	}

	err = oneTimeLogin(wac)
	if err != nil {
		color.Danger.Println("[-] Error creating connection: ", err)
		return nil, err
	}

	return wac, nil

}

func oneTimeLogin(wac *whatsapp.Conn) error {
	qr := make(chan string)

	go func() {
		qrcodeFilePath := saveCurrentQrCode(<-qr)
		color.BgMagenta.Println("[+] QrCode Saved as ", qrcodeFilePath)
		openURLInBrowser(qrcodeFilePath)
	}()

	_, err := wac.Login(qr)

	if err != nil {
		return fmt.Errorf("error during login: %v", err)
	}

	return nil
}
