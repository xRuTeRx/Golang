package structure

type Request struct {
	Id          int    `json:"Id"`
	City        string `json:"City"`
	RequestTime string `json:"RequestTime"`
	Temperature string `json:"Temperature"`
}
