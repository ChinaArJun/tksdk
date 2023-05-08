package tbkprivilegeget

import (
	json "encoding/json"

	"github.com/ChinaArJun/tksdk/tbopensdk/response"
)

// 淘宝客商品详情查询（简版） 免费 不需要授权
type Response struct {
	response.TopResponse
	TbkPrivilegeGetResult Result `json:"tbk_privilege_get_response"`
}

// 解析输出结果
func (t *Response) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type Result struct {
	Result    ResultData `json:"result"`
	RequestID string     `json:"request_id"`
}

type ResultData struct {
	Data Data `json:"data"`
}

type Data struct {
	CategoryID          int64          `json:"category_id"`
	CouponClickURL      string         `json:"coupon_click_url"`
	CouponEndTime       string         `json:"coupon_end_time"`
	CouponInfo          string         `json:"coupon_info"`
	CouponRemainCount   int64          `json:"coupon_remain_count"`
	CouponStartTime     string         `json:"coupon_start_time"`
	CouponTotalCount    int64          `json:"coupon_total_count"`
	CouponType          int64          `json:"coupon_type"`
	ItemID              interface{}    `json:"item_id"`
	ItemURL             string         `json:"item_url"`
	MaxCommissionRate   string         `json:"max_commission_rate"`
	YsylClickUrl        string         `json:"ysyl_click_url"`
	YsylTljFace         string         `json:"ysyl_tlj_face"`
	YsylTljSendTime     string         `json:"ysyl_tlj_send_time"`
	YsylCommissionRate  string         `json:"ysyl_commission_rate"`
	YsylTljUseStartTime string         `json:"ysyl_tlj_use_start_time"`
	YsylTljUseEndTime   string         `json:"ysyl_tlj_use_end_time"`
	MinCommissionRate   string         `json:"min_commission_rate"`
	RewardInfo          int64          `json:"reward_info"`
	TopnInfo            TopnInfo       `json:"topn_info"`
	MiniProgram         MiniProgramDto `json:"mini_program"`
}

type TopnInfo struct {
	TopnEndTime    int64 `json:"topn_end_time"`
	TopnQuantity   int64 `json:"topn_quantity"`
	TopnStartTime  int64 `json:"topn_start_time"`
	TopnTotalCount int64 `json:"topn_total_count"`
}
type MiniProgramDto struct {
	MiniProgramAppid     string `json:"mini_program_appid"`
	MiniProgramPath      string `json:"mini_program_path"`
	MiniProgramQrcodeUrl string `json:"mini_program_qrcode_url"`
}

/*
"topn_info": {
"topn_end_time": 1623599999000,
"topn_quantity": 1884,
"topn_start_time": 1623313762000,
"topn_total_count": 2000
}
*/
