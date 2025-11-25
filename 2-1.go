package main
import"fmt"
func main(){
var a float64
fmt.Scan(&a)
switch {
case a >=6.5 && a<=8.6 :
fmt.Println("Air layak minum")
case a<0 || a>14 :
fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
default:
fmt.Println("Air tidak layak minum")
}
}
