/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM Fungsi FoGoH
*/
package main
import "fmt"

var x float32
func main () {
	fmt.Print("Masukkan nilai X :")
	fmt.Scanln(&x)
	fmt.Println("f(",x,")=",fungsiF(x))
	fmt.Println("g(",x,")=",fungsiG(x))
	fmt.Println("h(",x,")=",fungsiH(x))
	fmt.Println("fungsi fogoh(",x,")=",fogoh(x))
	fmt.Println("\nNAMA	: BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM	: 1301194020")
}
func fungsiF (x float32) float32 {
	return x*x
}
func fungsiG (x float32) float32 {
	return x-2
}
func fungsiH (x float32) float32 {
	return x+1
}
func fogoh (x float32) float32 {
	return fungsiF(fungsiG(fungsiH(x)))
}