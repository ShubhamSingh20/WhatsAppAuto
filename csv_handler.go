package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"

	color "github.com/gookit/color"
)

const csvFilename = "student_schedule.csv"

//Message load detials from csv
type Message struct {
	StudentName   string
	StudentCode   string
	Subject       string
	ClassTime     string
	ContactNumber string
}

func (m *Message) getFieldValues() string {
	return (m.StudentName + ", " +
		m.StudentCode + ", " + m.Subject + ", " +
		m.ClassTime + ", " + m.ContactNumber)

}

var correctColumnRowMessage = Message{
	"student_name", "student_code",
	"subject", "time", "contact_no",
}

func verifyColumns(columnRow *Message) bool {
	return correctColumnRowMessage == *columnRow
}

func loadCSV() ([]Message, error) {
	csvFile, _ := os.Open(csvFilename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var message []Message

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			color.Danger.Println("[-] File Not Found, make sure file name is ", csvFilename)
			return nil, errors.New("File Not Found")
		}

		if len(line[3]) != 0 { //Don't load the row if the schedule/Time column is empty
			message = append(message, Message{
				StudentName:   line[0],
				StudentCode:   line[1],
				Subject:       line[2],
				ClassTime:     line[3],
				ContactNumber: line[4],
			})
		}

	}

	if !verifyColumns(&message[0]) {
		color.Danger.Println("[-] Valid Column names are ",
			correctColumnRowMessage.getFieldValues(), " in that order")

		color.Danger.Println("[-] Your column names were ", message[0].getFieldValues())

		return nil, errors.New("Invalid columns")
	}

	return message, nil

}
