package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	color "github.com/gookit/color"
)

func openURLInBrowser(url string) {
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
		color.Danger.Println("[-] Error occured while opening file")
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

func stringToBold(tempStr string) string {
	//fmt.Println(fmt.Sprintf("%+q", tempStr))
	return tempStr
}

func getMessage(m []Message) []string {
	var message []string
	for i := 1; i < len(m); i++ {
		tempString := m[i].StudentName + "\n" + "your timing for " +
			m[i].Subject + " class is " + m[i].ClassTime

		tempString = stringToBold(tempString)
		message = append(message, tempString)
	}

	return message
}
