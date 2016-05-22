package models

type EasyuiDataGrid struct {
	Total int         `json:"total"`
	Rows  interface{} `json:"rows"`
}
