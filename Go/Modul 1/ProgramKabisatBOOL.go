/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini adalah program pengecekan tahun kabisat yaitu tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100
*/
package main
import (
	"fmt"
)
var i int
var kabisat bool
func main () {
	fmt.Println("Program Tahun Kabisat")
	fmt.Print("Tahun         : ")
	fmt.Scanln(&i)
	kabisat = i%400 == 0 || i%4 == 0 && i%100 != 0
	fmt.Print("Kabisat       : ")
	fmt.Println(kabisat)
}