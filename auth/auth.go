package auth

import (
	"github.com/Inf85/go-project-demo/models/users"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"net/http"
	"time"
)

// signingKey set up a global string for our secret
var signingKey = []byte("knrjkevdckjh")

// JwtMiddleware handler for jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	},
	UserProperty:  "user",
	Debug:         false,
	SigningMethod: jwt.SigningMethodHS256,
})

// GetToken create a jwt token with user claims
func GetToken(user *users.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user.UUID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	signedToken, _ := token.SignedString(signingKey)
	return signedToken
}

// GetJSONToken create a JSON token string
func GetJSONToken(user *users.User) string {
	token := GetToken(user)
	jsonToken := "{\"id_token\": \"" + token + "\"}"

	return jsonToken
}

// GetUserClaimsFromContext return "user" claims as a map from request
func GetUserClaimsFromContext(req *http.Request) map[string]interface{} {
	claims := req.Context().Value("user").(jwt.Token).Claims.(jwt.MapClaims)

	return claims
}
