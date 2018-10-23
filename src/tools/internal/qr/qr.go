package qr

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode/qr"
	file2 "tools/internal/file"
	"models"
)

//GenerateQRCodeFile 根据内容生成二维码图片
func GenerateQRCodeFile(content string) models.File {
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)
	newQrCodeFile := file2.CreateEmptyFile("png")
	file, _ := os.Create(file2.GetFullFilePathByFileName(newQrCodeFile.FileName))
	defer file.Close()

	png.Encode(file, qrCode)
	return newQrCodeFile
}
