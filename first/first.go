package main                    // объявляем главный пакет, всегда пишется первым

import "fmt"                   //подключаем инструмент для вывода текста - пакет "формат" от слова форматирование

    var name string = "Даша"  //очевидное объявление переменной типа строка, МОЖНО ИСПОЛЬЗОВАТЬ ВНЕ ФУНКЦИЙ!!
    var age int = 24          //очевидное объявление целочисленной переменной
    
    var (
        shool = "Kata Academy"  //явное объявление переменных блоком
        language = "Golang"
    )

func main() {                //создаём главную функцию

    random := 100            //неявное объявление переменной int, программа сама понимает ее тип, ИСПОЛЬЗУЕТСЯ ТОЛЬКО ВНУТРИ ФУНКЦИИ!!
    city := "Москва"        //неявное объявление строки
    drob := 99.99          //неявное объявление дроби
    isStudent := true         //неявное объявление типа bool (да/нет)

    x, y, z := 10, 12, 14   //короткое объявление сразу нескольких переменных
    name2, name3 := "Комарова", "Алексеевна"

    fmt.Println("Привет, мир!") //выводим текст на экран
    fmt.Println("Учим golang емае")
    fmt.Println("")             //пустая строка
    fmt.Println("Меня зовут", name, "Мне", age)
    fmt.Println("Вывод числа, города, дроби и студентка ли я:", random, city, drob, isStudent)
    fmt.Println("Просто координаты:", x,y,z) 
    fmt.Println("Мои фамилия и отчество:", name2, name3)
    fmt.Println("")
    fmt.Println("Теперь сделаем задание на описание товара и его цены:")

        //задание на описание товара и его цены

        productName := "Ноутбук"
        price := 20.20
        quantity := 10                              //количество
        inStock := true

        const discont = 0.15                        //15% скидка
        discountAmount := price * discont           //размер скидки
        finalPrice := price - discountAmount        //окончательная цена
        
        fmt.Println("Товар:", productName)
        fmt.Println("Цена:", price)
        fmt.Println("Скидка:", discountAmount)
        fmt.Println("Со скидкой:", finalPrice)
        fmt.Println("Количество:", quantity)
        fmt.Println("В наличии:", inStock)

        //очередное задание с printf и спецификаторами форматирования
        //printf используется для форматирования текста, т.е. не просто тупой вывод, а приведение его к какому-то нужному виду

            fmt.Println("")
            n := 25
            p := 99.99
            user := "Alice"
            ok := true

            fmt.Printf("%v (%T)\n", n, n)
            fmt.Printf("%v (%T)\n",p, p)
            fmt.Printf("%v (%T)\n", user, user)
            fmt.Printf("%v (%T)\n", ok, ok)

                //еще одно задание

                distance := 15
                pricePerKm := 25.50

                cost := pricePerKm * float64(distance)          //перевод целочисленной переменной в дробную

                fmt.Printf("Distance: %d km\n", distance)       //%d - для целых чисел
                fmt.Printf("Price per km: %.2f\n", pricePerKm)  //%.2f - для дробей, два знака после запятой, указывается обязательно сама переменная
                fmt.Printf("Delivery cost: %.2f\n", cost)

                //и еще!!

                    fmt.Println("")
                    taxPercent := 13
                    currency := "RUB"
                    price, qty := 1500.00, 2            //объявление нескольких переменных в одной строке

                    total := price * float64(qty)      
                    tax := total * float64(taxPercent) / 100
                    final := total + tax

                    fmt.Printf("Price: %.2f\n", price)                  
                    fmt.Printf("Qty: %d\n", qty)                        
                    fmt.Printf("Tax: %.2f\n", tax)          
                    fmt.Printf("Final: %.2f %s\n", final, currency)     //%s - для строк. два спецификатора формата - две переменные после

}