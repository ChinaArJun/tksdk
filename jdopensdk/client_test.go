package jdopensdk

import (
	"encoding/json"
	"fmt"
	"github.com/ChinaArJun/tksdk/jdopensdk/request"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopengoodsjingfenquery"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopengoodspromotiongoodsinfoquery"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopengoodsquery"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopenorderrowquery"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopenpositioncreate"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopenpositionquery"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopenpromotionbysubunionidget"
	"github.com/ChinaArJun/tksdk/jdopensdk/response/jdunionopenpromotioncommonget"
	"io/ioutil"
	"os"
	"testing"
)

var (
	appKey, appSecret, sessionKey, key string
)

func init() {
	if _, err := os.Stat("../dev_env.json"); err == nil {
		if bytes, err := ioutil.ReadFile("../dev_env.json"); err == nil {
			var data struct {
				Jd struct {
					AppKey     string `json:"app_key"`
					AppSecret  string `json:"app_secret"`
					SessionKey string `json:"session_key"`
					Key        string `json:"key"`
				} `json:"jd"`
			}
			if err = json.Unmarshal(bytes, &data); err == nil {
				appKey = data.Jd.AppKey
				appSecret = data.Jd.AppSecret
				sessionKey = data.Jd.SessionKey
				key = data.Jd.Key
			}
		}
	}
}
func GetClient() *TopClient {
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)
	return client
}

func TestJdUnionOpenGoodsJingfenQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsJingfenQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"goodsReq":{"eliteId":"9999","pid":"","sortName":"commission","sort":"desc","pageSize":"50","pageIndex":"1"}}`)
	var getResponse DefaultResponse = &jdunionopengoodsjingfenquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodsjingfenquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPromotionCommonGetRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPromotionCommonGetRequest{}
	getRequest.AddParameter("360buy_param_json", `{"promotionCodeReq":{"materialId":"https://item.jd.com/10030840282202.html","siteId":"223997188","positionId":"1631770896","subUnionId":"user_1","couponUrl":"","giftCouponKey":"","ext1":"user_1"}}`)
	var getResponse DefaultResponse = &jdunionopenpromotioncommonget.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpromotioncommonget.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPromotionBysubunionidGetRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPromotionBysubunionidGetRequest{}
	getRequest.AddParameter("360buy_param_json", `{"promotionCodeReq":{"materialId":"https://item.jd.com/10030840282202.html", "positionId":"1631770896","subUnionId":"user_1","couponUrl":"","giftCouponKey":"","chainType":"3"}}`)
	var getResponse DefaultResponse = &jdunionopenpromotionbysubunionidget.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpromotionbysubunionidget.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenOrderRowQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenOrderRowQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"orderReq":{"pageIndex":"1", "pageSize":"20","type":"3","startTime":"2021-08-02 11:45:00","endTime":"2021-08-02 12:45:00"}}`)
	var getResponse DefaultResponse = &jdunionopenorderrowquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenorderrowquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenGoodsPromotiongoodsinfoQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsPromotiongoodsinfoQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{skuIds:"10030840282202"}`)
	var getResponse DefaultResponse = &jdunionopengoodspromotiongoodsinfoquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodspromotiongoodsinfoquery.Response)

		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenGoodsQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenGoodsQueryRequest{}
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"skuIds":[10045819630288],"fields":"videoInfo"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"iPhone+13"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"小米+13"}}`)
	//getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"民法典"}}`)
	getRequest.AddParameter("360buy_param_json", `{"goodsReqDTO":{"keyword":"空调"}}`)
	var getResponse DefaultResponse = &jdunionopengoodsquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopengoodsquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPositionQueryRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPositionQueryRequest{}
	getRequest.AddParameter("360buy_param_json", `{"positionReq":{"unionId":103431087,"pageIndex":1,"pageSize":100,"key":"`+key+`","unionType":4}}`)
	var getResponse DefaultResponse = &jdunionopenpositionquery.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpositionquery.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
	}
}

func TestJdUnionOpenPositionCreateRequest(t *testing.T) {
	client := GetClient()
	getRequest := &request.JdUnionOpenPositionCreateRequest{}
	getRequest.AddParameter("360buy_param_json", `{"positionReq":{"unionId":103431087,"key":"`+key+`","unionType":3,"type":4,"spaceNameList":["测试专用-01"]}}`)
	var getResponse DefaultResponse = &jdunionopenpositioncreate.Response{}
	if err := client.Exec(getRequest, getResponse); err != nil {
		fmt.Println(err)
	} else {
		commonGetResponse := getResponse.(*jdunionopenpositioncreate.Response)
		fmt.Println(commonGetResponse.IsError())
		fmt.Println(commonGetResponse.Body)
		// {"jd_union_open_position_create_responce":{"code":"0","createResult":"{\"code\":200,\"data\":{\"resultList\":{\"测试专用-01\":21001775614},\"siteId\":0,\"type\":4,\"unionId\":103431087},\"message\":\"success\",\"requestId\":\"o_0b126c7d_l531g96m_4147537\"}"}}
	}
}
