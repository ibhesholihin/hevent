package routes

import (
	h "github.com/ibhesholihin/hevent/apps/http/handler"
	"github.com/ibhesholihin/hevent/apps/http/middleware"
	"github.com/labstack/echo/v4"
)

//type Routes struct {}
//func NewRoutes() *Routes {		return &Routes{}	}

func MyRoutes(e *echo.Echo, middleware *middleware.Middleware, myHandler *h.Handlers) {

	apiV1 := e.Group("/api/v1")

	//Auth admin handler
	apiV1.POST("/auth/admin/signup", myHandler.AdminSignUp)
	apiV1.POST("/auth/admin/signin", myHandler.LoginAdmin)

	//Auth user handler
	apiV1.POST("/auth/user/signup", myHandler.UserSignUp)
	apiV1.POST("/auth/user/signin", myHandler.LoginUser)

	public := apiV1.Group("/public")
	public.GET("/category", myHandler.GetListCategory)
	public.GET("/category/:categoryid", myHandler.GetCategory)
	public.GET("/events", myHandler.GetListEvents)
	public.GET("/events/:eventid", myHandler.GetEvent)

	//Admin Route
	//Backend Management Route by Admin
	admin := apiV1.Group("/admin", middleware.JWTAuthAdmin())
	admin.GET("/profile", myHandler.GetAdminProfile)
	admin.PUT("/updateprofile", myHandler.UpdateAdmin)

	admin.POST("/category", myHandler.AddCategory)
	admin.PUT("/category/:categoryid", myHandler.UpdateEventCategory)
	//admin.DELETE("/category/:categoryid", myHandler.RemoveEventCategory)

	admin.POST("/events", myHandler.AddEvent)
	admin.PUT("/events/:eventid", myHandler.UpdateEvent)
	admin.DELETE("/events/:eventid", myHandler.RemoveEvent)

	admin.GET("/users", myHandler.GetListUsers)

	//User Route
	//Apps Management Route by user
	user := apiV1.Group("/user", middleware.JWTAuthUser())
	user.GET("/profile", myHandler.GetUserProfile)
	user.PUT("/updateprofile", myHandler.UpdateUserProfile)

	/*
		//Order Ticketing
		user.GET("/users/:userid/cart", userHandler.GetCartSession)
		user.POST("/users/:userid/cart/:sessionid/item", userHandler.AddItemToCart)
		user.GET("/users/:userid/cart/:sessionid/item", userHandler.GetItemsCart)
		user.DELETE("/users/:userid/cart/:sessionid/item/:itemid", userHandler.DeleteItemFromCart)
		user.POST("/users/:userid/order/create/:sessionid", userHandler.CreateOrder)
		user.GET("/users/:userid/order", userHandler.GetListOrders)
		user.GET("/users/:userid/order/:orderid", userHandler.GetOrderById)
		user.PUT("/users/:userid/order/:orderid/payment/:paymentid", userHandler.UploadReceipt)
	*/
}
