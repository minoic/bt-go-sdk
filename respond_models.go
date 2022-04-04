package bt_go_sdk

/*
 *定义返回的 json 解析到目标结构体 由 json-to-go 自动生成
 相应结构详见本目录的 api-doc.pdf
 或 https://www.bt.cn/api-doc.pdf
 对应功能若有不同则以本目录文档为准
*/

// NetWork 获取实时状态信息(CPU、内存、网络、负载)
// URI 地址：/system?action=GetNetWork
type NetWork struct {
	Load struct {
		Max     int     `json:"max"`     // 最高值
		Safe    float64 `json:"safe"`    // 安全值
		One     float64 `json:"one"`     // 1 分钟
		Five    float64 `json:"five"`    // 5 分钟
		Limit   int     `json:"limit"`   // 限制
		Fifteen float64 `json:"fifteen"` // 10 分钟
	} `json:"load"` // 负载实时信息
	Down      float64 `json:"down"`      // 下行流量 （KB）
	DownTotal int64   `json:"downTotal"` // 总接收 （Byte）
	Version   string  `json:"version"`   // 面板版本
	Mem       struct {
		MemFree     int `json:"memFree"`     // 可用内存（MB）
		MemTotal    int `json:"memTotal"`    // 总共内存（MB）
		MemCached   int `json:"memCached"`   // 缓存化内存（MB）
		MemBuffers  int `json:"memBuffers"`  // 系统缓冲（MB）
		MemRealUsed int `json:"memRealUsed"` // 实际使用内存（MB）
	} `json:"mem"` // 内存实时信息
	Up        float64 `json:"up"`        // 上行流量（KB）
	UpTotal   int64   `json:"upTotal"`   // 总发送 （Byte）
	UpPackets int     `json:"upPackets"` // 总发包 （个）
	Disk      []struct {
		Path   string   `json:"path"`   // 挂载点
		Inodes []string `json:"inodes"` // Inode使用信息 数组同下
		Size   []string `json:"size"`   // 0-总共（GB） 1-已用（GB） 2-可用（GB） 3-使用率（百分比 带%）
	} `json:"disk"` // 磁盘
	DownPackets int           `json:"downPackets"` // 总收包（个）
	CPU         []interface{} `json:"cpu"`         // 0-总体使用率 1-核心数 2-[0-CPU0 1-CPU1]使用率 3-CPU型号
}

// SystemTotal 获取系统基础统计
// URI 地址：/system?action=GetSystemTotal
type SystemTotal struct {
	CPURealUsed float64 `json:"cpuRealUsed"` // cpu使用率（百分比）
	MemTotal    int     `json:"memTotal"`    // 物理内存容量（MB）
	System      string  `json:"system"`      // 操作系统信息
	MemRealUsed int     `json:"memRealUsed"` // 物已使用的物理内存 （MB）
	CPUNum      int     `json:"cpuNum"`      // CPU 核心数
	MemFree     int     `json:"memFree"`     // 可用物理内存
	Version     string  `json:"version"`     // 面板版本
	Time        string  `json:"time"`        // 上次开机到现在的运行时间
	MemCached   int     `json:"memCached"`   // 缓存化的内存
	MemBuffers  int     `json:"memBuffers"`  // 系统缓冲 （MB）
	Isuser      int     `json:"isuser"`      // ？
}

// DiskInfo 获取磁盘分区信息
// URI 地址：/system?action=GetDiskInfo
type DiskInfo []struct {
	Path   string   `json:"path"`
	Inodes []string `json:"inodes"`
	Size   []string `json:"size"`
}

// UpdateStatus 检查面板更新
// URI 地址：/ajax?action=UpdatePanel
type UpdateStatus struct {
	Status    bool   `json:"status"`
	Version   string `json:"version"`
	UpdateMsg string `json:"updateMsg"`
}

// RespSites 获取网站列表
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

// SiteTypes 获取网站分类
// URI 地址：/site?action=get_site_types
type SiteTypes []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PHPVersions 获取已安装的 PHP 版本列表
// URI 地址：/site?action=GetPHPVersion
type PHPVersions []struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

// RespAddSite 创建网站
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

// RespMSG 通用消息结构
type RespMSG struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

// RespSiteBackups 获取网站备份列表
// URI 地址：/data?action=getData&table=backup
type RespSiteBackups struct {
	Data []struct {
		Name     string `json:"name"`
		Addtime  string `json:"addtime"`
		Pid      int    `json:"pid"`
		Filename string `json:"filename"`
		ID       int    `json:"id"`
		Size     int    `json:"size"`
	} `json:"data"`
	Where string `json:"where"`
	Page  string `json:"page"`
}

// SiteDomains 获取网站的域名列表
// URI 地址：/data?action=getData&table=domain
type SiteDomains []struct {
	Port    int    `json:"port"`
	Addtime string `json:"addtime"`
	Pid     int    `json:"pid"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

// RewriteList 伪静态可用列表
type RewriteList struct {
	Rewrites []string `json:"rewrite"`
}

// RespGetFile 获取指定文件
type RespGetFile struct {
	Status   bool   `json:"status"`
	Data     string `json:"data"`
	Encoding string `json:"encoding"`
}

type RespUserINI struct {
	Pass    bool `json:"pass"`
	Logs    bool `json:"logs"`
	Userini bool `json:"userini"`
	RunPath struct {
		Dirs    []string `json:"dirs"`
		RunPath string   `json:"runPath"`
	} `json:"runPath"`
}

// RespLimitNet 获取网络限制
type RespLimitNet struct {
	LimitRate int `json:"limit_rate"`
	Perserver int `json:"perserver"`
	Perip     int `json:"perip"`
}
