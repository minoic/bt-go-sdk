package bt_go_sdk

/*
 带注释为必填 其余为选填
 省略项默认填为默认值或空 推荐多填
*/

// 获取网站列表
// URI 地址：/data?action=getData&table=sites
type ReqSites struct {
	P int64
	Limit int64 // 必填
	Type int64
	Order string
	ToJS string
	Search string
}

