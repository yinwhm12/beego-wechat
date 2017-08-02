package util

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"

)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn	float64 `json:"expires_in"`
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg string
}

func GetAccessToken(appID, appSecret, accessTokenURL string)(string, error)  {
	requestLine := strings.Join([]string{accessTokenURL,"?grant_type=client_credential&appid=",appID,"&secret=",appSecret},"")
	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK{
		fmt.Println("获取accessToken 失败", err)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil{
		fmt.Println("获取accessToken body 失败", err)
		return "", err
	}
	if bytes.Contains(body, []byte("access_tokne")){
		atp := AccessTokenResponse{}
		err = json.Unmarshal(body,&atp)
		if err != nil{
			fmt.Println("解析token失败",err)
			return "", err
		}
		return atp.AccessToken, nil
	}else{
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil{
			return "", err
		}
		return "", fmt.Errorf("%s",ater.Errmsg)
	}
}

