package middlewares

import (
	"net/http"

	"API.com/utils"
	"github.com/gin-gonic/gin"
)

// Uma função extra que basicamente é executada no meio da solicitação, antes que o handler de solicitação final se torne ativo.

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// Interrompe a solicitação atual ou a geração da resposta atual e envia essa resposta
		// e nenhum outro manipulador de solicitação será executado posteriormente
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userID", userID)

	// Se chegar ao final dessa função de autenticação, se tivermos um token válido também devemos chamar uma função especial
	// no context, que garantirá que o proximo handler de solicitações na linha seja executado corretamente
	context.Next()
}
