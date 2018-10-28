package models

//File 代表一个有关文件请求操作的结果
type File struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	FileSize float64 `json:fileSize`
	FileType string `json:fileType`
	FileCreatedAt string `json:fileCreatedAt`
	Url      string `json:"url"`
}
