package bt_go_sdk

import (
	"fmt"
	"testing"
)

func TestClient_GetNetWork(t *testing.T) {
	client:=NewClient("http://10.0.0.14:8888","qviqWLiiUB623bfzJqQ37OGUEXwOXtVN")
	network, err := client.GetNetWork()
	if err != nil {
		t.Fail()
	}
	network, err = client.GetNetWork()
	if err != nil {
		t.Fail()
	}
	network, err = client.GetNetWork()
	if err != nil {
		t.Fail()
	}
	fmt.Println(network)
}