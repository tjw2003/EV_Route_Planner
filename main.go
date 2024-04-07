// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterRequest struct {
    NewUsername string `json:"newUsername"`
    NewPassword string `json:"newPassword"`
}


func main() {
	// 连接 MySQL 数据库
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/users")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("成功连接")
	}
	defer db.Close()
	// 测试连接是否正常
	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	var username string
	err = db.QueryRow("SELECT username FROM users WHERE id = ?", 1).Scan(&username)
	if err != nil {
		log.Fatal("Error fetching username:", err.Error())
	}
	fmt.Println("Username retrieved from database:", username)

	// 初始化 Gin 框架
	router := gin.Default()

	// 使用 CORS 中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"} // 允许的来源
	router.Use(cors.New(config))

	// 设置登录接口
	router.POST("/login", loginHandler)

	// 设置注册接口
	router.POST("/register", registerHandler)

	// 启动服务
	router.Run(":8888")
}

func getDatabaseName(db *sql.DB) string {
	var dbName string
	err := db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		log.Fatal("Failed to get database name:", err)
	}
	return dbName
}

func loginHandler(c *gin.Context) {
	if db == nil {
        log.Fatal("Database connection is nil")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// 可以从 req 中获取到正确的 username 和 password
	username := req.Username
	password := req.Password

	// 查询数据库，验证用户名和密码
	var storedPassword string
	// 打印数据库名字
	dbName := getDatabaseName(db)
	log.Println("Connected to database:", dbName)
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	log.Println(storedPassword)

	// 比对密码
	if storedPassword != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func registerHandler(c *gin.Context) {
	 // 解析注册请求的 JSON 数据
	 var req RegisterRequest
	 if err := c.ShouldBindJSON(&req); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		 return
	 }
 
	 // 输出请求体内容进行调试
	 log.Printf("Received register request: %+v\n", req)
 
	 // 从请求中获取用户名和密码
	 username := req.NewUsername
	 password := req.NewPassword
 
	 // 检查数据库连接是否有效
	 if db == nil {
		 log.Fatal("Database connection is nil")
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		 return
	 }
 
	 // 检查用户名是否已存在于数据库中
	 var existingUsername string
	 err := db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUsername)
	 if err == nil {
		 // 如果用户名已存在，则返回用户名已存在的错误
		 c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		 return
	 } else if err != sql.ErrNoRows {
		 // 如果查询出错，返回内部服务器错误
		 log.Println("Error checking username existence:", err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		 return
	 }
 
	 // 在数据库中插入新用户
	 _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	 if err != nil {
		 // 插入用户时发生错误，返回内部服务器错误
		 log.Println("Error registering user:", err)
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		 return
	 }
 
	 // 注册成功，返回成功消息
	 c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
