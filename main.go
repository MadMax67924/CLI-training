package main

import (
	//"fmt"

	//"github.com/spf13/pflag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var (
//
//	mode   string
//	params []int
type data struct {
	Mode   string `json:"mode"`
	Params []int  `json:"params"`
}

var datas = []data{}

func getdata(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)
}

func main() {
	r := gin.Default()
	r.GET("/calculator", getdata)
	r.POST("/calculator", func(c *gin.Context) {
		var ndata data
		if err := c.BindJSON(&ndata); err != nil {
			return
		}
		fmt.Println(datas)
		datas = append(datas, ndata)
		c.IndentedJSON(http.StatusCreated, ndata)
		fmt.Println(datas)
		fmt.Println("data created")
		fmt.Println(ndata)
	})
	r.Run("localhost:8080")
	//pflag.StringVar(&mode, "mode", "", "Mode of the calculator to use")
	//	pflag.IntSliceVar(&params, "params", []int{}, "Parameters")
	//	pflag.Parse()
	//	result := calculate()
	//	fmt.Println(len(params))
	//	fmt.Println("Results:", result)
}

//func calculate() int {
//
//	result := params[0]
//
//	for i := 1; i < len(params); i++ {
//		switch mode {
//		case "sum":
//			result += params[i]
//		case "sustract":
//			result -= params[i]
//		case "multiply":
//			result *= params[i]
//		case "divide":
//			result /= params[i]
//		default:
//			fmt.Println("Operación no válida")
//			return 0
//		}
//	}
//
//	return result
//}
