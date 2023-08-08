package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define router with user
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

const (
	// Two ways of Regex
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"

	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

	userIdKey = "userId"
	bizLogin  = "login"
)

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	server.POST("/users/signup", u.SignUp)

	server.POST("/users/login", u.Login)

	server.POST("/users/edit", u.Edit)

	server.GET("/users/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	// Bind method will parse the request body and store the result in req
	// Through Content-Type
	// If parse failed, it will return error 400
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// Validate email

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "System Error")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "Invalid Email")
		return
	}

	// Validate password
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "Password and Confirm Password are not the same")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "System Error")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "Password should contain at least 8 characters, 1 number and 1 special character")
		return
	}
	ctx.String(http.StatusOK, "Sign Up Success")
	fmt.Printf("%v", req)
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
