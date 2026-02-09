package main

import ("fmt"
    "strings"
)

func main() {
 from, amount, to := user()
 result := calculate(from, amount, to)
 fmt.Println(result)
}

func user() (string, float64, string ) {
 var from string
 for {
 fmt.Print("Введите исходную валюту (usd, eur, rub): ")
 fmt.Scan(&from)
 from = strings.ToLower(from)
 if from == "usd" || from == "eur"|| from == "rub" {
     break
 }
    fmt.Println("Ошибка, введите доступную валюту")
   
 }
 var amount float64
 for {
 fmt.Print("Введите сумму для конвертации:  ")
 fmt.Scan(&amount)
 if amount >= 0 {
    break
   } 
    fmt.Println("Введите положительное число")
    
 }
 var to string
 for {
 fmt.Print("Введите целевую валюту: ")
 fmt.Scan(&to)
 to = strings.ToLower(to)
 if to == from {
	fmt.Println("Целевая и исходная валюты не должны совпадать")
    continue
 }
 if to == "usd" || to == "eur" || to == "rub" {
			break
		}
		fmt.Println("Ошибка, введите доступную валюту")
    }
 return from, amount, to
}

func calculate(from string, amount float64, to string) float64 {
currencyRUB := map[string]float64{
   "rub": 1,
   "usd": 77.0,
   "eur": 91.70,
}
 result := amount * currencyRUB[from] / currencyRUB[to]
       
return result

}




