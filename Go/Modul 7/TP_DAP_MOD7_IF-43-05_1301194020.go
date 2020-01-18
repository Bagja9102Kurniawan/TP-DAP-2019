/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM menghitung jarak titik yang terdekat
*/

package main
import "fmt"
import "math"

type Point struct {
	x float64
	y float64
}

const N = 10000
var points[N]Point
var numpoints int = 0

/*
	fungsi jarak yang menerima dua parameter formal input (hanya-baca) p1 dan
	p2 dengan tipe Point, yang mengembalikan sebuah nilai real yaitu jarak antara kedua
	titik p1 dan p2.
*/
func Jarak(p1 Point, p2 Point)float64{
	var nilaiX, nilaiY float64

	if p1.x > p2.x {
		nilaiX = p1.x - p2.x
	}else{
		nilaiX = p2.x - p1.x
	}
	
	if p1.y > p2.y {
		nilaiY = p1.y - p2.y
	}else{
		nilaiY = p2.y - p1.y
	}
	
	return math.Sqrt((nilaiX)*(nilaiX)+(nilaiY)*(nilaiY))
}

/* 
	prosedur bacaTitik membaca semua titik dari input dan menyimpannya
	dalam array global points. Prosedure tersebut tidak mempunyai parameter formal. Setiap
	baris input berisi dua bilangan real, dengan penanda akhir input pasangan data 0.0.
	Jumlah titik yang terbaca dari input harus dicatat dalam variabel numpoints.
*/
func bacaTitik(){
	var bil Point

	fmt.Scanln(&bil.x, &bil.y)
	for bil.x != 0.0 || bil.y != 0.0 {
		points[numpoints] = bil
		numpoints = numpoints + 1
		fmt.Scanln(&bil.x, &bil.y)
	}
}

/*
	prosedur ambilTitikTerdekat yang mempunyai dua parameter formal output
	p1 dan p2 dengan tipe Point. Prosedur tersebut akan memeriksa jarak setiap pasang titik
	dalam array points dan mengembalikan pasangan titik yang mempunyai jarak terdekat.
*/
func ambilTitikTerdekat(p1*Point, p2*Point){
	var bil1, bil2 int
	var min,rec float64
	var terdekat float64

	bil2 = 0
	for bil2 < numpoints{
		bil2 = bil2 + 1
	}
	
	bil1, bil2 = 0, 1
	rec = Jarak(points[bil1],points[bil2])
	*p1 = points[bil1]
	*p2 = points[bil2]

	for bil1 < numpoints -1 {
	
		if Jarak(points[bil1], points[bil2]) < rec {
			rec = Jarak(points[bil1],points[bil2])
			*p1 = points[bil1]
			*p2 = points[bil2]
		}
		bil2 = bil2 + 1
		
		if bil2 == numpoints{
			bil1 = bil1 + 1
			bil2 = bil1 + 1
		}
	}
	
	bil1 = 0
	bil2 = bil1 + 1
	min = Jarak(points[bil1], points[bil2])
	
	for bil1 < numpoints - 1 {
	
		for bil2 < numpoints {
			terdekat = Jarak(points[bil1], points[bil2])
			
			if terdekat < min {
				min = Jarak(points[bil1], points[bil2])
			}
			
			bil2 = bil2 + 1
		}
		
		bil1 = bil1 + 1
		bil2 = bil1 + 1
	}
}

func main(){
	var titik1, titik2 Point
	bacaTitik()
	ambilTitikTerdekat(&titik1, &titik2)
	
	fmt.Print("Titik terdekat adalah (",titik1.x,".",titik1.y,") dan (",titik2.x,".",titik2.y,") dengan jarak ", Jarak(titik1, titik2))
	fmt.Println("\n\nNAMA	: BAGJA 9102 KURNIAWAN")
	fmt.Println("NIM	: 1301194020")
}