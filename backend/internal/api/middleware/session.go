package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
)

func UseSession(accountRepository *repositories.AccountRepository) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		sessionID, err := c.Cookie("sessionID")
		fmt.Printf("Cookie value: %s \n", sessionID)

		if err != nil {
			fmt.Println(err.Error())
			//error getting cookie
		}

		if sessionID == "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			account, err := accountRepository.GetBySession(sessionID)
			if err != nil {
				fmt.Println(err.Error())
			}

			c.Set("accountName", account.AccountName)
			c.Set("apiKey", account.APIKey)

		}

		// TODO Add token
		// token := context.GetHeader("Authorization")
		// err := ValidateToken(token)
		// if err != nil {
		// 	context.JSON(401, gin.H{"error": err.Error()})
		// 	context.Abort()
		// 	return
		// }

		c.Next()
	})
}
