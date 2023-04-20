package Nasabah

import (
	"fmt"
	"go_rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(n *gin.Context) {
	var Nasabah []models.Nasabah
	err := models.GetAll(&Nasabah)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, Nasabah)
	}
}

func Create(n *gin.Context) {
	var Nasabah models.Nasabah
	n.BindJSON(&Nasabah)
	err := models.Create(&Nasabah)
	if err != nil {
		fmt.Println(err.Error())
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, Nasabah)
	}
}

func Detail(n *gin.Context) {
	id := n.Params.ByName("id")
	var Nasabah models.Nasabah
	err := models.Detail(&Nasabah, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, Nasabah)
	}
}

func Update(n *gin.Context) {
	var Nasabah models.Nasabah
	id := n.Params.ByName("id")
	err := models.Detail(&Nasabah, id)
	if err != nil {
		n.JSON(http.StatusNotFound, Nasabah)
	}
	n.BindJSON(&Nasabah)
	err = models.Update(&Nasabah, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, Nasabah)
	}
}


func Delete(n *gin.Context) {
	var user models.Nasabah
	id := n.Params.ByName("id")
	err := models.Delete(&user, id)
	if err != nil {
		n.AbortWithStatus(http.StatusNotFound)
	} else {
		n.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
