package main
import "fmt"

func main() {
usd_to_eur, usd_to_rub := user()
 eur_to_rub  := calculate(usd_to_eur, usd_to_rub)
	fmt.Println(eur_to_rub)
}

func user() (float64, float64) {
	var usd_to_eur, usd_to_rub float64
	fmt.Print("Введите курс USD->EUR: ")
	fmt.Scan(&usd_to_eur)
	fmt.Print("Введите курс USD->RUB: ")
	fmt.Scan(&usd_to_rub)
return usd_to_eur, usd_to_rub

}

func calculate(usd_to_eur, usd_to_rub float64) float64{
	
	var eur_to_rub float64 = ((1 /usd_to_eur) * usd_to_rub)
	return eur_to_rub

}