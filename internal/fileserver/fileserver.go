package fileserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// Start 启动 HTTP 文件服务器
// 将指定目录的文件通过 HTTP 提供浏览和下载
func Start(ip string, port int, dir string) {
	// 验证目录是否存在
	if !isValidDirectory(dir) {
		log.Fatalf("✗ Invalid directory: %s", dir)
	}

	// 打印启动信息
	printServerInfo(ip, port)

	// 创建文件服务器
	handler := http.FileServer(http.Dir(dir))
	http.Handle("/", handler)

	addr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("✓ Serving directory [%s] on %s\n", dir, addr)
	fmt.Println("✓ Access it from your browser.")
	fmt.Println("-----------------------------------------")

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("✗ Failed to start server: %v", err)
	}
}

// isValidDirectory 检查路径是否是一个有效的、存在的目录
func isValidDirectory(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

// printServerInfo 打印服务器信息，包括局域网IP
func printServerInfo(ip string, port int) {
	fmt.Println("★ Go File Server is starting...")

	fmt.Println("✓ Local addresses:")
	fmt.Printf("  - http://localhost:%d\n", port)
	fmt.Printf("  - http://127.0.0.1:%d\n", port)

	if ip == "0.0.0.0" {
		ips := getLocalIPs()
		if len(ips) > 0 {
			fmt.Println("✓ On your network:")
			for _, addr := range ips {
				fmt.Printf("  - http://%s:%d\n", addr, port)
			}
		} else {
			fmt.Println("✗ Could not find local network IP addresses.")
		}
	} else {
		fmt.Println("✓ Listening on:")
		fmt.Printf("  - http://%s:%d\n", ip, port)
	}
}

// getLocalIPs 获取所有非环回的IPv4地址
func getLocalIPs() []string {
	var ips []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips
}
