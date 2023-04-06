package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id        uint64         `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;notnull"`
	Title     string         `json:"title" gorm:"column:title;type:varchar(200)"`
	Author    string         `json:"author" gorm:"column:author;type:varchar(200)"`
	Desc      string         `json:"desc" gorm:"column:des;type:varchar(200)"`
	Create_at *time.Time     `json:"create_at"`
	Update_at *time.Time     `json:"update_at"`
	Delete_at gorm.DeletedAt `json:"-" gorm:"index"`
	// Delete_at gorm.DeletedAt sama artinya
	//Delete_at *time.Time `json:"-" gorm:"index"`
	// Delete_at gorm.DeletedAt sama artinya
	//tapi dalam pengujian soft delete tidak bekerja apabila
	//saya tidak menggunakan default delete_at dari gorm
}

//untuk mengcostom name table yg di buat
// func (Book) TableName() string {
// 	return "tess"
// }
