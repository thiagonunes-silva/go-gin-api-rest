package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagonunes.silva/go-gin-api-rest/database"
	"github.com/thiagonunes.silva/go-gin-api-rest/models"
)

func ListarAlunos(ctx *gin.Context) {
	var listaAlunos []models.Aluno
	database.DB.Find(&listaAlunos)
	ctx.JSON(http.StatusOK, listaAlunos)
}

func DetalharAluno(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	aluno := consultarAlunoById(id)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado!",
		})
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}

func consultarAlunoById(id string) models.Aluno {
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	return aluno
}

func Saudacao(ctx *gin.Context) {
	nome := ctx.Params.ByName("nome")
	ctx.JSON(http.StatusOK, gin.H{
		"API diz:": "E aí " + nome + ", tudo bom?",
	})
}

func CriarAluno(ctx *gin.Context) {
	var aluno models.Aluno
	err := ctx.ShouldBindJSON(&aluno)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	ctx.JSON(http.StatusOK, aluno)
}

func DeletarAluno(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	aluno := consultarAlunoById(id)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado!",
		})
		return
	}
	database.DB.Delete(&models.Aluno{}, id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func AtualizarAluno(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	aluno := consultarAlunoById(id)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado!",
		})
		return
	}
	err := ctx.ShouldBindJSON(&aluno)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	ctx.JSON(http.StatusOK, aluno)
}

func ConsultarAlunoByCpf(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado!",
		})
		return
	}
	ctx.JSON(http.StatusOK, aluno)
}
