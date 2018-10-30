package tools

import (
	"models"
	"time"
	"tools/internal/guid"
	"tools/internal/qr"
	"tools/internal/file"
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
	//newFile.Date = time.Now().String()
	return newFile
}

//保存文件到本地磁盘
func SaveFileToDisk(fileName string, data []byte) (bool, error){
	return file.SaveFile(fileName,data)
}

//GetFileContent 获取文件内容
func GetFileContent(fileName string) ([]byte,error){
	return file.GetFileContent(fileName)
}