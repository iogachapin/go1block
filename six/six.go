package main

import "fmt"

func sumToLimit(limit int) {
	sum2 := 0
	fmt.Println("Суммируем числа до превышения лимита")
	for i := 1; i <= 100; i++ {
		sum2 += i
		if sum2 > limit {
			fmt.Printf("Превысили лимит %d! Выходим из цикла\n", limit)
			break
		}
	}
	fmt.Printf("Итоговая сумма: %d\n", sum2)
}

func forContinue(limit int) {
	fmt.Println("Печать только нечётных чисел от 1 до 10:")
	for i := 1; i <= limit; i++ {
		if i%2 == 0 {
			continue //пропускает ВСЕ четные числа
		}
		fmt.Printf("%d ", i) //выдает числа в одну строчку с пробелами
	}
	fmt.Println("") //делает строчку в конце
}

func rightContinue(limit int) {
	i := 0
	for i < limit {
		if i == 5 {
			i++
			continue
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println("")
}

func breakOuter(limit int) {
	fmt.Printf("Ищем пару чисел, произведение которых будет больше %d\n", limit)
outer:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			product := i * j
			if product > limit {
				fmt.Printf("Нашли: %d * %d > %d\n", i, j, limit)
				break outer
			}
		}
	}
	fmt.Println("Поиск завершен")
}

func continueRows() {
	fmt.Println("Пропускаем строку с цифрой 2")
rows:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 {
				continue rows //continue rows - следующая итерация цикла rows
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
}

func fizzBizzGame(limit int) {
	fmt.Printf("Игра FizzBizz для чисел от 1 до %d\n", limit)
	for i := 1; i <= limit; i++ {
		if i%15 == 0 {
			fmt.Printf("%d: FizzBizz\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("%d: Bizz\n", i)
			continue
		}
		if i%3 == 0 {
			fmt.Printf("%d: Fizz\n", i)
			continue
		}
		fmt.Printf("%d: %d\n", i, i)
	}
}

func earlyBreak(number int) {
	fmt.Printf("Ищем первый делитель числа %d (кроме 1 и самого числа):\n", number)
	for i := 1; i < number; i++ {
		if number%i == 0 {
			fmt.Printf("Число %d делится на %d без остатка\n", number, i)
			break
		}
	}
}

func sumOneTarget(targetSum int) {
	fmt.Printf("Ищем первую пару чисел, сумма которых равна %d\n", targetSum)
search:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			sum := i + j
			if sum == targetSum {
				fmt.Printf("Нашли пару чисел: %d и %d в сумме дают %d\n", i, j, targetSum)
				break search // выход при первом совпадении. если написать просто break, то найдутся все числа, которые в сумме дают 15
			}
		}
	}
}

func shouldSkip() { //вывести на экран все числа, кроме 6 и 7
	for i := 0; i <= 10; i++ {
		if i == 6 || i == 7 {
			continue
		}
		fmt.Println(i)
	}
}

func onlyOdd(namber int) {
	fmt.Printf("Вывести только нечетные числа от 1 до %d:\n", namber)
	for i := 1; i <= namber; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}

func sumAndBreak(limit int) {
	fmt.Printf("Суммируем все числа от 1, пока сумма не превысит %d:\n", limit)
	sum := 0
	lastNamber := 0
	for i := 1; ; i++ { //пропуск обязателен, так как означает, что условие всегда true
		sum += i
		lastNamber = i
		if sum > limit {
			break
		}
	}
	fmt.Printf("Сумма равна: %d\n", sum)
	fmt.Printf("Последнее добавленное число: %d\n", lastNamber)
}

func firstTarget(n, limit int) {
exit:
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			proiz := i * j
			if proiz > limit {
				fmt.Printf("Нашли первую пару чисел, произведение которых больше %d: %d*%d = %d\n", limit, i, j, proiz)
				break exit
			}
		}
	}

}

func enumeration() {
	for n := 1; n <= 100; n++ {
		if n%7 == 0 && n%3 == 1 {
			fmt.Printf("Нашли первое число, которое делится без остатка на 7 и при делении на 3 дает остаток 1: n = %d\n", n)
			break
		}
	}
}

func searchFiveNumbers() {
	sum := 0
	for n := 1; n <= 100; n++ {
		if n%3 != 0 || n%2 == 0 {
			continue
		}
		sum += 1
		if sum == 5 {
			fmt.Println(sum)
			break
		}
	}
}

func targetProiz() {
search:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == 45 {
				fmt.Printf("Нашли первую пару чисел, произведение которых равно 45, это %d и %d\n", i, j)
				break search
			}
		}
	}
}

func onlyNotOdd() {
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func sumToLimit2(limit int) {
	sum := 0
	lastNumber := 0
	for i := 1; ; i++ {
		sum += i
		lastNumber = i
		if sum > limit {
			fmt.Printf("Сумма чисел превысила лимит в %d. Последнее добавленное число = %d, сумма равна %d\n", limit, lastNumber, sum)
			break
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Printf("%d: FizzBuzz\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("%d: Buzz\n", i)
			continue
		}
		if i%3 == 0 {
			fmt.Printf("%d: Fizz\n", i)
			continue
		}
		fmt.Printf("%d: %d\n", i, i)
	}
}

func divisorN(n int) {
	for divisor := 2; ; divisor++ {
		if n%divisor == 0 {
			fmt.Printf("Делитель числа %d равен %d\n", n, divisor)
			quotient := n / divisor
			fmt.Printf("Частное от деление равно %d\n", quotient)
			break
		}
	}

}

func numbersGame(n int) {
	step := 0
	for n != 1 {
		if step > 1000 {
			fmt.Println("Счетчик превысил 1000. Выходим из цикла")
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		step += 1
		fmt.Println(n)
	}
	fmt.Printf("Число стало равно %d. Счетчик: %d\n", n, step)
}

func main() {
	sumToLimit(50)
	fmt.Println(" ")

	forContinue(10)
	fmt.Println(" ")

	rightContinue(10)
	fmt.Println(" ")

	breakOuter(50)
	fmt.Println(" ")

	continueRows()
	fmt.Println(" ")

	fizzBizzGame(10)
	fmt.Println(" ")

	earlyBreak(91)
	fmt.Println(" ")

	sumOneTarget(15)
	fmt.Println(" ")

	shouldSkip()
	fmt.Println(" ")

	onlyOdd(10)
	fmt.Println(" ")

	sumAndBreak(100)
	fmt.Println(" ")

	firstTarget(10, 40)
	fmt.Println(" ")

	enumeration()
	fmt.Println(" ")

	searchFiveNumbers()
	fmt.Println(" ")

	targetProiz()
	fmt.Println(" ")

	onlyNotOdd()
	fmt.Println(" ")

	sumToLimit2(100)
	fmt.Println(" ")

	fizzBuzz()
	fmt.Println(" ")

	divisorN(91)
	fmt.Println(" ")

	numbersGame(27)

}
