package bt_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ehang-io/nps/lib/crypt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	BTAddress string
	BTKey     string
	cookies   []*http.Cookie
}

func NewClient(address string, key string) *Client {
	return &Client{
		BTAddress: address,
		BTKey:     key,
	}
}

func (this *Client) btAPI(data map[string][]string, endpoint string) ([]byte, int) {
	requestURL, err := url.Parse(this.BTAddress + endpoint)
	if err != nil {
		panic(err)
	}
	// fmt.Println(this.cookies)
	nowTime := string(time.Now().Unix())
	requestToken, requestTime := crypt.Md5(nowTime+crypt.Md5(this.BTKey)), nowTime
	body := url.Values{
		"request_token": {requestToken},
		"request_time":  {requestTime},
	}
	for k, v := range data {
		body[k] = v
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{Jar: jar}
	if len(this.cookies) != 0 {
		client.Jar.SetCookies(requestURL, this.cookies)
	}
	resp, err := client.PostForm(requestURL.String(), body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode >= 400 {
		fmt.Println("failed post at ", requestURL.String())
	}
	this.cookies = resp.Cookies()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, resp.StatusCode
}

func (this *Client) GetNetWork() (NetWork, error) {
	resp, status := this.btAPI(map[string][]string{}, "/system?action=GetNetWork")
	if status >= 400 {
		return NetWork{}, errors.New(string(resp))
	}
	var dec NetWork
	if err := json.Unmarshal(resp, &dec); err != nil {
		return NetWork{}, err
	}
	return dec, nil
}

func (this *Client) GetSystemTotal() (SystemTotal, error) {
	resp, status := this.btAPI(map[string][]string{}, "/system?action=GetSystemTotal")
	if status >= 400 {
		return SystemTotal{}, errors.New(string(resp))
	}
	var dec SystemTotal
	if err := json.Unmarshal(resp, &dec); err != nil {
		return SystemTotal{}, err
	}
	return dec, nil
}

func (this *Client) GetDiskInfo() (DiskInfo, error) {
	resp, status := this.btAPI(map[string][]string{}, "/system?action=GetDiskInfo")
	if status >= 400 {
		return DiskInfo{}, errors.New(string(resp))
	}
	var dec DiskInfo
	if err := json.Unmarshal(resp, &dec); err != nil {
		return DiskInfo{}, err
	}
	return dec, nil
}

func (this *Client) GetTaskCount() int {
	resp, status := this.btAPI(map[string][]string{}, "/ajax?action=GetTaskCount")
	if status >= 400 {
		return 0
	}
	dec, err := strconv.Atoi(string(resp))
	if err != nil {
		return 0
	}
	return dec
}

func (this *Client) GetUpdateStatus(check bool, force bool) (UpdateStatus, error) {
	data := map[string][]string{
		"check": {strconv.FormatBool(check)},
		"force": {strconv.FormatBool(force)},
	}
	resp, status := this.btAPI(data, "/ajax?action=UpdatePanel")
	if status >= 400 {
		return UpdateStatus{}, errors.New(string(resp))
	}
	var dec UpdateStatus
	if err := json.Unmarshal(resp, &dec); err != nil {
		return UpdateStatus{}, err
	}
	return dec, nil
}

func (this *Client) GetSites(params *ReqSites) (RespSites, error) {
	data := map[string][]string{
		"p":      {strconv.FormatInt(params.P, 1)},
		"limit":  {strconv.FormatInt(params.Limit, 15)},
		"type":   {strconv.FormatInt(params.Type, -1)},
		"order":  {params.Order},
		"tojs":   {params.ToJS},
		"search": {params.Search},
	}
	resp, status := this.btAPI(data, "/data?action=getData&table=sites")
	if status >= 400 {
		return RespSites{}, errors.New(string(resp))
	}
	var dec RespSites
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespSites{}, err
	}
	return dec, nil
}

func (this *Client) AddSite(params *ReqAddSite) (RespAddSite, error) {
	webname, err := json.Marshal(params)
	if err != nil {
		return RespAddSite{}, err
	}
	data := map[string][]string{
		"webname":      {string(webname)},
		"path":         {params.Path},
		"type_id":      {strconv.FormatInt(params.TypeID, 0)},
		"type":         {params.Type},
		"version":      {strconv.FormatInt(params.Version, 72)},
		"port":         {strconv.FormatInt(params.Port, 80)},
		"ps":           {params.PS},
		"ftp":          {strconv.FormatBool(params.FTP)},
		"ftp_username": {params.FTPUserName},
		"ftp_password": {params.FTPPassword},
		"sql":          {strconv.FormatBool(params.SQL)},
		"codeing":      {params.Codeing},
		"datauser":     {params.DataUser},
		"datapassword": {params.DataPassword},
	}
	resp, status := this.btAPI(data, "/site?action=AddSite")
	if status >= 400 {
		return RespAddSite{}, errors.New(string(resp))
	}
	var dec RespAddSite
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespAddSite{}, err
	}
	return dec, nil
}

func (this *Client) DeleteSite(params *ReqDeleteSite) (RespMSG, error) {
	data := map[string][]string{
		"id":      {strconv.FormatInt(params.ID, 0)},
		"webname": {params.WebName},
	}
	if params.FTP {
		data["ftp"] = []string{"1"}
	}
	if params.Database {
		data["database"] = []string{"1"}
	}
	if params.Path {
		data["path"] = []string{"1"}
	}
	resp, _ := this.btAPI(data, "/site?action=DeleteSite")
	var dec RespMSG
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespMSG{}, err
	}
	return dec, nil
}

func (this *Client) StopSite(id int64, name string) (RespMSG, error) {
	data := map[string][]string{
		"id":   {strconv.FormatInt(id, 0)},
		"name": {name},
	}
	resp, _ := this.btAPI(data, "/site?action=SiteStop")
	var dec RespMSG
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespMSG{}, err
	}
	return dec, nil
}

func (this *Client) StartSite(id int64, name string) (RespMSG, error) {
	data := map[string][]string{
		"id":   {strconv.FormatInt(id, 0)},
		"name": {name},
	}
	resp, _ := this.btAPI(data, "/site?action=SiteStart")
	var dec RespMSG
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespMSG{}, err
	}
	return dec, nil
}

func (this *Client) SetSiteEdate(id int64, edate string) (RespMSG, error) {
	data := map[string][]string{
		"id":    {strconv.FormatInt(id, 0)},
		"edate": {edate},
	}
	resp, _ := this.btAPI(data, "/site?action=SetEdate")
	var dec RespMSG
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespMSG{}, err
	}
	return dec, nil
}

func (this *Client) SetSitePS(id int64, ps string) (RespMSG, error) {
	data := map[string][]string{
		"id": {strconv.FormatInt(id, 0)},
		"ps": {ps},
	}
	resp, _ := this.btAPI(data, "/data?action=setPs&table=sites")
	var dec RespMSG
	if err := json.Unmarshal(resp, &dec); err != nil {
		return RespMSG{}, err
	}
	return dec, nil
}
