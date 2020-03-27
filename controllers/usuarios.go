package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/marceloagmelo/udemy-golang-mvc/models"
)

//Home é a página inicial da aplicação
func Home(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, Wordl!")
	var usuarios []models.Usuarios

	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar recuperar os dados",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Lista de Usuários",
		"usuarios": usuarios,
	}
	return c.Render(http.StatusOK, "index.html", data)
}

//Add adicionar usuário
func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

//Atualizar atualizar usuário
func Atualizar(c echo.Context) error {
	var usuarioID, _ = strconv.Atoi(c.Param("id"))

	var usuario models.Usuarios

	resultado := models.UsuarioModel.Find("id=?", usuarioID)
	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possivel encontrar o usuário!",
		})
	}

	if err := resultado.One(&usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possivel encontrar o usuário!",
		})

	}

	var data = map[string]interface{}{
		"usuario": usuario,
	}

	return c.Render(http.StatusOK, "update.html", data)

}

//Inserir os dados no banco de dados
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"mensagem": "Não foi possível adicionar o registro no banco!",
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Campos obrigatórios!",
	})
}

//Deletar registro no banco de dados
func Deletar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := models.UsuarioModel.Find("id=?", usuarioID)
	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possivel encontrar o usuário!",
		})
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possivel deletar o usuário!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Usuario deletado como sucesso!",
	})

}

//Update registro no banco de dados
func Update(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario = models.Usuarios{
		ID:    usuarioID,
		Nome:  nome,
		Email: email,
	}

	resultado := models.UsuarioModel.Find("id=?", usuarioID)
	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possivel encontrar o usuário!",
		})
	}

	if err := resultado.Update(usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar atualizar o usuário!",
		})
	}

	return c.JSON(http.StatusAccepted, usuario)
}
