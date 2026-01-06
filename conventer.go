package main

import "fmt"

func main() {
 currencyFrom, amount, currencyTo := user()
 result := calculate(currencyFrom, amount, currencyTo)
 fmt.Println(result)
}

func user() (string, float64, string ) {
 var currencyFrom string
 fmt.Print("Введите исходную валюту (usd, eur, rub): ")
 fmt.Scan(&currencyFrom)
 var amount float64
 fmt.Print("Введите сумму для конвертации:  ")
 fmt.Scan(&amount)
 var currencyTo string
 fmt.Print("Введите целевую валюту: ")
 fmt.Scan(&currencyTo)
 if currencyTo == currencyFrom {
	fmt.Println("Целевая и исходная валюты не должны совпадать")
	return currencyFrom, amount, currencyTo
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