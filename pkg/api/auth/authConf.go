package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	config "github.com/munsheerck79/Ecom_project.git/pkg/config"
	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

// ========================== JWT Token and cookie session  ========================== //

func JwtCookieSetup(c *gin.Context, name string, userId uint) bool {
	//time = 10 mins
	cookieTime := time.Now().Add(10 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(userId),
		ExpiresAt: cookieTime,
	})

	// Generate signed JWT token using env var of secret key
	if tokenString, err := token.SignedString([]byte(config.GetJWTConfig())); err == nil {

		// Set cookie with signed string if no error time = 10 hours
		c.SetCookie(name, tokenString, 10*3600, "", "", false, true)

		fmt.Println("JWT sign & set Cookie successful")
		return true
	}
	fmt.Println("Failed JWT setup")
	return false

}

//=======================================================token validation==========================================================

func ValidateToken(tokenString string) (jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.GetJWTConfig()), nil
		},
	)
	if err != nil || !token.Valid {
		fmt.Println("not valid token")
		return jwt.StandardClaims{}, errors.New("not valid token")
	}

	// then parse the token to claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		fmt.Println("can't parse the claims")
		return jwt.StandardClaims{}, errors.New("can't parse the claims")
	}

	return *claims, nil
}

//============================================================== otp (twilio)========================================================

var (
	AUTHTOKEN  string
	SERVICESID string
	ACCOUNTSID string
	client     *twilio.RestClient
)

func TwilioSendOTP(phoneNumber string) (string, error) {

	SERVICESID = config.GetConfig().SERVICESID
	ACCOUNTSID = config.GetConfig().ACCOUNTSID
	AUTHTOKEN = config.GetConfig().AUTHTOKEN
	// // ACCOUNTSID = "AC2c8bf06da44f10b978038088c472580e"
	// // AUTHTOKEN = "9ac2047bdc66673ca4d6703303a33f0a"
	// // SERVICESID = "VA61d4b739330f391e167f7b0cad47d861"
	// fmt.Println("ser", SERVICESID)
	// fmt.Println("acc", ACCOUNTSID)
	// fmt.Println("auth", AUTHTOKEN)
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: AUTHTOKEN,
		Username: ACCOUNTSID,
	})
	if client != nil {
		fmt.Println("Twilio connected")
	} else {
		fmt.Println("Twilio connection error")
	}

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(SERVICESID, params)

	if err != nil {
		fmt.Println("error at twilio")
		return "", err
	}
	fmt.Println("return from auth")
	return *resp.Sid, nil
}

//=======================================================twilio verification ===============================================

func TwilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(SERVICESID, params)

	if err != nil {
		return err
	} else if *resp.Status != "approved" {
		return errors.New("OTP verification failed")
	}

	return nil
}

func GetUserIdFromContext(ctx *gin.Context) uint {
	userIdStr := ctx.GetString("userId")
	userIdInt, _ := strconv.Atoi(userIdStr)
	return uint(userIdInt)
}
