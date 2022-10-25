package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"exemple_oauth/config"
	"exemple_oauth/server"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type OauthHandler struct {
	server server.Server
}

func NewOauthHandler(server *server.Server) OauthHandler {
	return OauthHandler{server: *server}
}

func (o OauthHandler) GetInfo(c echo.Context) error {
	cfg := config.LoadOAUTHConfiguration()

	oauthState := generateStateOauthCookie(c.Response())

	u := cfg.AuthCodeURL(oauthState)

	err := c.Redirect(http.StatusTemporaryRedirect, u)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (o OauthHandler) CallBack(c echo.Context) error {
	data, err := getUserDataFromGoogle(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		err = c.Redirect(http.StatusTemporaryRedirect, "/")
		if err != nil {
			return err
		}
		return fmt.Errorf("redirect")
	}

	fprintf, err := fmt.Fprintf(c.Response(), "UserInfo: %s\n", data)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fprintf)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	cfg := config.LoadOAUTHConfiguration()

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	return state
}
