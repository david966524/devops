package model

type Line struct {
	BaseUrl    string `json:"base_url"`
	ResUrl     string `json:"res_url"`
	SocketIP   string `json:"socket_ip"`
	SocketPort int    `json:"socket_port"`
	Timeout    int    `json:"timeout"`
	Ssl        int    `json:"ssl"`
	Remark     string `json:"remark"`
	Type       int    `json:"type"`
}

type Data struct {
	Lines []Line `json:"lines"`
}

type Datas struct {
	Data Data `json:"data"`
	Ok   bool `json:"ok"`
}
