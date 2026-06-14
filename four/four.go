package main

import "fmt"

func сashInBalance(balance float64, price float64) { //функция проверки баланса
	if balance >= price {
		fmt.Println("Покупка одобрена")
		fmt.Printf("Ваш баланс: %.2f\n", balance-price)
	} else {
		fmt.Println("Денег на вашем счету не хватает")
		fmt.Printf("Ваш баланс: -%.2f\n", price-balance)
	}
}

func rating(score int) { //функция проверки балла для оценки
	if score >= 90 {
		fmt.Printf("Ваш бал: %d\n", score)
		fmt.Println("Это оценка: А (отлично)")
	} else if score >= 75 {
		fmt.Printf("Ваш бал: %d\n", score)
		fmt.Println("Это оценка: В (хорошо)")
	} else if score >= 60 {
		fmt.Printf("Ваш бал: %d\n", score)
		fmt.Println("Это оценка: C (удовлетворительно)")
	}
}

func permission(age int, hasPermission bool) { //функция разрешения доступа
	if age < 18 || !hasPermission {
		fmt.Println("Доступ запрещен")
	} else {
		fmt.Println("Доступ разрешён")
	}
}

func benefits(age int) { //функция определения наличия льготы
	if age <= 14 || age >= 65 {
		fmt.Println("Льготный билет")
	} else {
		fmt.Println("Обычный билет")
	}
}

func weekend(isWeekend bool) { //функция определение выходного
	if !isWeekend {
		fmt.Println("Сегодня рабочий день")
	} else {
		fmt.Println("Сегодня выходной!!")
	}
}

func priceStudent(age int, hasPermission, isWeekend bool) { //студенческая скидка
	if (age >= 18 && age <= 25) && (hasPermission || isWeekend) {
		fmt.Println("Студенческая скидка доступна")
	} else {
		fmt.Println("Студенческая скидки нет")
	}
}

func permission2(age int, hasPermission bool) bool { //функция возвращает булево значение - тру или фальш
	return age >= 18 && hasPermission
}

func delZero(den, num int) {
	if den != 0 && num/den > 1 {
		fmt.Println("Результат больше 1")
	} else {
		fmt.Println("Деление на ноль предотвращено")
	}
}

func remains(num1, num2 int) {
	if remainder := num1 % num2; remainder == 0 {
		fmt.Printf("%d делится без остатка на %d\n", num1, num2)
	} else {
		fmt.Printf("%d делится с остатком %d на %d\n", num1, remainder, num2)
	}
}

func chekLVL(score int) {
	if score >= 90 {
		fmt.Println("Ваш уровень 10!")
	} else if score >= 70 {
		fmt.Println("Ваш уровень 9!")
	} else if score >= 50 {
		fmt.Println("Ваш уровень 8!")
	}
}

func delOst(num1 int) {
	if del := num1 % 2; del == 0 {
		fmt.Printf("%d четное\n", num1)
	} else {
		fmt.Printf("%d нечетное\n", num1)
	}
}

func hasDostup(hasCard, pinCorrect, isEmergency bool) {
	if (hasCard && pinCorrect) || isEmergency {
		fmt.Println("Доступ разрешён")
		if hasCard && !pinCorrect {
			fmt.Println("Внимание: экстренный доступ без пин")
		}
	} else {
		fmt.Println("Доступ запрещён")
		if hasCard && !pinCorrect {
			fmt.Println("Причина: неверный PIN")
		}
		if !hasCard && pinCorrect {
			fmt.Println("Причина: нет карты")
		}
	}
}

func leapYear(year int) {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d - високосный год\n", year)
		fmt.Println("В феврале 29 дней")
	} else {
		fmt.Printf("%d - обычный год\n", year)
		fmt.Println("В феврале 28 дней")
	}
}

func hasDelivery(weight int, city string) {
	if weight <= 100 && (city == "Москва" || city == "Санкт-Петербург") {
		fmt.Println("Доставка возможна")
		if city == "Москва" {
			fmt.Println("Доставка в течение 1-2 дней")
		}
		if city == "Санкт-Петербург" {
			fmt.Println("Доставка в течение 3-4 дней")
		}
	} else {
		fmt.Println("Доставка невозможна")
		if weight > 100 {
			fmt.Printf("Вес посылки слишком большой: %d больше 100 кг\n", weight)
		} else {
			fmt.Printf("Доставка в город %s не осуществляется\n", city)
		}
	}
}

func main() {

	fmt.Println("Задание 1: баланс на счету")
	сashInBalance(2000.00, 1000.00)

	fmt.Println("")
	fmt.Println("Задание 2: расчет оценки по баллам")
	rating(85)

	fmt.Println("")
	fmt.Println("Задание 3: разрешение")
	permission(20, true)

	fmt.Println("Или то же самое:")
	if permission2(16, true) { //функция принимает два значения, делает расчет, возвращает значение. Первая строчка всегда по умолчанию тру, и сравнивать дополнительно с тру не нужно
		fmt.Println("Доступ разрешён")
	} else { //вторая автоматически фальш
		fmt.Println("Доступ запрещён")
	}

	fmt.Println("")
	fmt.Println("Задание 4: льгота на билет")
	benefits(68)

	fmt.Println("")
	fmt.Println("Задание 5: выходной или будний день")
	weekend(true)

	fmt.Println("")
	fmt.Println("Задание 6: скидка для студентов")
	priceStudent(20, true, true)

	fmt.Println("")
	fmt.Println("Задание 7: деление на ноль")
	delZero(0, 10)

	fmt.Println("")
	fmt.Println("Задание 8: деление с остатком")
	remains(17, 5)

	fmt.Println("")
	fmt.Println("Задание 9: лвл по баллам")
	chekLVL(85)

	fmt.Println("")
	fmt.Println("Задание 10: чет или нечет")
	delOst(5)

	fmt.Println("")
	fmt.Println("Задание 11: есть ли доступ")
	hasDostup(true, false, false)

	fmt.Println("")
	fmt.Println("Задание 12: високосный ли год")
	leapYear(2026)

	fmt.Println("")
	fmt.Println("Задание 13: возможна ли доставка")
	hasDelivery(90, "Москва")
	hasDelivery(120, "Москва")
	hasDelivery(90, "Пермь")
}
