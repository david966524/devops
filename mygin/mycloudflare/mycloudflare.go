package mycloudflare

import (
	"context"
	"fmt"
	"log"
	"mygin/myutils"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

var api *cloudflare.API

func init() {
	api = myutils.GetCfApi()
}

// 获取全部域名列表
func GetAllDomain() []Myzone {
	//api, err := cloudflare.New("300c6bba9f8c1cb0ae3a436d1828639c1b9ab", "2898802425@qq.com")
	// api := cfutils.GetCfApi()
	fmt.Println(api)

	resp, err := api.ListZones(context.Background())

	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}
	var MyzoneSlice []Myzone
	for _, cz := range resp {
		myzone := toMyzone(cz)
		MyzoneSlice = append(MyzoneSlice, myzone)
	}
	return MyzoneSlice
}

// 查询单个域名
func QueryDomain(Doname string) []Myzone {
	// api := cfutils.GetCfApi()
	resp, err := api.ListZones(context.Background(), Doname)
	if err != nil {
		// 处理错误
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	var MyzoneSlice []Myzone
	for _, cz := range resp {
		myzone := toMyzone(cz)
		MyzoneSlice = append(MyzoneSlice, myzone)
	}

	return MyzoneSlice
}

// 添加单个域名
func AddDomain(domainName string) []Myzone {
	// api := cfutils.GetCfApi()
	zoneParams := cloudflare.ZoneCreateParams{
		Name:      domainName,
		JumpStart: false,
	}
	zone, err := api.CreateZone(context.Background(), zoneParams.Name, zoneParams.JumpStart, cloudflare.Account{}, "")

	if err != nil {
		fmt.Println(err.Error())
	}
	var MyzoneSlice []Myzone
	myzone := toMyzone(zone)
	MyzoneSlice = append(MyzoneSlice, myzone)

	return MyzoneSlice
}

// 删除单个域名
func DeleteDomain(id string) bool {
	// api := cfutils.GetCfApi()
	zone_id, err := api.DeleteZone(context.Background(), id)
	if err != nil {
		log.Println("error:" + err.Error())
	}
	log.Println(zone_id.ID)
	if zone_id.ID != "" {
		return true
	} else {
		return false
	}

}

// 获得单个域名的dns解析
func GetDNSRecord(domainName string) []cloudflare.DNSRecord {
	// api := cfutils.GetCfApi()
	zoneID, err := api.ZoneIDByName(domainName)
	if err != nil {
		log.Fatal(err)
	}

	recs, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		log.Fatal(err)
	}
	return recs
}

// 添加单个域名 解析
func AddDnsRecord(zoneID string, dnsRecordParams cloudflare.CreateDNSRecordParams) (bool, string) {

	// api := cfutils.GetCfApi()

	recordID, err := api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), dnsRecordParams)
	if err != nil {
		errMsg := fmt.Sprintf("Error %v", err)
		fmt.Println(errMsg)
		return false, errMsg
	}
	log.Println(recordID.Response)
	log.Println(recordID.Result)
	log.Println(recordID.ResultInfo)
	return true, "add Success"
}

// 删除单个域名解析
func DeleteDnsRecord(zone_id string, recordID string) (bool, string) {
	// api := cfutils.GetCfApi()
	err := api.DeleteDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zone_id), recordID)
	if err != nil {
		errMsg := fmt.Sprintf("Error %v", err)
		fmt.Println(errMsg)
		return false, errMsg
	}
	return true, "delete Success"
}

// zone 结构体 定义所需要的 字段
type Myzone struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedOn   time.Time `json:"created_on"`
	ModifiedOn  time.Time `json:"modified_on"`
	NameServers []string  `json:"name_servers"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
}

// cloudflare.zong 转换成 myzone
func toMyzone(cz cloudflare.Zone) Myzone {
	return Myzone{
		ID:          cz.ID,
		Name:        cz.Name,
		CreatedOn:   cz.CreatedOn,
		ModifiedOn:  cz.ModifiedOn,
		NameServers: cz.NameServers,
		Status:      cz.Status,
		Type:        cz.Type,
	}
}
