package models

//File 代表一个有关文件请求操作的结果
type File struct {
	Result
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	Url      string `json:"url"`
}
