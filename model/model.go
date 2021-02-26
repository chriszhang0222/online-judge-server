package model

type Problem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Diff string `json:"diff"`
}

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Role int32 `json:"role"`
}

type Submit struct {
	ProblemId string `json:"id"`
	UserId string `json:"userId"`
}