package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	cron "gopkg.in/robfig/cron.v2"
)

func main() {
	RunCron()
	Init()

}

func Init() {
	r := gin.Default()
	r.GET("/sent", SendHelloMessage)
	r.Run(":8080")
}

func RunCron() {
	c := cron.New()
	c.AddFunc("@every 00h00m02s", sentMessage)
	c.Start()
}

func sentMessage() {
	resp, err := http.Get("http://localhost:8080/sent")
	if err != nil {
		fmt.Println("error from cron:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error from read all:", err)
	}
	fmt.Println(string(body))
}

func SendHelloMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello",
	})
}
