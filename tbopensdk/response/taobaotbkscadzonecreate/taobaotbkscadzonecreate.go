package taobaotbkscadzonecreate

import (
	"encoding/json"
	"github.com/ChinaArJun/tksdk/tbopensdk/response"
)

// taobao.tbk.sc.adzone.create( 淘宝客-服务商-创建推广者位 )
type Response struct {
	response.TopResponse
	TbkScAdzoneCreateResponse TbkScAdzoneCreateResponse `json:"tbk_sc_adzone_create_response"`
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

type TbkScAdzoneCreateResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Model string `json:"model"`
}
