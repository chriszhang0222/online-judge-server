package form

type ProblemForm struct{
	Name string `form:"name" json:"name" binding:"required"`
	Diff string `form:"diff" json:"diff" binding:"required"`
	Desc string `form:"desc" json:"diff" binding:"required"`
	Id string `form:"id" json:"id" binding:"required"`
}
