package main

import(
	"fmt"
	"os"
	"log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator"
	"github.com/jinzhu/gorm"
	- "github.com/go-sql-driver/mysql"

)

type UserInfo struct {
	user_id  string  `json:"user_id" form:"user_id" validate:"required"`
	password string  `json:"password" form:"password" validate:"required"`
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

type successResponse struct {


}

type apiError struct {
	message string
	cause string
}




func main()  {
	e:=echo.New();
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.GET("/", func(context echo.Context) error {
		return context.String(200, "Hello World")
	})
	e.GET("/users/:user_id", func(context  echo.Context) error {
		userId:=context.Param("user_id")


		return context.JSON(200,"HelloWorld")
	})
	e.POST("/signup", func(c echo.Context) error {
		inputedUser:=new (UserInfo)
		if err:=c.Bind(inputedUser);err!=nil {
			return c.JSON(500,"Server Error")
		}
		if err:=c.Validate(inputedUser); err!=nil {
			return  c.JSON(400, "invalid Input")
		}


		return c.JSON(200,"GoodBye")
	})
	e.PATCH("/users/{user_id}", func(context echo.Context) error {
		return context.JSON(200,"Seeyou")
	})
	e.DELETE("/close", func(context echo.Context) error {
		return  context.JSON(200,"GoodBtw")
	})
	fmt.Println("Start the Server");
	e.Logger.Fatal(e.Start(":"+port))

}
