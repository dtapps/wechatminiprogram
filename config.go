package wechatminiprogram

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
)

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.appId = appId
	c.config.appSecret = appSecret
	return c
}

// ConfigRedisClient 缓存数据库
func (c *Client) ConfigRedisClient(client *dorm.RedisClient) {
	c.cache.redisClient = client
}

// ConfigRedisCachePrefixFunWechatAccessToken 缓存前缀
func (c *Client) ConfigRedisCachePrefixFunWechatAccessToken(config string) error {
	c.cache.wechatAccessTokenPrefix = config
	if c.cache.wechatAccessTokenPrefix == "" {
		return redisCachePrefixNoConfig
	}
	return nil
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}
