package main
import "fmt"
func main() {
const usd_to_eur float64 = 0.8686
const usd_to_rub float64 = 80.5
const eur_to_rub float64 = ((1 / usd_to_eur) * usd_to_rub)
fmt.Println(eur_to_rub)

	
}