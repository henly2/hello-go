package models

type (
	UserInfo struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Password string `json:"password"`
	}

	User struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	ChangePwd struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		OldPassword string `json:"oldpassword" binding:"required"`
		NewPassword string `json:"newpassword" binding:"required"`
	}
)
