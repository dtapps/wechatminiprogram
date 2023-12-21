package wechatminiprogram

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"go.dtapp.net/gojson"
	"strings"
)

type UserInfo struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
}

type UserInfoResponse struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
	Watermark struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

type UserInfoResult struct {
	Result UserInfoResponse // 结果
	Err    error            // 错误
}

func newUserInfoResult(result UserInfoResponse, err error) *UserInfoResult {
	return &UserInfoResult{Result: result, Err: err}
}

// UserInfo 解密用户信息
func (c *Client) UserInfo(ctx context.Context, param UserInfo) *UserInfoResult {
	var response UserInfoResponse
	aesKey, err := base64.StdEncoding.DecodeString(param.SessionKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	cipherText, err := base64.StdEncoding.DecodeString(param.EncryptedData)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	ivBytes, err := base64.StdEncoding.DecodeString(param.Iv)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = c.pkcs7Unpaid(cipherText, block.BlockSize())
	if err != nil {
		return newUserInfoResult(response, err)
	}
	err = gojson.Unmarshal(cipherText, &response)
	if err != nil {
		return newUserInfoResult(response, err)
	}
	if response.Watermark.AppID != c.GetAppId() {
		return newUserInfoResult(response, errors.New("c id not match"))
	}
	return newUserInfoResult(response, err)
}

func (u *UserInfoResponse) UserInfoAvatarUrlReal() string {
	return UserInfoAvatarUrlReal(u.AvatarUrl)
}

func UserInfoAvatarUrlReal(avatarUrl string) string {
	return strings.Replace(avatarUrl, "/132", "/0", -1)
}
