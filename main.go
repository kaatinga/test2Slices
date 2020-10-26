package main

import "fmt"

//Дано: Два отсортированных целочисленных массива
//Задача: Объединить два массива в один отсортированный массив


func main(){
	//Input:
	var arr1 = []int{1, 3, 4, 5}
	var arr2 = []int{2, 4, 6, 8}

	//Output:
	//var output = []int{1, 2, 3, 4, 4, 5, 6, 8}

	// размер нового слайса
	var outputSize = len(arr1)+len(arr2)

	// создаём заранее слайс нужного размера
	var output = make([]int, outputSize)

	// переменные цикла
	var arr1i, arr2i, outputI int

	var arr1finished, arr2finished bool

	for {

		if arr1i == len(arr1) {
			arr1finished = true
		}

		if arr2i == len(arr2) {
			arr2finished = true
		}

		fmt.Println("=== текущий указатель output", outputI)
		fmt.Println("arr1finished", arr1finished)
		fmt.Println("arr2finished", arr2finished)

		// пока оба слайса не исчерпаны, проверяем какое значение больше и его пишем в текущую ячейку output
		if !arr1finished && !arr2finished {
			if arr1[arr1i] < arr2[arr2i] {
				output[outputI] = arr1[arr1i]
				arr1i++
				fmt.Println("берем arr1i")
			} else {
				output[outputI] = arr2[arr2i]
				arr2i++
				fmt.Println("берем arr2i")
			}
		}

		if !arr1finished && arr2finished {
			output[outputI] = arr1[arr1i]
			fmt.Println("arr2i кончился")
			arr1i++
		}

		if arr1finished && !arr2finished {
			output[outputI] = arr2[arr2i]
			fmt.Println("arr1i кончился")
			arr2i++
		}

		fmt.Println("arr1i=", arr1i, "arr2i=", arr2i)

		// меняем итератор output
		outputI++

		fmt.Println(output)

		if outputSize == outputI {
			break
		}
	}
}

