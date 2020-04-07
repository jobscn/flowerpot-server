package main

import (
	"fmt"
	"jobscn/ai-flower-pot/controller"
	"reflect"
	"runtime"
)
type TestStruct struct {
	Name string `json:"Name"`
	Num int64 `json:"Num"`
}

func Test(a *TestStruct) string {
	fmt.Println("this is Test method")
	return "this.Name"
}

func (p *TestStruct)Test2(param map[string]interface{}) string {
	fmt.Printf("this is Test method %+v\n", param)
	return "this.Name"
}


func main() {
	userController := controller.UserController{}
	testxxx := userController.Login
	testvvv := runtime.FuncForPC(reflect.ValueOf(testxxx).Pointer()).Name()
	fmt.Printf("%+v", testvvv)

	// 传参obj + 被执行方法名

	//inter := &TestStruct{}
	//m := []byte(`{"Name": "name1","Num": 100}`)
	//mmm := make(map[string]interface{})
	//err := json.Unmarshal(m, &mmm)
	//if err != err {
	//	panic(err)
	//}
	//fmt.Println("test: ", mmm)
	//
	//in := []reflect.Value{reflect.ValueOf(mmm)}
	//
	//fmt.Println(reflect.ValueOf(inter).MethodByName("Test2").Call(in))
}