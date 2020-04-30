package bt_go_sdk

import (
	"fmt"
	"testing"
)

var client *Client

// 运行全部测试时请先填写对应信息

func init() {
	client = NewClient("http://10.0.0.14:8888", "qviqWLiiUB623bfzJqQ37OGUEXwOXtVN")
}

func Test(t *testing.T) {
	r1, _ := client.GetSites(&ReqSites{
		Limit: 15,
	})
	fmt.Println(r1)
	r, _ := client.GetSiteBackups(&ReqSiteBackups{
		Limit:  15,
		Type:   0,
		Search: 4,
	})
	fmt.Println(r)
}

func TestClient_GetNetWork(t *testing.T) {
	r, _ := client.GetNetWork()
	fmt.Println(r)
}

func TestClient_GetSystemTotal(t *testing.T) {
	r, err := client.GetSystemTotal()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetDiskInfo(t *testing.T) {
	r, err := client.GetDiskInfo()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetRewriteList(t *testing.T) {
	r, err := client.GetRewriteList("10.0.0.14")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetSites(t *testing.T) {
	r, err := client.GetSites(&ReqSites{
		P:     1,
		Limit: 15,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetSiteDomains(t *testing.T) {
	r, err := client.GetSiteDomains(4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetTaskCount(t *testing.T) {
	r := client.GetTaskCount()
	fmt.Println(r)
}

func TestClient_GetUpdateStatus(t *testing.T) {
	r, err := client.GetUpdateStatus(true, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_GetPHPVersion(t *testing.T) {
	r, err := client.GetPHPVersion()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_AddSite(t *testing.T) {
	r, err := client.AddSite(&ReqAddSite{
		WebName: struct {
			Domain     string   `json:"domain"`
			DomainList []string `json:"domainlist"`
			Count      int      `json:"count"`
		}{
			Domain:     "w1.hao.com",
			DomainList: []string{},
			Count:      0,
		},
		Path:         "/www/wwwroot/w1.hao.com",
		TypeID:       0,
		Type:         "PHP",
		Version:      73,
		Port:         80,
		PS:           "test",
		FTP:          true,
		FTPUserName:  "ftpusername",
		FTPPassword:  "ftppassword",
		SQL:          true,
		Codeing:      "utf8",
		DataUser:     "datauser",
		DataPassword: "datapassword",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r)
}

func TestClient_AddDomain(t *testing.T) {
	r2, err := client.AddDomain(11, "w1.hao.com", "w2.hao.com")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_DelDomain(t *testing.T) {
	r2, err := client.DelDomain(11, "w1.hao.com", "w2.hao.com", 80)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_SiteBackup(t *testing.T) {
	r2, err := client.SiteBackup(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_GetSiteBackups(t *testing.T) {
	r2, err := client.GetSiteBackups(&ReqSiteBackups{
		P:      1,
		Limit:  15,
		Search: 11,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_DeleteSiteBackup(t *testing.T) {
	r2, err := client.DeleteSiteBackup(540)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_SetSitePS(t *testing.T) {
	r2, err := client.SetSitePS(11, "testps")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_SetSiteEdate(t *testing.T) {
	r2, err := client.SetSiteEdate(11, "0000-00-00")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_StopSite(t *testing.T) {
	r2, err := client.StopSite(11, "w1.hao.com")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_StartSite(t *testing.T) {
	r2, err := client.StartSite(11, "w1.hao.com")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}

func TestClient_DeleteSite(t *testing.T) {
	r2, err := client.DeleteSite(&ReqDeleteSite{
		ID:       10,
		WebName:  "w1.hao.com",
		FTP:      true,
		Database: true,
		Path:     true,
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(r2)
}
