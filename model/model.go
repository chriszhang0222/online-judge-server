package model

type Problem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Diff string `json:diff`
}
