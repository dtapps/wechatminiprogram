package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type BusinessGetLiveInfoResponse struct {
	Errcode  int    `json:"errcode"` //  // 错误码，0代表成功，1代表未创建直播间
	Errmsg   string `json:"errmsg"`  // 错误信息
	Total    int    `json:"total"`
	RoomInfo []struct {
		Name       string `json:"name"`   // 直播间名称
		Roomid     int    `json:"roomid"` // 直播间ID
		CoverImg   string `json:"cover_img"`
		ShareImg   string `json:"share_img"`
		LiveStatus int    `json:"live_status"`
		StartTime  int    `json:"start_time"`
		EndTime    int    `json:"end_time"`
		AnchorName string `json:"anchor_name"`
		Goods      []struct {
			CoverImg        string `json:"cover_img"`
			Url             string `json:"url"`
			Name            string `json:"name"`
			Price           int    `json:"price"`
			Price2          int    `json:"price2"`
			PriceType       int    `json:"price_type"`
			GoodsId         int    `json:"goods_id"`
			ThirdPartyAppid string `json:"third_party_appid"`
		} `json:"goods"`
		LiveType      int    `json:"live_type"`
		CloseLike     int    `json:"close_like"`
		CloseGoods    int    `json:"close_goods"`
		CloseComment  int    `json:"close_comment"`
		CloseKf       int    `json:"close_kf"`
		CloseReplay   int    `json:"close_replay"`
		IsFeedsPublic int    `json:"is_feeds_public"`
		CreaterOpenid string `json:"creater_openid"`
		FeedsImg      string `json:"feeds_img"`
	} `json:"room_info"`
}

type BusinessGetLiveInfoResult struct {
	Result BusinessGetLiveInfoResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func newBusinessGetLiveInfoResult(result BusinessGetLiveInfoResponse, body []byte, http gorequest.Response, err error) *BusinessGetLiveInfoResult {
	return &BusinessGetLiveInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// BusinessGetLiveInfo 获取直播间列表
// 调用此接口获取直播间列表及直播间信息
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/liveplayer/studio-api.html
func (c *Client) BusinessGetLiveInfo(notMustParams ...gorequest.Params) *BusinessGetLiveInfoResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/business/getliveinfo?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response BusinessGetLiveInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newBusinessGetLiveInfoResult(response, request.ResponseBody, request, err)
}
