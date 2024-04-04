package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tb")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 注册路由处理函数
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	// 配置跨域
	cors := setupCors()
	handler := cors.Handler(http.DefaultServeMux)

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

// 设置CORS
func setupCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},                   // 允许的源
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的HTTP方法
		AllowedHeaders:   []string{"*"},                                       // 允许的请求头
		AllowCredentials: true,                                                // 允许发送身份验证凭据（cookies）
	})
	return c
}

// 登录处理函数
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 添加响应头
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "不允许的请求方法", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("接收到的用户名:", username)
	fmt.Println("接收到的密码:", password)

	// 查询数据库以验证用户
	row := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		http.Error(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, `{"message": "登录成功", "userID": %d}`, id)
}

// 注册处理函数
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 添加响应头
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "不允许的请求方法", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// 将新用户插入到数据库中
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		http.Error(w, "注册失败", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"message": "注册成功"}`)
}
