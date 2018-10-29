package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"models"
	"time"
)

var dbConnectionString = "root:11111111@/core"

// Store File Info
func CreateFileObject(file models.File) (models.File,error){
	db,err := sql.Open("mysql",dbConnectionString)
	if err != nil {
		return models.File{},err
	}
	defer db.Close()

	statement := "insert into files (file_id,file_name,file_type,file_size,file_create_date) values (?,?,?,?,?)"
	stmt,err := db.Prepare(statement)
	if err != nil {
		return models.File{},err
	}

	createdAt,err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST",file.FileCreatedAt)
	if err != nil {
		return models.File{},err
	}

	_,err = stmt.Exec(file.FileId,file.FileName,file.FileType,file.FileSize,createdAt)
	if err != nil {
		return models.File{}, err
	}

	return file,nil
}


