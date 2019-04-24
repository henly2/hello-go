package main

import (
	"fmt"
	"net/http"
	"time"
)

type Cookie struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

func HandConn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	fmt.Println("r.url", r.URL)
	fmt.Println("r.Method=", r.Method)
	fmt.Println("r.Header=", r.Header)
	fmt.Println("r.body=", r.Body)
	//COOKIE_MAX_MAX_AGE     = time.Hour * 24 / time.Second   // 单位：秒。
	//maxAge = int(COOKIE_MAX_MAX_AGE)
	uid := "10"
	uid_cookie := &http.Cookie{
		Name:     "uid",
		Value:    uid,
		Path:     "/",
		HttpOnly: false,
		//MaxAge:   maxAge
	}

	http.SetCookie(w.Writer, uid_cookie)
}

func main() {
	http.HandleFunc("/", HandConn)
	//监听绑定
	http.ListenAndServe(":8000", nil)
}
