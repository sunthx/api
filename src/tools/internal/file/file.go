package file

import (
	"common"
	"io/ioutil"
	"models"
	"os"
	"tools/internal/guid"
					)

var fileStoragePath = "./public/files"

//CreateEmptyFile 通过指定文件扩展名生成一个未使用的文件对象
func CreateEmptyFile(ext string) models.File {
	fileID := guid.New()

	return models.File{
		FileId:   fileID,
		FileName: fileID + "." + ext,
	}
}

//GetFullFilePathByFileName 通过文件名获取文件的实际路径（不管文件存在还是不存在）
func GetFullFilePathByFileName(fileName string) string {
	return fileStoragePath + "/" + fileName
}

//SaveFile 保存文件到硬盘
func SaveFile(fileName string, data []byte) (bool, error) {
	exist, err := common.PathExist(fileStoragePath)
	if err == nil && !exist {
		os.Mkdir(fileStoragePath, os.ModePerm)
	}

	filePath := fileStoragePath + "/" + fileName
	newFile, _ := os.Create(filePath)
	defer newFile.Close()

	ioutil.WriteFile(fileStoragePath+"/"+fileName, data, os.ModePerm)
	return true, nil
}