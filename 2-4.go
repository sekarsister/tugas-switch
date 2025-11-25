package main
import "fmt"
func main(){
var a int
var c string

fmt.Scanln(&a)
fmt.Scanln(&c)
switch{
case c == "USD":
fmt.Println(a/16000)
case c == "JPY":
fmt.Println(a/110)
case c == "EUR":
fmt.Println(a/17500)
}
}
