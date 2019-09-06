package ding

import (
	"errors"
)

type Client struct {
	requestHandler *requestHandler
	TokenContainer *tokenContainer
}

func NewClient(scheme, host string) *Client {
	client := &Client{}
	client.requestHandler = &requestHandler{
		scheme: scheme,
		host: host,
		charset: "UTF-8",
		contentType: "application/json",
		queryString:make(map[string]interface{}),
	}
	client.TokenContainer = &tokenContainer{}
	return client
}

func (c *Client) GetAccessToken(appKey, appSecret string) (token string, errCode int, errMsg string, expiresIn int64, err error) {
	for {
		if appKey == "" {
			err = errors.New("app key can't be empty")
			break
		}
		if appSecret == "" {
			err = errors.New("app secret can't be empty")
			break
		}

		queryString := make(map[string]interface{})
		queryString["appkey"] = appKey
		queryString["appsecret"] = appSecret

		c.requestHandler.setQueryString(&queryString)
		c.requestHandler.path = "gettoken"
		url, err := c.requestHandler.makeUrl()
		if err != nil {
			break
		}
		tokenJson, err := c.requestHandler.get(url)
		if err != nil {
			break
		}
		tokenMap, err := Json2map(tokenJson)
		if err != nil {
			break
		}
		token = ToString(tokenMap["access_token"])
		errMsg = ToString(tokenMap["errmsg"])
		errCode = int(tokenMap["errcode"].(float64))
		expiresIn = int64(tokenMap["expires_in"].(float64))
		break
	}
	defer func() {
		c.requestHandler.queryString = make(map[string]interface{}, 0)
	}()
	return token, errCode, errMsg, expiresIn, err
}

func (c *Client) Send(requestParams *map[string]interface{}) (resp map[string]interface{}, err error) {
	var respStr string
	for {
		if len(*requestParams) == 0 && c.requestHandler.method == "post" {
			err = errors.New("request body can't be empty")
			break
		}
		accessToken := &map[string]interface{}{
			"access_token": c.TokenContainer.accessToken,
		}

		c.requestHandler.setQueryString(accessToken)
		url, err := c.requestHandler.makeUrl()
		if err != nil {
			break
		}
		switch c.requestHandler.method {
		case "post":
			respStr, err = c.requestHandler.post(url, requestParams)
		case "get":
			c.requestHandler.setQueryString(requestParams)
			respStr, err = c.requestHandler.get(url)
		default:
			c.requestHandler.setQueryString(requestParams)
			respStr, err = c.requestHandler.get(url)
		}

		if err != nil {
			break
		}
		resp, err = Json2map(respStr)
		break
	}
	defer func() {
		c.requestHandler.queryString = make(map[string]interface{}, 0)
	}()
	return resp, err
}



