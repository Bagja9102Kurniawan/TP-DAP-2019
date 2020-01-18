package main
import "fmt"

var nam float64
var nmk string
func main () {
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nmk = "A"
		}	else if nam > 72.5 {
				nmk = "AB"
		}	else if nam > 65 {
				nmk = "B"
		}	else if nam > 57.5 {
				nmk = "BC"
		}	else if nam > 50 {
				nmk = "C"
		}	else if nam > 40 {
				nmk = "D"
		}	else if nam <= 40 {
				nmk = "E"
		}
	fmt.Println("Nilai mata kuliah: ", nmk)
}