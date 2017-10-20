package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"crud_2/models"

)
var DB *gorm.DB

func init(){
	var err error
	DB, err = gorm.Open("mysql", "root:root@/ptk?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Println(err.Error)
	}else{
		log.Println("koneksi berhasil")
	}
	DB.SingularTable(true)
	DB.CreateTable(models.Ptk{})
}