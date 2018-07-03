package sort

import (
	"fmt"
)

func main() {
	var testing = []int{3, 1, 9}
	// fmt.Println(BubbleSort(testing))
	quickSort(testing, 0, len(testing)-1)
}

/**
 * 冒泡排序
 */
func BubbleSort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return values
}

/**
 * 快速排序 s
 * @param  {[type]} values []int         [description]
 * @param  {[type]} left   [description]
 * @param  {[type]} right  int           [description]
 * @return {[type]}        [description]
 */
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
	fmt.Println(values)
}
