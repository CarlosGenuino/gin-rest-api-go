package routes

import (
	"github.com/CarlosGenuino/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	addr := "127.0.0.1:8500"
	r := gin.Default()
	r.GET("/api/alunos", controllers.ExibirTodosAlunos)
	r.GET("/api/alunos/:id", controllers.ExibirAlunosPorId)

	r.POST("/api/alunos", controllers.CriarNovoAluno)

	r.GET("/saudacao/:nome", controllers.Saudacao)

	r.Run(addr)
}
