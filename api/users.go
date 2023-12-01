package api

import (
	"context"
	"encoding/json"
	"github.com/Inf85/go-project-demo/auth"
	"github.com/Inf85/go-project-demo/models/users"
	"net/http"
)

type UserJSON struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (api *API) UserSignUp(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsonData := UserJSON{}
	err := decoder.Decode(&jsonData)

	if err != nil || jsonData.UserName == "" || jsonData.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(context.Background(), jsonData.UserName) {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(context.Background(), jsonData.UserName, jsonData.Password)

	jsonToken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonToken))
}

func (api *API) UserLogin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsonData := UserJSON{}
	err := decoder.Decode(&jsonData)

	if err != nil || jsonData.UserName == "" || jsonData.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}
	user := api.users.FindUser(context.Background(), jsonData.UserName)
	if user.UserName == "" {
		http.Error(w, "username not found", http.StatusBadRequest)
		return
	}

	if !api.users.CheckPassword(user.Password, jsonData.Password) {
		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *users.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(context.Background(), userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
