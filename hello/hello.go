package hello
import "fmt"
import "github.com/yixy/go-version-dojo/version"
func SayHello(){
	fmt.Println("hello-v1.0.0",version.EchoVersion())
}
