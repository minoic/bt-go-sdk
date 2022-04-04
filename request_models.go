package bt_go_sdk

/*
 *定义请求参数较为复杂的结构体
 带注释为必填 其余为选填
 省略项默认填为默认值或空 推荐多填
*/

// ReqSites 获取网站列表
// URI 地址：/data?action=getData&table=sites
type ReqSites struct {
	P      int64
	Limit  int64 // 必填
	Type   int64
	Order  string
	ToJS   string
	Search string
}

// ReqAddSite 创建网站
// URI 地址：/site?action=AddSite
type ReqAddSite struct {
	WebName struct {
		Domain     string   `json:"domain"`     // 必填
		DomainList []string `json:"domainlist"` // 必填
		Count      int      `json:"count"`      // 必填
	}
	Path         string // 必填
	TypeID       int64  // 必填
	Type         string // 必填
	Version      int64  // 必填
	Port         int64  // 必填
	PS           string // 必填
	FTP          bool   // 必填
	FTPUserName  string // FTP 为 true 时 必填
	FTPPassword  string // FTP 为 true 时 必填
	SQL          bool   // 必填
	Codeing      string // SQL 为 true 时 必填
	DataUser     string // SQL 为 true 时 必填
	DataPassword string // SQL 为 true 时 必填
}

// ReqDeleteSite 删除网站
// URI 地址：/site?action=DeleteSite
type ReqDeleteSite struct {
	ID       int64  // 必填
	WebName  string // 必填
	FTP      bool
	Database bool
	Path     bool
}

// ReqSiteBackups 获取网站备份列表
// URI 地址：/data?action=getData&table=backup
type ReqSiteBackups struct {
	P      int64
	Limit  int64 // 必填
	Type   int64 // 必不填或填0
	ToJS   string
	Search int64 // 必填
}
