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

//func UploadFileHandler(writer http.ResponseWriter, request *http.Request){
//	request.ParseMultipartForm(1024*1024*1024)
//	file := request.MultipartForm.File["uploaded"][0]
//
//	fileReader,err := file.Open()
//	if err != nil {
//		return
//	}
//
//	data,err := ioutil.ReadAll(fileReader)
//	if err != nil {
//		return
//	}
//
//	fileId :=
//	fileExt := path.Ext(file.Filename)
//	fileName := fileId + fileExt
//	result,err := common.SaveFile(fileName,data)
//	if err != nil || !result{
//		return
//	}
//
//	currentFile := models.File{
//		FileId:fileId,
//		FileName:fileName,
//		Url: request.Host + "/static/" + fileId,
//	}
//
//	currentFile.Base.Date =time.Now().String()
//
//	value,err := json.Marshal(&currentFile)
//	if err != nil{
//		return
//	}
//
//	writer.Header().Set("Content-Type","application/json")
//	writer.Header().Set("Access-Control-Allow-Origin", "*")
//	writer.Write(value)
//}
