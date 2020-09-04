#利用golang爬取天气预报

> 爬取城市天气，通过用一个HTML模板，渲染一个html页面，使用电子邮件每天推送给指定的用户。

recipient.json 目标用户邮箱及用户所在城市
```json 
{
  "from": [
    {
      "city": "chengdu",
      "user": [
        "*****@qq.com",
        "*****@mrray.cn",
        "*****@qq.com",
        "*****@mrray.cn",
        "*****@mrray.cn",
        "*****@mrray.cn",
        "*****@mrray.cn",
        "*****@qq.com",
        "*****@qq.com"
      ]
    },
    {
      "city": "beijing",
      "user": [
        "*****@qq.com",
        "*****@qq.com"
      ]
    }
  ]
}
