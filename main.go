package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Mode   string `json:"mode"`
	Params []int  `json:"params"`
}

func main() {
	r := gin.Default()
	//	r.GET("/calculator", getdata)
	r.POST("/calculator", func(c *gin.Context) {
		var newdata data
		var result int
		if err := c.BindJSON(&newdata); err != nil {
			return
		}
		result = calculate(newdata.Params, newdata.Mode)
		c.IndentedJSON(http.StatusCreated, result)
	})
	r.Run("localhost:8080")
}

func calculate(params []int, mode string) int { //with a array
	result := params[0]

	for i := 1; i < len(params); i++ {
		switch mode {
		case "sum":
			result += params[i]
		case "sustract":
			result -= params[i]
		case "multiply":
			result *= params[i]
		case "divide":
			result /= params[i]
		default:
			return 0
		}
	}

	return result
}
