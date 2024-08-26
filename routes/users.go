package routes

import (
	"net/http"

	"API.com/models"
	"API.com/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	//Mapeia os campos do JSON da requisição HTTP para os campos correspondentes no struct user.
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message:": "User created succesfully!"})
}

func login(context *gin.Context) {
	// Extrair as credenciais (email e senha que foram enviados pelo cliente)
	// e queremos validá-los se uma combinação válida de email e senha foi enviada
	var user models.User

	// Quero vincular o corpo da solicitação de entrada a essa struct
	// Quero criar essa struct com os dados que estão anexados à solicitação
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Depois de verificar o erro e preencher o usuário, é hora de validar as credenciais que foram enviadas com a solicitação
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}



	context.JSON(http.StatusOK, gin.H{"message":"Login sucessfully!", "token": token})
}
