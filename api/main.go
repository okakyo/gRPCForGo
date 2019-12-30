package api

import(
	"fmt"
	"github.com/labstack/echo",
	"./handler"

)

type userType struct {
	user_id  string `json:"name" form:"name" query:"name"`
	password string `json:"email" form:"email" query:"email"`
	nickname string `json:" nickname" form:"nickname" query:"nickname"`
	comment  string `json:"comment" form:"comment" query:"comment"`
}



func main(){
	e:=echo.New();
	e.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello World")
	})
	e.GET("/users/{user_id}", func(context  echo.Context) error {
		return context.JSON(200,"HelloWorld")
	})
	e.POST("/signup", func(context echo.Context) error {
		return context.JSON(200,"GoodBye")
	})
	e.PATCH("/users/{user_id}", func(context echo.Context) error {
		return context.JSON(200,"Seeyou")
	})
	e.DELETE("/close", func(context echo.Context) error {
		return  context.JSON(200,"GoodBtw")
	})
	fmt.Println("Start the Server");
	e.Start(":3000");



}

