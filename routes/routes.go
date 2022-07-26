package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagonunes.silva/go-gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ListarAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarAluno)
	r.GET("/alunos/:id", controllers.DetalharAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.AtualizarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.ConsultarAlunoByCpf)
	r.Run()
}
