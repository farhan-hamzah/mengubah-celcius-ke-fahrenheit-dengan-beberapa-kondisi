package main
import "fmt"

func tempt(c float64)float64{
	return ((9.0/5.0)*c)+32
}
func main(){
	var celcius, hasil float64
	var i int
	fmt.Scan(&i)
	for i > 0{
		fmt.Scan(&celcius)
		hasil = tempt(celcius)
		fmt.Println(hasil)
		i -= 1
	}
}
