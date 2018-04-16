# About
企业微信API官方文档: https://work.weixin.qq.com/api/doc

# Director

    ├── api // API 接口
    ├── examples // API接口的测试用例
    ├── config // 设置Debug模式
    ├── utils // 一些基础方法
    └── README.md

# Usage
```
import (
	"github.com/scarletcc/weworkapi_go/api"
)

// 实例化 API 类
corpAPI := api.NewCorpAPI("CORP_ID", "APP_SECRET")

// 发送消息
response, err := corpAPI.MessageSend(map[string]interface{}{
    "touser":  "",
    "toparty": "",
    "totag":   "",
    "msgtype": "text",
    "agentid": "APP_ID",
    "text": map[string]interface{}{
        "content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
    },
    "safe": 0,
})
```

# POST请求时url带参数的情况

```
postData := make(map[string]interface{})
str := `
{
    "urlArgs": {  // 将url参数放在"urlArgs"的map下，提交时会自动加到url中，并从postData中删除此字段
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
```

# 关于token的缓存
token是需要缓存的，不能每次调用都去获取token（否者会中频率限制）
在本库的设计里，token是以类里的一个变量缓存的
在类的生命周期里，这个accessToken都是存在的， 当且仅当发现token过期，CorpAPI类会自动刷新token
刷新机制在 api/API.go
所以，使用时，只需要全局实例化一个CorpAPI类，不要析构它，就可一直用它调函数，不用关心 token
