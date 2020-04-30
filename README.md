用 http 对接 go 和宝塔 api 的 sdk

参考 https://www.bt.cn/bbs/thread-20376-1-1.html

* 返回 json 自动解析为 struct
* 保存返回的 cookies 并在之后的请求中使用来提高效率
* 所有 API 通过单元测试 测试版本为 6.9.8（免费版）