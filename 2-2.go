package main 
import "fmt"
func main(){
var a,b int
var c string
fmt.Scan(&c)
fmt.Scan(&a)
switch {
case c == "motor" :
b = a*2000
case c == "mobil" :
b= a*5000
case c == "truk" :
b= a*8000
}
fmt.Println("Rp ",b)
}
