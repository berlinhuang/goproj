package myfunc
import "fmt"

var str = "Hello Go"
var File = "lib.go"
// 公有函数的名字以大写字母开头
func SayHello(){
	fmt.Println(str)
}
//私有函数的名字以小写字母开头
func sayHi(){
	fmt.Println("hi")
}

