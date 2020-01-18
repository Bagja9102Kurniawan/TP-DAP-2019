/*	
	NAMA	: BAGJA 9102 KURNIAWAN
	NIM		: 1301194020
	PROGRAM pengurutan dalam array
*/
package main
import "fmt"

const MAXSIZE = 20192020
type RecType struct {
  count int
  size int
}
type ArrType[MAXSIZE]RecType

func iSort(tab *ArrType, nsize int)  {
//I.S. Array tab berisi elemen RecType sebanyak nsize
//F.S. Array tab terurut membesar terhadap field count
  var i, j, t int

  i = 1
  for i < nsize {
    t = tab[i].count
    j = i - 1
    for j >= 0 && tab[j].count > t {
      tab[j+1].count = tab[j].count
      j = j - 1
    }
    tab[j+1].count = t
    i++
  }
}

func mSort(tab *ArrType, nsize int) {
//I.S. Array tab berisi elemen RecType sebanyak nsize
//F.S. Array tab terurut mengecil terhadap field size

  i := 0
  for i < nsize {
    max := i
    j := i
    for j < nsize {
      if tab[j].size > tab[max].size {
        max = j
      }
      j++
    }
    t := tab[max].size
    tab[max].size = tab[i].size
    tab[i].size = t
    i++
  }
}

func isFound(tab ArrType, v, n int)bool {
//I.S. Array tab berisi elemen RecType sebanyak n nilai yang dicari dalam parameter v
//F.S. Mengembalikan nilai true jika nilai v ada dalam field count di salah satu elemen tab, atau false sebaliknya.

	l := 0
	r := n
	m := (l+r)/2
	found := false
	for l < r && !found {
		m = (l+r)/2
		if tab[m].count < v {
			l = m + 1
		} else if tab[m].count > v {
			r = m - 1
		} else if tab[m].count == v {
			found = true
		}
	}
	return found
}

func main() {
	fmt.Println("\nNAMA\t:\tBAGJA 9102 KURNIAWAN")
	fmt.Println("NIM\t:\t1301194020")
}