package handler

import (
	dblayer "cloud-storage/db"
	"cloud-storage/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890"
)

// SignUpHandler: 处理用户注册请求
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	// 对密码做加盐hash加密
	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))

	// 注册
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

type SigninResponseData struct {
	Token    string
	Username string
	Location string
}

// SignInHandler: 用户登陆接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	enc_passwd := util.Sha1([]byte(password + pwd_salt))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, enc_passwd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2. 生成访问凭证token
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}
	// 3. 登陆成功后重定向到首页
	responseData := SigninResponseData{
		Token:    token,
		Username: username,
		Location: "http://" + r.Host + "/static/view/home.html",
	}

	// 使用json格式写回
	data, err := json.Marshal(responseData)
	if err != nil { // json转换失败
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

// 根据用户名生成用户的token，token是40位的字符串
func GenToken(username string) string {
	// md5(username + timestamp + token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
