package main

import "fmt"

func greeting() { //создание новой функции
	fmt.Println("Привет!")
}

func divider() {
	fmt.Println("----")
}

func printMenu() {
	fmt.Println("1) Добавить")
	fmt.Println("2) Удалить")
	fmt.Println("3) Выйти")
}

func introduceNameAge(name string, age int) { //создаем функцию с переменными
	fmt.Printf("Имя %s, возраст %d\n", name, age) //указываем, что функция выводит, какие данные мы можем внести в переменные
} //и где они находятся при выводе

func summa(a, b int) { //переменные одного типа можно указывать через запятую
	fmt.Println("Сумма равна:", a+b)
}

func infoCity(nameCity string, population int, area float64) {
	fmt.Printf("Город: %s\n", nameCity)
	fmt.Printf("Население: %d\n", population)
	fmt.Printf("Площадь: %.2f\n", area)
}

func personalGreeting(name string, age int) {
	fmt.Printf("Привет, %s! Тебе %d лет\n", name, age)
}

func calculateArea(length, width float64) {
	area := length * width
	fmt.Printf("Площадь: %.2f\n", area)
}

func formatProduct(name string, price float64, quantity int) {
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Цена: %.2f\n", price)
	fmt.Printf("Количествор: %d\n", quantity)
	fullSumm := price * float64(quantity)
	fmt.Printf("Общая сумма: %.2f\n", fullSumm)
}

// Функция square возвращает квадрат числа n
func square(n int) int { // int после параметров - тип возврата
	return n * n
}

func double(x int) int {
	return x * 2
}

func addTen(x int) int {
	return x + 10
}

func exclaim(text string) string {
	return text + "!"
}

func greet(name string) string {
	return "Здравствуйте, " + name
}

func formatPrice(price float64) string {
	return fmt.Sprintf("%.2f руб.", price)
}

func isEven(n int) bool { //чет или нечет
	return n%2 == 0
}

func formatTime(hours, minutes int) string {
	return fmt.Sprintf("%02d:%02d\n", hours, minutes) //нужен в формате 00:00
}

func cube(n int) int { //куб числа n
	return n * n * n
}

func average(a, b float64) float64 { //среднее арифметическое двух чисел
	return (a + b) / 2.0
}

func triplle(n int) int {
	return n * 3
}

func celsiusToFahrenheit(c float64) (float64, float64) { //возращает два значения
	f := c*1.8 + 32 //фаренгейты
	return c, f     //два значения

}

func showBanner(name string) {
	fmt.Println("Привет,", name)
}

func showFooter(name string) {
	fmt.Println("Пока,", name)
}

func Banner(name string) string {
	return "Как дела, " + name + "?"
}

func artBanner() {
	fmt.Println(" /\\_/\\")
	fmt.Println("( o.o )")
	fmt.Println(" > ^ <")
}

func main() {
	fmt.Println("Задание 1:")
	greeting() //вызов функции внутри основной функции
	greeting()
	greeting()

	fmt.Println("")
	fmt.Println("Задание 2")
	fmt.Println("Start")
	divider()
	fmt.Println("Middle")
	divider()
	fmt.Println("End")

	fmt.Println("")
	fmt.Println("Задание 3")
	fmt.Println("Меню:")
	printMenu()
	fmt.Println("---")
	printMenu()
	fmt.Println("Выберите пункт...")

	fmt.Println("")
	fmt.Println("Задание 4")
	introduceNameAge("Даша", 24) //вызываем функцию, передавая ей нужные нам аргументы

	fmt.Println("")
	fmt.Println("Задание 5")
	summa(2, 3)

	fmt.Println("")
	fmt.Println("Задание 6")
	infoCity("Москва", 50000, 29.99)

	fmt.Println("")
	fmt.Println("Задание 7")
	personalGreeting("Анна", 28)

	fmt.Println("")
	fmt.Println("Задание 8")
	calculateArea(21.12, 10.11)
	calculateArea(22.1, 11.1)

	fmt.Println("")
	fmt.Println("Задание 9")
	formatProduct("Яблоки", 29.99, 5)

	fmt.Println("")
	fmt.Println("Задание 10")
	result := square(5)
	fmt.Println("Квадрат числа 5 равен: ", result)
	fmt.Println("Квадрат числа 7:", square(7))

	fmt.Println("")
	fmt.Println("Задание 11")
	doublet := double(3) //возвращенное значение функции можно сохранить в другую переменную
	fmt.Println("Удвоение:", doublet)
	fmt.Println("Прямо в print:", double(3)) //использовать в выражении (проще)
	result = addTen(double(3))               //использовать в другой функции
	fmt.Println("Композиция функций:", result)
	total := double(5) + double(6) //использовать в вычислениях
	fmt.Println("Сумма удвоенных:", total)

	fmt.Println("")
	fmt.Println("Задание 12")
	greeti := greet("Александр")
	fmt.Println(greeti)
	fmt.Println(greet("Саша"))
	fmt.Println(exclaim(greet("Мария")))
	fmt.Println("Цена:", formatPrice(99.99))
	fmt.Println(isEven(5))
	fmt.Println(isEven(4))

	fmt.Println("")
	fmt.Println("Задание 13")
	fmt.Print(formatTime(2, 9))
	fmt.Print(formatTime(20, 20))

	fmt.Println("")
	fmt.Println("Задание 14")
	fmt.Println(cube(3))
	fmt.Println(average(3, 5))

	fmt.Println("")
	fmt.Println("Задание 15")
	fmt.Println(triplle(3))
	celsius, fahrenheit := celsiusToFahrenheit(26) //задаем две переменные, которые принадлежат переменным c и f соответственно в функции и задаем значение переменной, которую принимает эта функция
	fmt.Printf("%.1f в Цельсиях будет %.1f в Фаренгейтах\n", celsius, fahrenheit)

	fmt.Println("")
	fmt.Println("Задание 16")
	showBanner("Даша")
	showFooter("Даша")
	fmt.Println(Banner("Даша"))
	artBanner()
}
