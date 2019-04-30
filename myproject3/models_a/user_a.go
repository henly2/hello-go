package models_a

type (
	UserInfo struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
		Sex      string `json:"sex"`
		Password string `json:"password"`
	}
	AdminInfo struct {
		Adminname     string `json:"adminname"`
		Adminage      int    `json:"adminage"`
		Adminsex      string `json:"adminsex"`
		Adminpassword string `json:"adminpassword"`
	}

	User struct {
		Name     string `form:"name" json:"name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	InforMation struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password" binding:"required"`
	}
	Loginlog struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)
