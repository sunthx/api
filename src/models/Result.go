package models

//Result 请求结果
type Result struct {
	Date    string `json:"date"`
	Content interface{} `json:"content"`
}