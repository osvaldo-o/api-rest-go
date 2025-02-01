package middleware

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var _user = os.Getenv("USER")
var _password = os.Getenv("PASSWORD")

func BasicAuthMiddleware(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el encabezado Authorization"})
		c.Abort()
		return
	}

	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Encabezado Authorization no válido"})
		c.Abort()
		return
	}

	credentials := strings.SplitN(string(decodedBytes), ":", 2)
	if len(credentials) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de credenciales inválido"})
		c.Abort()
		return
	}

	username, password := credentials[0], credentials[1]
	if username == _user && password == _password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		c.Abort()
		return
	}

	c.Next()
}
