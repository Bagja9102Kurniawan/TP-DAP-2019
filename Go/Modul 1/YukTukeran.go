/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini bertujuan untuk memindahkan input pertama yang awalnya di depan kalimat berubah menjadi di akhir kalimat
	sedangkan input ke 2 berubah menjadi di depan kalimat 
	dan input yang ke 3 menjadi di tengah kalimat
*/
package main
import "fmt"

var a,b,c,d string
func main() {
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&c)
	fmt.Println("Output awal = "+a+""+b+""+c)
	d = a
	a = b
	b = c
	c = d
	fmt.Println("Output akhir = "+a+""+b+""+c)
}