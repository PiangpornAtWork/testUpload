package controller


import (

   "net/http"


   "example.com/test/config"

   "example.com/test/entity"

   "github.com/gin-gonic/gin"

)


func GetAll(c *gin.Context) {


   db := config.DB()


   var genders []entity.Genders

   db.Find(&genders)


   c.JSON(http.StatusOK, &genders)


}