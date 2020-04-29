package bt_go_sdk

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	client := NewClient("http://10.0.0.14:8888", "qviqWLiiUB623bfzJqQ37OGUEXwOXtVN")
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
