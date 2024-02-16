package main

import (
	//"fmt"

	//"github.com/spf13/pflag"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var (
//	mode   string
//	params []int
//)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	r.Run()
	//pflag.StringVar(&mode, "mode", "", "Mode of the calculator to use")
	//	pflag.IntSliceVar(&params, "params", []int{}, "Parameters")
	//	pflag.Parse()
	//	result := calculate()
	//	fmt.Println(len(params))
	//	fmt.Println("Results:", result)
}

//func calculate() int {
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
