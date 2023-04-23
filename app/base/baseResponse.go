package base

//分頁資訊
type Pagination struct {
	Page       int64 `json:"page"`        //當前頁數
	TotalRows  int64 `json:"total_rows"`  //全部行數
	TotalPages int64 `json:"total_pages"` //全部頁數
	Size       int64 `json:"size"`        //一頁幾筆資料
}
