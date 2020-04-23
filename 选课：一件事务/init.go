package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}

	if !db.HasTable(&Choose{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Choose{}).Error; err != nil {
			panic(err)
		}
	}
}

type Class struct {//原有的表，里面存可选课的所有信息
	Dacourse_num string `gorm:"type:varchar(256);not null;"`
	Course_num   string `gorm:"type:varchar(256);not null;"`
	Day          string `gorm:"type:varchar(256);not null;"`
	Lesson       string `gorm:"type:varchar(256);not null;"`
	Course       string `gorm:"type:varchar(256);not null;"`
	Teacher      string `gorm:"type:varchar(256);not null;"`
	Classroom    string `gorm:"type:varchar(256);not null;"`
	Rewweek      string `gorm:"type:varchar(256);not null;"`
	Type         string `gorm:"type:varchar(256);not null;"`
}

type Choose struct {//该表中day和lesson仅用于计算时间冲突
	Stunum    int    `gorm:"type:int(255);not null;primary_key"`
	Username   string `gorm:"type:varchar(256);not null;"`
	Date       string `gorm:"type:varchar(1024);not null;"`//A04192A1110020009
	Timeday    string `gorm:"type:varchar(1024);not null;"`//1 # 3 # 5
	Timelesson string `gorm:"type:varchar(1024);not null;"`//3-4 # 7-8 # 3-4
}