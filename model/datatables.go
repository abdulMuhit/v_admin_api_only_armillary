package model

type DatatablesResult struct{
	Draw uint8 `json:"draw"`
	RecordsTotal int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Data []map[string]interface{} `json:"data"`
}

type DatatablesSource struct{
	Draw uint8 `form:"draw"`
	OrdCol uint8 `form:"order[0][column]"`
	OrdDir string `form:"order[0][dir]"`
	Start int `form:"start"`
	Length int `form:"length"`
	Keyword string `form:"search[value]"`
}
