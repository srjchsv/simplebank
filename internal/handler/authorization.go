package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/simplebank/internal/services/responses"
	"github.com/srjchsv/simplebank/util"
)

const (
	tokenTTL        = 15 * time.Minute
	cookieAgeSignIn = int(tokenTTL / time.Second)
)

type authId struct {
	Id int `json:"id"`
}

type token struct {
	Token string `json:"token"`
}

type signIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var client http.Client

func (h *Handler) UserIdentity(ctx *gin.Context) {
	token, err := ctx.Cookie("access_token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	cookie := &http.Cookie{
		Name:   "access_token",
		Value:  token,
		MaxAge: cookieAgeSignIn,
	}
	requestUrl := fmt.Sprintf("http://%s:%d/api", os.Getenv("AUTH_HOST"), util.AuthPort())
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	req.AddCookie(cookie)
	response, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var req signIn
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	requestUrl := fmt.Sprintf("http://%s:%d/auth/sign-in", os.Getenv("AUTH_HOST"), util.AuthPort())

	postBody, err := json.Marshal(map[string]string{
		"username": req.Username,
		"password": req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}

	if response.StatusCode != 200 {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse(err))
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}

	var token token
	json.Unmarshal([]byte(body), &token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse(err))
		return
	}
	ctx.SetCookie(
		"access_token",
		token.Token,
		cookieAgeSignIn,
		"/",
		"/",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token.Token,
	})
}

func SignUp(name, username, password string) (int, error) {
	requestURL := fmt.Sprintf("http://%s:%d/auth/sign-up", os.Getenv("AUTH_HOST"), util.AuthPort())

	postBody, _ := json.Marshal(map[string]string{
		"name":     name,
		"username": username,
		"password": password,
	})

	response, err := http.Post(requestURL, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return 0, err
	}
	if response.StatusCode != 200 {
		return 0, errors.New("auth error")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	var id authId
	json.Unmarshal([]byte(body), &id)
	if err != nil {
		return 0, err
	}
	return id.Id, err
}
