package main

import(
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

)

func main()  {
	e:=echo.New();
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
	e.Logger.Fatal(e.Start(":1323"))

}
