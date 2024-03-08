package wechatminiprogram

import (
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/golog"
)

// 缓存前缀
// wechat_mini:wechat_access_token:
type redisCachePrefixFun func() (wechatAccessToken string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId     string `json:"app_id"` // 小程序唯一凭证，即 appId
	AppSecret string // 小程序唯一凭证密钥，即 appSecret
}

// Client 实例
type Client struct {
	config struct {
		appId           string // 小程序唯一凭证，即 appId
		appSecret       string // 小程序唯一凭证密钥，即 appSecret
		accessToken     string // 接口调用凭证
		jsapiTicket     string // 签名凭证
		selfAccessToken bool   // 自己设置接口调用凭证
	}
	cache struct {
		redisClient             *redis.Client // 缓存数据库
		wechatAccessTokenPrefix string        // AccessToken
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret

	return c, nil
}
