package models


type Ptk struct{
	Id int `gorm:"primary_key" json`
	Nama string `gorm:"varchar:50;not_null" json`
	Alamat string `gorm:"varchar:50;not_null" json`
}