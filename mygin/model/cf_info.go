package model

import (
	"time"
)

type CloudflareRequest struct {
	DomainName  string `form:"domainName"`
	DomainId    string `form:"domainId"`
	DnsRecordId string `form:"dnsRecordId"`
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
