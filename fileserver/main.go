package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	Username   string
	Password   string
	Port       string
	DirToShare string
)

func init() {
	flag.StringVar(&Username, "u", "admin", "用户名")
	flag.StringVar(&Password, "p", "123456", "密码")
	flag.StringVar(&Port, "port", "8080", "端口")
	flag.StringVar(&DirToShare, "dir", "./", "共享目录路径")
	flag.Parse()
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取浏览器弹窗输入的账号密码
		username, password, ok := r.BasicAuth()

		// 验证账号密码是否正确
		if !ok || username != Username || password != Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted Files"`)
			http.Error(w, "Unauthorized (拒绝访问，请输入正确的密码)", http.StatusUnauthorized)
			return
		}
		// 验证通过，放行继续访问文件
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 创建基础的文件服务器
	fileServer := http.FileServer(http.Dir(DirToShare))

	// 套上我们写的“密码验证外壳”
	protectedFileServer := authMiddleware(fileServer)

	log.Printf("文件服务器已启动，正在监听 http://localhost%s , 共享目录: %s", Port, DirToShare)
	if err := http.ListenAndServe(":"+Port, protectedFileServer); err != nil {
		log.Fatal(err)
	}
}
