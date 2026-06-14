package main

import "fmt"

func sumToN(n int) int { //счет суммы от 1 до числа n через формулу гаусса
	sum := n * (n + 1) / 2
	return sum
}

func sumToN1(n int) int { //то же самое но с циклом for
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sumEven(n int) int { //вычисление суммы четных чисел от 2 до n
	sum := 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func power(x, n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= x
	}
	return result
}

func sumDigits(n int) int {
	sum := 0
	ost := n % 10
	chislo := n / 10
	for i := 0; i <= n; i++ {
		chislo = chislo / 10
		ost = ost % 10
		sum += ost
	}
	return sum
}

func doubleUntilLimit(start, limit int) int {
	if start >= limit {
		return start
	}
	for start < limit {
		start *= 2
	}
	return start
}

func sumSkipMultiples(limit, skip int) int { //сумма с пропуском кратных
	if limit < 1 || skip <= 1 {
		return 0
	}
	sum1 := 0
	for i := 1; i <= limit; i++ {
		if i%skip != 0 {
			sum1 += i
		}
	}
	return sum1
}

func main() {
	for i := 1; i <= 5; i++ { //для переменной i=1, выполнять цикл, пока i<=5, и прибавлять после каждого шага цикла i+1
		fmt.Println(i)
	}
	fmt.Println("")

	for i := 10; i >= 1; i-- { //уменьшение
		fmt.Println(i)
	}
	fmt.Println("")

	count := 5
	for count > 0 { //цикл for как "пока": пока count>0 выполнять цикл и уменьшать count на 1
		fmt.Printf("Осталось попыток %d\n", count)
		count--
	}
	fmt.Println("")

	value := 1
	limit := 100
	for value < limit { //пока value < limit выполняется цикл
		fmt.Printf("Текущее значение: %d\n", value) //печатается текст с значением
		value *= 2                                  //удваивается переменная
	}
	fmt.Printf("Финальное значение: %d (превысило предел %d)\n", value, limit) //как только удвоенная переменная становиться больше limita,
	fmt.Println("")                                                            // цикл заканчивается и печатается финальная фраза

	sum := 0
	count1 := 0
	fmt.Println("Суммируем числа от 1 до бесконечности, пока сумма не превысит 20")
	for { //бесконечный цикл
		count1++ //ставим условия внутри цикла
		sum += count1
		if sum > 20 { //проверку внутри цикла для выхода из бесконечного цикла
			fmt.Printf("Сумма превысила 20! Просуммировали числа от 1 до %d\n", count1)
			break //выходим из этого цикла, а не из всей функции, как при return
		}
	}
	fmt.Println("")

	for i := 0; i < 10; i++ {
		if i != 6 && i != 7 { //печатаем все числа от 0 до 9 без цифр 6 и 7
			fmt.Println(i)
		}
	}
	fmt.Println("")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println("")

	i := 7 //табл умножения на 7
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d*%d=%d ", i, j, i*j)
		if j%2 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("")

	n := 6
	factorial := 1
	fmt.Printf("Вычисление факториала числа %d\n", n)
	for j := 1; j <= n; j++ {
		factorial *= j
	}
	fmt.Printf("!%d = %d\n", n, factorial)
	fmt.Println("")

	for j := 1; j <= 50; j++ {
		p := j * j
		if p >= 50 {
			fmt.Printf("Первое число, квадрат которого больше 50 это %d = %d\n", j, p)
			break
		}
	}
	p := 1
	for p*p <= 50 {
		fmt.Printf("%d² = %d\n", p, p*p)
		p++
	}
	fmt.Printf("%d² = %d > 50\n", p, p*p)
	fmt.Println("")

	fmt.Println("Треугольник из звездочек :)")
	heigh := 5
	for i := 1; i <= heigh; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("")
	fmt.Println(sumToN(5))

	fmt.Println("")
	fmt.Println(sumToN1(5))

	fmt.Println("")
	fmt.Println(power(2, 2))

	fmt.Println("")
	fmt.Println(sumDigits(12345))

	fmt.Println("")
	fmt.Println(doubleUntilLimit(1, 100))

	fmt.Println("")
	fmt.Println(sumSkipMultiples(10, 3))
}
