package main

import "gopkg.in/gin-gonic/gin.v1"
import "github.com/manveru/faker"
import "time"

func queryData(c chan string) {
	fake, _ := faker.New("en")
	time.Sleep(100 * time.Millisecond)
	c <- fake.Name()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		chA := make(chan string)
		chB := make(chan string)
		go queryData(chA)
		go queryData(chB)
		c.JSON(200, gin.H{
			"data1": <-chA,
			"data2": <-chB,
		})
	})
	r.Run(":8000")
}
