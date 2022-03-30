package middleware

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/jayesh29patidar/golang-react-todo/database"
// 	"github.com/jayesh29patidar/golang-react-todo/models"
// )

// func AddUser(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	fmt.Println(r.Method)
// 	if r.Method == "OPTIONS" {
// 		return
// 	}

// 	var user models.User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	fmt.Println("user done", user)
// 	if user.Username == "" {
// 		return
// 	}
// 	var listedUser = database.InsertUser(user)
// 	json.NewEncoder(w).Encode(listedUser)
// }

// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	token := database.TokenGenerator()
// 	fmt.Println("login Token", token)

// 	var user models.User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	fmt.Println("login hit with user", user)

// 	var loggedUser = database.LoginCheck(user)
// 	var userWithToken models.UserWithToken
// 	userWithToken.Token = token
// 	userWithToken.User = loggedUser

// 	json.NewEncoder(w).Encode(userWithToken)
// }
