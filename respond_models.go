package bt_go_sdk

/*
 相应结构详见本目录的 api-doc.pdf
 或 https://www.bt.cn/api-doc.pdf
 对应功能若有不同则以本目录文档为准
*/

// 获取实时状态信息(CPU、内存、网络、负载)
// URI 地址：/system?action=GetNetWork
type NetWork struct {
	Load struct {
		Max     int     `json:"max"`
		Safe    float64 `json:"safe"`
		One     float64 `json:"one"`
		Five    float64 `json:"five"`
		Limit   int     `json:"limit"`
		Fifteen float64 `json:"fifteen"`
	} `json:"load"`
	Down      float64 `json:"down"`
	DownTotal int64   `json:"downTotal"`
	Version   string  `json:"version"`
	Mem       struct {
		MemFree     int `json:"memFree"`
		MemTotal    int `json:"memTotal"`
		MemCached   int `json:"memCached"`
		MemBuffers  int `json:"memBuffers"`
		MemRealUsed int `json:"memRealUsed"`
	} `json:"mem"`
	Up        float64 `json:"up"`
	UpTotal   int64   `json:"upTotal"`
	UpPackets int     `json:"upPackets"`
	Disk      []struct {
		Path   string   `json:"path"`
		Inodes []string `json:"inodes"`
		Size   []string `json:"size"`
	} `json:"disk"`
	DownPackets int       `json:"downPackets"`
	CPU         []float64 `json:"cpu"`
}

// 获取系统基础统计
// URI 地址：/system?action=GetSystemTotal
type SystemTotal struct {
	CPURealUsed float64 `json:"cpuRealUsed"`
	MemTotal    int     `json:"memTotal"`
	System      string  `json:"system"`
	MemRealUsed int     `json:"memRealUsed"`
	CPUNum      int     `json:"cpuNum"`
	MemFree     int     `json:"memFree"`
	Version     string  `json:"version"`
	Time        string  `json:"time"`
	MemCached   int     `json:"memCached"`
	MemBuffers  int     `json:"memBuffers"`
	Isuser      int     `json:"isuser"`
}

// 获取磁盘分区信息
// URI 地址：/system?action=GetDiskInfo
type DiskInfo []struct {
	Path   string   `json:"path"`
	Inodes []string `json:"inodes"`
	Size   []string `json:"size"`
}

// 检查面板更新
// URI 地址：/ajax?action=UpdatePanel
type UpdateStatus struct {
	Status    bool   `json:"status"`
	Version   string `json:"version"`
	UpdateMsg string `json:"updateMsg"`
}

// 获取网站列表
// URI 地址：/data?action=getData&table=sites
type RespSites struct {
	Data []struct {
		Status      string `json:"status"`
		Ps          string `json:"ps"`
		Domain      int    `json:"domain"`
		Name        string `json:"name"`
		Addtime     string `json:"addtime"`
		Path        string `json:"path"`
		BackupCount int    `json:"backup_count"`
		Edate       string `json:"edate"`
		ID          int    `json:"id"`
	} `json:"data"`
	Where string `json:"where"`
	Page  string `json:"page"`
}

// 获取网站分类
// URI 地址：/site?action=get_site_types
type SiteTypes []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 获取已安装的 PHP 版本列表
// URI 地址：/site?action=GetPHPVersion
type PHPVersions []struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

// 创建网站
// URI 地址：/site?action=AddSite
type RespAddSite struct {
	FtpStatus      bool   `json:"ftpStatus"`
	DatabaseUser   string `json:"databaseUser"`
	DatabaseStatus bool   `json:"databaseStatus"`
	FtpUser        string `json:"ftpUser"`
	DatabasePass   string `json:"databasePass"`
	SiteStatus     bool   `json:"siteStatus"`
	FtpPass        string `json:"ftpPass"`
}

// 通用消息结构
type RespMSG struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}
