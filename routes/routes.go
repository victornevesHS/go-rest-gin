package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victornevesHS/go-rest-gin/controllers"
)

func HandleReqiests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
