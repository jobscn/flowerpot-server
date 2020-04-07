package startup

import (
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/controller"
)

//func c(f interface{}, methodName string) gin.HandlerFunc {
//
//	fmt.Println("test: ")
//	return func(ctx *gin.Context) {
//		args := make(map[string]interface{})
//		err := ctx.BindUri(args)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("args: ", args)
//
//		in := []reflect.Value{
//			reflect.ValueOf(ctx),
//			reflect.ValueOf(args),
//		}
//
//		reflect.ValueOf(f).MethodByName(methodName).Call(in)
//
//	}
//}

func RegisterGinRouter(app *gin.Engine) {
	userController := controller.UserController{}

	v1 := app.Group("/v1")
	{
		v1.GET("/user/login", userController.Login)
		v1.GET("/user/register", userController.Register)
	}
}

func RegisterDependencyInjector() {

}