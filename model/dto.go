package model

type LoginForm struct {
	AuthType string `json:"authType" form:"authType"`
	Identifier string `json:"identifier" form:"id" binding:"required"`
	Credential string `json:"credential" form:"pwd" binding:"required"`
}

type SignUpForm struct {
	AuthType string `json:"authType" form:"authType"`
	Identifier string `json:"identifier" form:"id" binding:"required"`
	Credential string `json:"credential" form:"pwd" binding:"required"`
}
