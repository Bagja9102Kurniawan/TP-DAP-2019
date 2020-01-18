/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini adalah program yang menerima 
	input 2 buah bilangan real positif
	yang akan terus meminta input 
	hingga salah satunya bernilai 9 atau lebih
*/
package main
import "fmt"

var i,c float64
func main () {
	for (i<9)&&(c<9) {
		fmt.Print("Masukkan berat belanjaan di kedua kantong : ")
		fmt.Scanln(&i, &c)
	}
	fmt.Println("Proses selesai.")
	fmt.Println("NAMA : BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM  : 1301194020")
	fmt.Scanln( )
}