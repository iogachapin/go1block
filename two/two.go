package main

import "fmt"

func CInF(celsius float64) {
	fmt.Printf("Значение в Цельсиях: %.2f\n", celsius)
	fahrenheit := celsius*(9/5) + 32

	fmt.Printf("В Фаренгейтах это: %.1f°C = %.1f°F\n", celsius, fahrenheit)
}

func purchaseReceipt(product string, price float64, quantity int) float64 { //это функция говорит, что она возвращает какое-то значение типа дроби
	fmt.Printf("Куплено: %s\n", product) //указываем значения переменных
	fmt.Printf("По цене: %.2f\n", price)
	fmt.Printf("В количестве: %d\n", quantity)

	return price * float64(quantity) //находим значение, которое функция возращает (return)

}

func main() {

	fmt.Println("Задание 1:")
	age := 24 //калькулятор возраста
	daysInYear := 365
	ageInYear := age * daysInYear
	fmt.Printf("Мне %d года, это %d в днях\n", age, ageInYear)

	fmt.Println("")
	fmt.Println("Задание 2:")
	CInF(20)

	fmt.Println("")
	fmt.Println("Задание 3:")
	total := purchaseReceipt("Ноутбук", 99.99, 2) //total  получает значение возращаемое функцией
	fmt.Printf("Итого сумма: %.2f\n", total)
}
