/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM pencarian nilai maximum, minimum, atau mencari suatu nilai tertentu di antara sekumpulan
*/
package main
import "fmt"

const N=2019
type RecType struct {
	f1 string
	f2 int
	f3 float64
	}
type ArrType[N]RecType

func rmax(data ArrType)float64 {
	var max float64

	max = data[0].f3
	i := 0
	for i < len(data) {
		if data[i].f3 > max {
			max = data[i].f3
		} 
		i++
	}
	return max
}

func imin(data ArrType)int {
	var min, idx int

	min = data[0].f2
	i := 0
	for i < len(data) {
		if data[i].f2 < min && data[i].f2 != 0 {
			min = data[i].f2
			idx = i
		} 
		i++
	}
	return idx
}

func found(data ArrType,key string)bool{

	i := 0
	for i < len(data) {
		if data[i].f1  == key {
			return true
		}
		i++
	}
	return false
}

func pos(data ArrType,key string)int{
	var indexkiri, indexkanan, med int

	indexkiri = 0
	indexkanan = N
	for indexkiri < indexkanan {
		med = (indexkiri + indexkanan)/2
		if data[med].f1 < key {
			indexkiri = med + 1
		}else if data[med].f1 > key {
			indexkanan = med
		}else {
			return med
		}
	}
	return -1
}

func main() {
	fmt.Println("\n\nNAMA\t:\tBAGJA 9102 KURNIAWAN")
	fmt.Println("NIM\t:\t1301194020")
}