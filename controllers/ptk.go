package controllers

import(
	"net/http"
	"github.com/labstack/echo"
	"crud_2/database"
	"crud_2/models"
	"fmt"
	"strconv"

)


func CreateUser(c echo.Context) error{
	var ptk models.Ptk
	u := &models.Ptk{}
	if err := c.Bind(u); err != nil{
		return err
	}
	ptk.Id = 0
	ptk.Nama = u.Nama
	ptk.Alamat = u.Alamat
	database.DB.Create(&ptk)
	fmt.Println("___________________________________")
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error{
	var ptk models.Ptk
	id, _ := strconv.Atoi(c.Param("id"))
	sts := database.DB.Find(&ptk, id)
	if sts.RecordNotFound(){
		a := map[string] string{"message":"Not Found"}
		return c.JSON(http.StatusOK, a)
	}
	return c.JSON(http.StatusOK, ptk)
}

func UpdateUser(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var ptk models.Ptk
	a := database.DB.Find(&ptk,id)
	if a.RecordNotFound(){
		a := map[string] string{"message":"Not Found"}
		return c.JSON(http.StatusOK, a)
	}
	u := &models.Ptk{}
	if err := c.Bind(u); err != nil{
		return err
	}
	ptk.Nama = u.Nama
	ptk.Alamat = u.Alamat
	database.DB.Save(&ptk)
	message := map[string]string{"message":"suceess"}
	return c.JSON(http.StatusOK, message)
}

func Delete(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	var ptk models.Ptk
	err := database.DB.Where("id=?", id).Delete(&ptk)
	if err.Error!=nil{
		fmt.Println(err.Error)
	}
	return c.NoContent(http.StatusNoContent)
}

func GetAll(c echo.Context) error{
	var ptk []models.Ptk
	err := database.DB.Find(&ptk)
	if err.Error!= nil{
		fmt.Println(err.Error)
	}
	return c.JSON(http.StatusOK, ptk)
}
