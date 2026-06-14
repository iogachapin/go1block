package main

import (
	"fmt"
)

func massive() { //функция массива
	var arr1 [5]int //[5]int - является типом целиком, то есть в типе указывается длина массива
	var arr2 [5]int

	arr1[0] = 10 //добавляем первой позиции в массиве значение 10
	arr2[0] = 20

	fmt.Printf("arr1: %v, тип: %T\n", arr1, arr1) //%v - значение(например булево); %T - тип
	fmt.Printf("arr2: %v, тип: %T\n", arr2, arr2)
}

func slice() { //функция среза
	var slice1 []int               //нулевой срез, который не имеет значений и не имеет длины
	slice2 := []int{1, 2, 3, 4, 5} //срез принимающий такие значения, но имеют один тип

	fmt.Printf("slice1 (nil): %v, тип: %T\n", slice1, slice1)
	fmt.Printf("slice2: %v, тип: %T\n", slice2, slice2)
}

func zero() { //нулевая инициализация
	var zeroArr [3]int  //[0, 0, 0]	//нулевой массив
	var zeroSlice []int //nil(нулевой срез)

	fmt.Printf("Нулевой массив: %v, тип: %T\n", zeroArr, zeroArr)
	fmt.Printf("Нулевой срез: %v, тип (nil): %T\n", zeroSlice, zeroSlice)
}

func comparisonMassive() { //сравнение массивов
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1 == arr2: %v\n", arr1 == arr2)
	fmt.Printf("arr2 == arr3: %v\n", arr2 == arr3)
}

func comparisonSlice() {
	var nilSlice []int    //пустой срез
	emptySlice := []int{} //нулевой срез (незаполненный)

	fmt.Printf("nilSlice == nil: %v\n", nilSlice == nil)
	fmt.Printf("emptySlice == nil: %v\n", emptySlice == nil)
}

