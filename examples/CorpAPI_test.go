package examples

import (
	"testing"
	"github.com/scarletcc/weworkapi_go/api"
	"fmt"
	"github.com/scarletcc/weworkapi_go/config"
	"encoding/json"
)

func getCorpAPI() *api.CorpAPI {
	config.DEBUG = true
	return api.NewCorpAPI(TestConf["CORP_ID"], TestConf["APP_SECRET"])
}

// 测试获取token
func TestGetAccessToken(t *testing.T)  {
	corpAPI := getCorpAPI()
	token := corpAPI.GetAccessToken()
	fmt.Println(token)
}

// 测试获取应用信息
func TestAgentGet(t *testing.T) {
	corpAPI := getCorpAPI()
	response, err := corpAPI.AgentGet(map[string]interface{}{
		"agentid": TestConf["APP_ID"],
	})
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
}

// 测试设置应用
func TestAgentSet(t *testing.T) {
	corpAPI := getCorpAPI()
	response, err := corpAPI.AgentSet(map[string]interface{}{
		"agentid": TestConf["APP_ID"],
		"report_location_flag": 0,
		"name": "",
		"description": "",
		"redirect_domain": "",
		"isreportenter":0,
		"home_url":"",
	})
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
}

// 测试获取用户信息
func TestAgentUserGet(t *testing.T) {
	corpAPI := getCorpAPI()
	response, err := corpAPI.UserGet(map[string]interface{}{
		"userid": "",
	})
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
}

// 测试获取应用列表
func TestAgentGetList(t *testing.T) {
	corpAPI := getCorpAPI()
	response, err := corpAPI.AgentGetList(nil)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
}

// 测试创建菜单
func TestMenuCreate(t *testing.T) {
	corpAPI := getCorpAPI()
	postData := make(map[string]interface{})
	str := `
	{
		"urlArgs": {
			"agentid": "%v"
		},
		"button":[
		   {    
			   "type":"click",
			   "name":"今日歌曲",
			   "key":"V1001_TODAY_MUSIC"
		   },
		   {
			   "name":"菜单",
			   "sub_button":[
				   {
					   "type":"view",
					   "name":"搜索",
					   "url":"http://www.soso.com/"
				   },
				   {
					   "type":"click",
					   "name":"赞一下我们",
					   "key":"V1001_GOOD"
				   }
			   ]
		  }
		]
	}
	`
	json.Unmarshal([]byte(fmt.Sprintf(str, TestConf["APP_ID"])), &postData)

	response, err := corpAPI.MenuCreate(postData)

	fmt.Println(response, err)
}

// 测试删除菜单
func TestMenuDelete(t *testing.T) {
	corpAPI := getCorpAPI()
	response, err := corpAPI.MenuDelete(map[string]interface{}{
		"agentid": TestConf["APP_ID"],
	})
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
}

// 测试给指定用户发送消息
func TestMessageSend(t *testing.T) {
	corpAPI := getCorpAPI()
	postData := map[string]interface{}{
		"touser":  "",
		"toparty": "",
		"totag":   "",
		"msgtype": "text",
		"agentid": TestConf["APP_ID"],
		"text": map[string]interface{}{
			"content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
		},
		"safe": 0,
	}
	response, err := corpAPI.MessageSend(postData)

	fmt.Println(response, err)
}
