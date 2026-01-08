package main

import ("fmt"
    "strings"
)

func main() {
 currencyFrom, amount, currencyTo := user()
 result := calculate(currencyFrom, amount, currencyTo)
 fmt.Println(result)
}

func user() (string, float64, string ) {
 var currencyFrom string
 for {
 fmt.Print("Введите исходную валюту (usd, eur, rub): ")
 fmt.Scan(&currencyFrom)
 currencyFrom = strings.ToLower(currencyFrom)
 if currencyFrom == "usd" || currencyFrom == "eur"|| currencyFrom == "rub" {
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
 var currencyTo string
 for {
 fmt.Print("Введите целевую валюту: ")
 fmt.Scan(&currencyTo)
 currencyTo = strings.ToLower(currencyTo)
 if currencyTo == currencyFrom {
	fmt.Println("Целевая и исходная валюты не должны совпадать")
    continue
 }
 if currencyTo == "usd" || currencyTo == "eur" || currencyTo == "rub" {
			break
		}
		fmt.Println("Ошибка, введите доступную валюту")
    }
 return currencyFrom, amount, currencyTo
}

func calculate(currencyFrom string, amount float64, currencyTo string) float64 {
 if currencyFrom == "usd" && currencyTo == "eur" {
  return amount * 0.92
 } else if currencyFrom == "eur" && currencyTo == "usd" {
  return amount * 1.08
 } else if currencyFrom == "rub" && currencyTo == "usd" {
  return amount / 90.0
 } else if currencyFrom == "usd" && currencyTo == "rub" {
  return amount * 90.0
 }
 return 0
 
}