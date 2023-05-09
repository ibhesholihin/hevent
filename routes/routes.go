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
	public.GET("/events/price/:eventid", myHandler.GetEventPrice)

	//Admin Route
	//Backend Management Route by Admin
	admin := apiV1.Group("/admin", middleware.JWTAuthAdmin())
	admin.GET("/profile", myHandler.GetAdminProfile)
	admin.PUT("/updateprofile", myHandler.UpdateAdmin)

	admin.GET("/users", myHandler.GetListUsers)

	admin.POST("/category", myHandler.AddCategory)
	admin.PUT("/category/:categoryid", myHandler.UpdateEventCategory)
	//admin.DELETE("/category/:categoryid", myHandler.RemoveEventCategory)

	admin.POST("/events", myHandler.AddEvent)
	admin.PUT("/events/:eventid", myHandler.UpdateEvent)
	admin.DELETE("/events/:eventid", myHandler.RemoveEvent)

	//events price
	admin.POST("/events/price/:eventid", myHandler.AddEventPrice)

	//testimg payment handler
	admin.GET("/testpayment", myHandler.TestPayment)

	//User Route
	//Apps Management Route by user
	user := apiV1.Group("/user", middleware.JWTAuthUser())
	user.GET("/profile", myHandler.GetUserProfile)
	user.PUT("/updateprofile", myHandler.UpdateUserProfile)

	//Order Ticketing
	user.GET("/users/:userid/cart", myHandler.GetCartSession)
	user.POST("/users/:userid/cart/:sessionid/item", myHandler.AddItemToCart)
	user.GET("/users/:userid/cart/:sessionid/item", myHandler.GetItemsCart)
	user.DELETE("/users/:userid/cart/:sessionid/item/:itemid", myHandler.DeleteItemFromCart)
	user.POST("/users/:userid/order/create/:sessionid", myHandler.CreateOrder)
	user.GET("/users/:userid/order", myHandler.GetListOrders)
	user.GET("/users/:userid/order/:orderid", myHandler.GetOrderById)
	user.PUT("/users/:userid/order/:orderid/payment/:paymentid", myHandler.UploadReceipt)

}
