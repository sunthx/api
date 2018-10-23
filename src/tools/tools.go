package tools

import (
	"models"
	"time"
	"tools/internal/guid"
	"tools/internal/qr"
)

//NewUUID 生成一个UUID
func NewUUID() models.Result {
	UUID := models.Result{
		Content:guid.New(),
		Date:time.Now().String(),
	}

	return UUID
}

//GenerateQRCodeFile 生成二维码文件
func GenerateQRCodeFile(content string) models.File {
	newFile := qr.GenerateQRCodeFile(content)
	newFile.Date = time.Now().String()
	return newFile
}