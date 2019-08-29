package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func openURLBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func getContactNumber(m []Message) []string {
	var contactNumber []string

	for i := 1; i < len(m); i++ {
		tempString := INDCode + strings.Replace(m[i].ContactNumber, " ", "", -1)
		contactNumber = append(contactNumber, tempString)
	}

	return contactNumber
}

func getMessage(m []Message) []string {
	var message []string
	for i := 1; i < len(m); i++ {
		tempString := m[i].StudentName + "\n" + "your timing for " +
			m[i].Subject + " class is " + m[i].ClassTime

		message = append(message, tempString)
	}

	return message
}
