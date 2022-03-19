package OAuth2

import (
	"awesomeProject1/Utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetLoginFunc(c *gin.Context) {

	code := c.Query("code")

	token, _ := GetToken(code)

	userInfo, _ := GetUserInfo(token)

	c.JSON(200, gin.H{
		"state": 10000,
		"data": map[string]interface{}{
			"AccessToken": Utils.CreateNewToken(fmt.Sprintln("%s", userInfo.Login)),
		},
	})

}

func GetToken(code string) (string, error) {

	url := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		ClientID, ClientSecret, code,
	)

	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return "", err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return "", err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

func GetUserInfo(token string) (githubAccount, error) {

	var userInfoUrl = "https://api.github.com/user"
	var req *http.Request
	var err error
	var userInfo githubAccount
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return userInfo, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response

	if res, err = client.Do(req); err != nil {
		return userInfo, err
	}

	ans, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(ans, &userInfo)

	fmt.Println(userInfo)

	return userInfo, nil

}
