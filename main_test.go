package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/CarlosGenuino/gin-api-rest/controllers"
	"github.com/CarlosGenuino/gin-api-rest/database"
	"github.com/CarlosGenuino/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var CPF string

func SetupRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	CPF = "12312312312"
	aluno := models.Aluno{Nome: "Mock", RG: "123456789", CPF: CPF}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVericaStatusCodeSaudacao(t *testing.T) {
	r := SetupRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	nome := "carlos"
	req, _ := http.NewRequest("GET", "/"+nome, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockResposta := `{"api_diz":"Olá carlos, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListarTodosAlunos(t *testing.T) {
	database.ConectaBancoDeDados()
	CriarAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos", controllers.ExibirTodosAlunos)
	req, _ := http.NewRequest("GET", "/api/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaBancoDeDados()
	CriarAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos/cpf/"+CPF, controllers.BuscarAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/api/alunos/cpf/"+CPF, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, CPF, "12312312312")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaBancoDeDados()
	CriarAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRotasDeTeste()
	r.GET("/api/alunos/:id", controllers.ExibirAlunosPorId)
	path := "/api/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Mock", alunoMock.Nome)
}

func TestDeleteAluno(t *testing.T) {
	CriarAlunoMock()
	r := SetupRotasDeTeste()
	r.DELETE("/api/alunos/:id", controllers.DeletarAlunos)
	path := "/api/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "HTTP_STATUS deveriam ser iguais")
}

func TestEditaAluno(t *testing.T) {
	CriarAlunoMock()
	defer DeleteAlunoMock()
	r := SetupRotasDeTeste()
	r.PUT("/api/alunos/:id", controllers.EditarAluno)
	path := "/api/alunos/" + strconv.Itoa(ID)
	aluno := models.Aluno{Nome: "Mockado", CPF: "17205578921", RG: "227779547"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "HTTP_STATUS deveriam ser iguais")
	var alunoAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoAtualizado)
	assert.Equal(t, aluno.Nome, alunoAtualizado.Nome)
	assert.Equal(t, aluno.CPF, alunoAtualizado.CPF)
	assert.Equal(t, aluno.RG, alunoAtualizado.RG)
}
