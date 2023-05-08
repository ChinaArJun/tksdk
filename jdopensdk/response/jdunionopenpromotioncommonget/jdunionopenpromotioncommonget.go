package jdunionopenpromotioncommonget

import (
	"encoding/json"
	"github.com/ChinaArJun/tksdk/jdopensdk/response"
)

// Response jd.union.open.promotion.common.get 网站/APP来获取的推广链接
type Response struct {
	response.TopResponse
	Responce Responce `json:"jd_union_open_promotion_common_get_responce"`
}

// WrapResult 解析输出结果
func (t *Response) WrapResult(result string) {
	err := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if err != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Message = err.Error()
	} else {
		//解析queryResult
		if t.Responce.QueryResultStr == "" {
			t.ErrorResponse.Code = -2
			t.ErrorResponse.Message = "empty queryResult"
		} else {
			if err = json.Unmarshal([]byte(t.Responce.QueryResultStr), &t.Responce.QueryResult); err != nil {
				t.ErrorResponse.Code = -1
				t.ErrorResponse.Message = err.Error()
			} else {
				t.ErrorResponse.Code = t.Responce.QueryResult.Code
				t.ErrorResponse.Message = t.Responce.QueryResult.Message
				t.ErrorResponse.RequestID = t.Responce.QueryResult.RequestID
			}
		}
	}
	t.Responce.QueryResultStr = ""
}

// Responce 响应结果
type Responce struct {
	Code           string `json:"code"`
	QueryResultStr string `json:"getResult"`
	QueryResult    QueryResult
}

// QueryResult 具体内容
type QueryResult struct {
	Code      int64  `json:"code"`
	Data      Data   `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type Data struct {
	ClickURL string `json:"clickURL"`
	JCommand string `json:"jCommand"`
}
