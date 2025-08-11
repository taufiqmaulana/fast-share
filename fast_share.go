package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	allArgs := os.Args
	filePath := allArgs[1]
	fileName := filepath.Base(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	randomPath := GenerateRandomString(5)
	http.HandleFunc(fmt.Sprintf("/fast-share/%s", randomPath), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeContent(w, r, fileName, fileInfo.ModTime(), file)
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	IP := GetLocalIP()
	path := fmt.Sprintf("http://%s:%d/fast-share/%s", IP, port, randomPath)

	fmt.Printf("\n\n%s\n", fileName)
	PrintQr(path)
	fmt.Printf(">> Scan QR or open this URL\n>> %s\n", path)
	fmt.Println(">> Pres CTRL + C to end share session")

	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}

func PrintQr(content string) {
	qrterminal.GenerateHalfBlock(content, qrterminal.M, os.Stdout)
}

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)
	return localAddress.IP
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
