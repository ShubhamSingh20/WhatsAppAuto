package main

import (
	"os"
	"path/filepath"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func saveCurrentQrCode(QrcodeString string) string {
	currentDir, _ := os.Getwd()
	qrcodeFolder := filepath.Join(currentDir, "qrcode")

	if _, err := os.Stat(qrcodeFolder); os.IsNotExist(err) {
		os.Mkdir(qrcodeFolder, os.ModeDir)
	}

	currentTime := time.Now()
	qrcodeFileName := "whatsapp_qrcode_" +
		currentTime.Format("01-02-2006 15-04-05 Monday") + ".png"

	qrcodeFilePath := filepath.Join(qrcodeFolder, qrcodeFileName)

	qrcode.WriteFile(QrcodeString, qrcode.Low, 256, qrcodeFilePath)

	return qrcodeFilePath
}
