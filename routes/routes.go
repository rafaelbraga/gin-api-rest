package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbraga/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.POST("/alunos/:id", controllers.DeletaAluno)
	r.POST("/alunos/", controllers.CriaNovoAluno)
	r.PATCH("/:alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	r.Run()
}
