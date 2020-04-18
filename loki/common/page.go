package common

type PageData struct{
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	Total int64 `json:"total"`
	List interface{} `json:"list"`
}

func NewPageData(page int, pageSize int, total int64, list interface{}) *PageData {
	return &PageData{Page: page, PageSize: pageSize, Total: total, List: list}
}



