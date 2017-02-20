package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var mySigningKey = []byte("mySuperSecret")

// JwtController -> Auth Controller
type JwtController struct {
	beego.Controller
}

// Auth -> Generate token
func (c *JwtController) Auth() {
	type Auth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type User struct {
		name  string
		roles []string
	}

	var a Auth
	json.Unmarshal(c.Ctx.Input.RequestBody, &a)

	if a.Username == "admin" && a.Password == "admin" {
		u := &User{
			name:  a.Username,
			roles: []string{"admin"},
		}

		claims := make(jwt.MapClaims)
		claims["user"] = u.name
		claims["roles"] = u.roles
		claims["iat"] = time.Now().Unix()
		claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

		// Create the token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		// Sign the token with our secret
		tokenString, _ := token.SignedString(mySigningKey)

		c.Data["json"] = tokenString
	} else {
		c.Data["json"] = "Invalid credentials"
	}

	// Finally, write the token to the browser window
	c.ServeJSON()
}

// IsTokenValid -> check if token is valida
func (c *JwtController) IsTokenValid(req *http.Request) bool {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	if token == nil {
		return false
	}

	if token.Valid {
		return true
	}

	return false
}
