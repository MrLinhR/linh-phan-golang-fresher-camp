package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "Note"
}

func main() {
	dsn := "root:sa123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	//Truy vấn và thêm note vào database
	note4 := Note{Id: 9, Title: "Check", Content: "Create"}
	db.Create(&note4)

	//Truy vấn và lấy ra note có id = 8
	var note Note
	db.Where("id = ?", 8).Find(&note)
	fmt.Println(note)

	//Truy vấn và xoá note có title là "Hello"
	var note1 Note
	db.Where("title = ?", "Hello").Delete(&note1)

	//Truy vấn và cập nhật lại dữ liệu title là "Here" có id = 8
	var note2 Note
	note2.Title = "Here"
	db.Where("id = ?", 8).Updates(&note2)
}
