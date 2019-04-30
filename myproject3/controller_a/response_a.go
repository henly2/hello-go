package controller_a

type (
	Response struct {
		Err    int         `json:"err"`
		ErrMsg string      `json:"errmsg"`
		Result interface{} `json:"result"`
	}
)
