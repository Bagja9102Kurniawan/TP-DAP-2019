/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini adalah program pengecekan selisih 
	berat benda yang akan dibawa pada motor.
	motor akan oleng ke satu bagian jika 
	selisihnya lebih dari sama dengan 9kg dengan output
	bernilai true(motor oleng), 
	program akan berhenti jika isi kedua kantong melebihi 150kg 
	atau salah satu kantong bernilai negatif.
*/
package main
import "fmt"

var i,c float64
func main () {
	for (i+c<150)&&(i>=0&&c>=0) {
		fmt.Print("Masukkan berat belanjaan di kedua kantong : ")
		fmt.Scanln(&i, &c)
		fmt.Println("Sepeda motor pak Andi akan oleng          :",i-c>=9||c-i>=9)
	}
	fmt.Println("Proses selesai.")
	fmt.Println("NAMA : BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM  : 1301194020")
	fmt.Scanln( )
}