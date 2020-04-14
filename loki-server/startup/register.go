package startup

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki-server/controller"
	dao_impl "jobscn/ai-flower-pot/loki-server/dao/impl"
	service_impl "jobscn/ai-flower-pot/loki-server/service/impl"
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
	var userController controller.UserController

	v1 := app.Group("/v1")
	{
		// relative to user controller
		{
			v1.GET("/user/login", userController.Login)
			v1.GET("/user/register", userController.Register)
		}
	}

	// register controller to injector
	RegisterDependencyInjector(
		&inject.Object{Value: &userController},
		&inject.Object{Value: &service_impl.UserService{}},
		&inject.Object{Value: &dao_impl.UserDao{}},
	)
}

func RegisterDependencyInjector(v ...*inject.Object) {
	graph := inject.Graph{}

	if err := graph.Provide(v...); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
}
