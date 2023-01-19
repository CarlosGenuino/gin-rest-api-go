package routes

import (
	"github.com/CarlosGenuino/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	addr := "127.0.0.1:8500"
	alunosPath := "/api/alunos"
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET(alunosPath, controllers.ExibirTodosAlunos)
	r.GET(alunosPath+"/:id", controllers.ExibirAlunosPorId)
	r.GET(alunosPath+"/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	r.GET("/saudacao/:nome", controllers.Saudacao)
	r.GET("/index", controllers.PaginaIndex)
	r.GET("/alunos", controllers.PaginaAlunos)

	r.POST("/api/alunos", controllers.CriarNovoAluno)
	r.DELETE("/api/alunos/:id", controllers.DeletarAlunos)
	r.PUT(alunosPath+"/:id", controllers.EditarAluno)
	r.Run(addr)
}
