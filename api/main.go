package api

import(
	"fmt"
	"github.com/labstack/echo",
	"./handler"

)

func responseValue(){
	return{}
}


func main(){
	e:=echo.New();

	e.GET("/users/{user_id}", func(context  echo.Context) error {

	})
	e.POST("/signup")
	e.PATCH("/users/{user_id}")
	e.DELETE("/close")
	fmt.Println("Start the Server");
	e.Start(":3000");



}

