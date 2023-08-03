package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/auth"
)

//===================================================user middleware==============================================

func UserAuthentication(ctx *gin.Context) {
	authHelper(ctx, "user")
}
// =================================================admin middle ==================================================

func AdminAuthentication(ctx *gin.Context) {
	authHelper(ctx, "admin")
}

func authHelper(ctx *gin.Context, user string) {

	tokenString, err := ctx.Cookie(user + "-auth") // get cookie for user or admin with name

	if err != nil {
		fmt.Println("didnt get jwt")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
		})
		return
	}

	claims, err := auth.ValidateToken(tokenString) // auth function validate the token and return claims with error
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login token not valid",
		})
		return
	}

	// check the cliams expire time
	if time.Now().Unix() > claims.ExpiresAt {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}
	// claim the" userId and set it on context
	ctx.Set("userId", fmt.Sprint(claims.Id))

	//
}
