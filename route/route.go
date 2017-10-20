package route

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"crud_2/controllers"
)

func Route(){
	e := echo.New()
	
		// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
		// Routes
	e.POST("/users", controllers.CreateUser)
	e.GET("/getuser/:id", controllers.GetUser )
	e.PUT("/getuser/:id", controllers.UpdateUser)
	e.DELETE("/getuser/:id", controllers.Delete)
	e.GET("/getAll", controllers.GetAll)
		
	e.Logger.Fatal(e.Start(":1323"))
}