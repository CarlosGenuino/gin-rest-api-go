package controllers

import (
	"net/http"

	"github.com/CarlosGenuino/gin-api-rest/database"
	"github.com/CarlosGenuino/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibirTodosAlunos(ctx *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	ctx.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"api_diz": "Olá " + nome + ", tudo beleza?",
	})
}

func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibirAlunosPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAlunos(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})
		return
	}
	database.DB.Delete(&aluno, id)
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}

func BuscarAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "aluno não encontrado",
		})
	}
	c.JSON(http.StatusOK, aluno)
}

func PaginaAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "alunos.html", gin.H{
		"alunos": alunos,
	})
}
func Pagina404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "index.html", gin.H{
		"message": "Pagina não encontrada! =[",
	})
}
