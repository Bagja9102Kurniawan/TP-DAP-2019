/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM PENGHITUNG NILAI PI
	HINGGA SELISIH PI = 0.0001
*/
package main
import "fmt"

var A,B,C float64
func main () {
	i := 1
	k := 1.0
	selisih := 0.0
	for selesai:=false; !selesai;{
		B = C 
			if i % 2 != 0 {
				A += 1.0/k
			}else {
				A -= 1.0/k
			}
			k = k + 2.0
		C = A * 4
		selisih = B - C
		if selisih < 0 {
			selisih = selisih * -1
		}
		selesai =  selisih < 0.00001
		i++
	}
	fmt.Println("Hasil PI:",B)
	fmt.Println("Hasil PI:",C)
	fmt.Println("Pada i ke:",i)
	fmt.Println("\nNAMA	: BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM	: 1301194020")
}