func lenCapMassiveSlice() {
	arr := [5]int{1, 2, 3, 4, 5} // Массив: len и cap всегда равны размеру массива

	fmt.Printf("Массив %v\n", arr)
	fmt.Printf("	len(arr): %d\n", len(arr))
	fmt.Printf("	cap(arr): %d\n", cap(arr))
	fmt.Println()

	slice := make([]int, 3, 5) //создаем срез длиной 3 и емкостью 5
	slice[0] = 10              //присваиваем значения в срез
	slice[1] = 20
	slice[2] = 30

	fmt.Printf("Срез: %v\n", slice)
	fmt.Printf("Доступные элементы (len(slice)): %d\n", len(slice))
	fmt.Printf("Доступная емкость (cap(slice)): %d\n", cap(slice))
	fmt.Println()

	var nilSlice []int          //nil срез, который не имеет значений и не имеет длины
	emptySlice := []int{}       //пустой срез
	zeroSlice := make([]int, 0) //нулевой срез

	fmt.Println("Сравнение nil, пустого и zero срезов")
	fmt.Printf("nil: %d,	len = %d, 	cap = %d, 	==nil: %v\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("empty: %d,	len = %d, 	cap = %d, 	==nil: %v\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("zero: %d,	len = %d, 	cap = %d, 	==nil: %v\n", zeroSlice, len(zeroSlice), cap(zeroSlice), zeroSlice == nil)

}

func creatureSliceMassive() {
	//создание срезов и массивов, способы
	colors := []string{"красный", "зеленый", "синий"} //создание среза с помощью литерала
	fmt.Printf("Литерал: %v\n", colors)
	fmt.Printf(" len: %d, cap: %d\n\n", len(colors), cap(colors))

	scores := make([]int, 5, 10) //создание среза с помощью функции make
	for i := 0; i < len(scores); i++ {
		scores[i] = (i + 1) * 10

	}
	fmt.Printf("make с длиной: %v\n", scores)
	fmt.Printf(" len: %d, cap: %d\n\n", len(scores), cap(scores))

	buffer := make([]int, 10, 64) //создание среза с помощью функции make с длиной и емкостью
	fmt.Printf("make с длиной и емкостью: %v\n", buffer)
	fmt.Printf(" len: %d, cap: %d\n\n", len(buffer), cap(buffer))

	result := make([]int, 0, 10) //пустой, но с резервом для накопления данных
	fmt.Printf("make пустой срез с резервом: %v\n", result)
	fmt.Printf(" len: %d, cap: %d\n\n", len(result), cap(result))

	for i := 0; i < 5; i++ {
		result = append(result, i*i) //добавление элементов в срез без реалокации
	}
	fmt.Printf("После добавления: %v\n", result)
	fmt.Printf("len: %d, cap: %d\n\n", len(result), cap(result))
}

func srez() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //создаем исходный срез
	fmt.Printf("Исходный срез: %v\n", data)
	fmt.Printf("len(data): %v, cap(data): %v\n\n", len(data), cap(data))

	//различные формы срезов
	middle := data[3:6] //срез с 3 по 6 индекс (3 значения)
	fmt.Printf("Срез с 3 по 6 индекс: %v\n", middle)
	fmt.Printf("len(middle): %v, cap(middle): %v\n\n", len(middle), cap(middle))

	first := data[:4] //срез с начала до 4 индекса (4 значения)
	fmt.Printf("Срез с первыми четырмя значениями: %v\n", first)
	fmt.Printf("len(first): %v, cap(first): %v\n\n", len(first), cap(first))

	last := data[5:] //срез с 5 индекса до конца (5 значений)
	fmt.Printf("Срез с 5 индекса до конца: %v\n", last)
	fmt.Printf("len(last): %v, cap(last): %v\n\n", len(last), cap(last))

	//создание нового среза на основе существующего
	newSlice := append(data[:5], 100, 200, 300) //создаем новый срез на основе первых 5 элементов исходного среза и добавляем новые элементы
	fmt.Printf("Новый срез: %v\n", newSlice)
	fmt.Printf("len(newSlice): %v, cap(newSlice): %v\n\n", len(newSlice), cap(newSlice))

	//очистка среза
	workSlice := []int{100, 200, 300, 400, 500} //создаем срез для очистки
	fmt.Printf("Исходный срез для очистки: %v\n\n", workSlice)

	workSlice = workSlice[:0] //очищаем срез, оставляя емкость прежней
	fmt.Printf("Очищенный срез: %v\n", workSlice)
	fmt.Printf("len(workSlice): %v, cap(workSlice): %v\n\n", len(workSlice), cap(workSlice))

	//можем использовать емкость очищенного среза для добавления новых элементов без выделения новой памяти
	workSlice = append(workSlice, 1, 2, 3) //добавляем новые элементы в очищенный срез
	fmt.Printf("Срез после добавления новых элементов: %v\n", workSlice)
	fmt.Printf("len(workSlice): %v, cap(workSlice): %v\n\n", len(workSlice), cap(workSlice))

	/* // Частая ошибка в циклах
			for i := 0; i <= len(s); i++ {  // <= вместо < (ошибка!)
	   		 fmt.Println(s[i])  // Ошибка на последней итерации
		}
	*/

}

func equalSlices(a, b []int) bool { //функция сравнения срезов
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func summElements() {
	s := []int{10, 20, 30, 40, 50}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	fmt.Printf("Срез: %v\n", s)
	fmt.Printf("Сумма элементов: %d\n", sum)
	fmt.Printf("Длина среза: %d\nЕмкость среза: %d\n", len(s), cap(s))
}

func zeroingSlice() {
	slice := []int{1, 2, 3}
	fmt.Printf("Срез: %v\n", slice)
	fmt.Printf("Длина среза: %d\nЕмкость среза: %d\n\n", len(slice), cap(slice))

	// Обнуляем срез
	slice = slice[:0] //обнуляем срез, оставляя емкость прежней
	fmt.Printf("Обнуленный срез: %v\n", slice)
	fmt.Printf("Длина среза: %d\nЕмкость среза: %d\n\n", len(slice), cap(slice))

	//добавляем элементы в обнуленный срез
	slice = append(slice, 4, 5, 6)
	fmt.Printf("Срез после добавления элементов: %v\n", slice)
	fmt.Printf("Длина среза: %d\nЕмкость среза: %d\n\n", len(slice), cap(slice))
}

func main() {
	massive()
	fmt.Println("")

	slice()
	fmt.Println("")

	zero()
	fmt.Println("")

	comparisonMassive()
	fmt.Println("")

	comparisonSlice()
	fmt.Println("")

	lenCapMassiveSlice()
	fmt.Println()

	creatureSliceMassive()
	fmt.Println()

	srez()

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(equalSlices(a, b))
	fmt.Println()

	summElements()
	fmt.Println()

	zeroingSlice()
	fmt.Println()
}
