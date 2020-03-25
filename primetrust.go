package primetrust

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

const (
	Version             = "1.0.17"
	SandboxAPIPrefix    = "https://sandbox.primetrust.com/v2"
	ProductionAPIPrefix = "https://api.primetrust.com/v2"
)

var _apiPrefix string
var _authHeader string
var _jwt string

func basicAuth(username string, password string) string {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return auth
}

type JWT struct {
	Token string `json:"token`
}

func getJWT() string {
	url := "https://sandbox.primetrust.com/auth/jwts"
	method := "POST"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", _authHeader)
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	jwt := new(JWT)
	if err = json.Unmarshal(body, &jwt); err != nil {
		fmt.Println("primetrust.init().getJWT error parsing body:", err)
		return ""
	}
	if len(jwt.Token) == 0 {
		fmt.Println("error creating JWT")
		panic(0)
	}
	bearer := fmt.Sprintf("Bearer %+v", jwt.Token)

	return bearer
}

func Init(sandbox bool, login string, password string) {
	if sandbox {
		_apiPrefix = SandboxAPIPrefix
	} else {
		_apiPrefix = ProductionAPIPrefix
	}
	_authHeader = fmt.Sprintf("Basic %s", basicAuth(login, password))
	color.Blue("%s", _authHeader)
	_jwt = getJWT()
	color.Blue("%s", _jwt)
}
