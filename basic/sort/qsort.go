package sort

func QuickSort(num []int) {
	quickSort(num, 0, len(num)-1)
}

func quickSort(num []int, l, r int) {
	if l >= r {
		return
	}
	x := num[l]
	i, j := l, r
	for i < j {
		for ; i < j && num[i] < x; i++ {
		}
		for ; i < j && num[j] > x; j-- {
		}
		num[i], num[j] = num[j], num[i]
	}
	quickSort(num, l, j)
	quickSort(num, j+1, r)
}
