package taobaotbkitemclickextract

import (
	"encoding/json"

	"github.com/ChinaArJun/tksdk/tbopensdk/response"
)

// taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )
type Response struct {
	response.TopResponse
	TbkItemClickExtractResponse Result `json:"tbk_item_click_extract_response"`
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
	ItemID     string `json:"item_id"`
	OpenIID    string `json:"open_iid"`
	BizSceneID string `json:"biz_scene_id"`
}
