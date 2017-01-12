package main

import "gopkg.in/gin-gonic/gin.v1"
import "github.com/manveru/faker"
import "time"

func queryData() string {
	fake, _ := faker.New("en")
	time.Sleep(100 * time.Millisecond)
	return fake.Name()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"data": queryData(),
		})
	})
	r.Run(":8000")
}
