/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM PENGHITUNG NILAI PI DI SUKU KE-N
*/
package main
import "fmt"

var A,B float64
var n int
func main () {
	fmt.Print("N suku pertama :")
	fmt.Scan(&n)
	i := 1
	k := 1.0
	for i <= n {
		if i % 2 == 0 {
			A -= 1.0/k
		}else {
			A += 1.0/k
		}	
		k = k + 2.0		
		i++
	}	
	B = A*4
	fmt.Println("Hasil PI:",B)
	fmt.Println("\nNAMA	: BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM	: 1301194020")
}