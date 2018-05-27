package util

// ArrayStringShuffle ... string配列をシャッフルする
func ArrayStringShuffle(arr []string) []string {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		j := IntRand(0, i+1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// ArrayIntShuffle ... int配列をシャッフルする
func ArrayIntShuffle(arr []int) []int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		j := IntRand(0, i+1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// ArrayStringInsert ... string配列の任意の場所に挿入する
func ArrayStringInsert(arr []string, v string, i int) []string {
	return append(arr[:i], append([]string{v}, arr[i:]...)...)
}

// ArrayIntInsert ... int配列の任意の場所に挿入する
func ArrayIntInsert(arr []int, v int, i int) []int {
	return append(arr[:i], append([]int{v}, arr[i:]...)...)
}

// ArrayStringDelete ... string配列の任意の値を削除する
func ArrayStringDelete(arr []string, i int) []string {
	return append(arr[:i], arr[i+1:]...)
}

// ArrayIntDelete ... int配列の任意の値を削除する
func ArrayIntDelete(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

// ArrayStringShift ... string配列の先頭を切り取る
func ArrayStringShift(arr []string) (string, []string) {
	return arr[0], arr[1:]
}

// ArrayIntShift ... int配列の先頭を切り取る
func ArrayIntShift(arr []int) (int, []int) {
	return arr[0], arr[1:]
}

// ArrayStringBack ... string配列の後尾を切り取る
func ArrayStringBack(arr []string) (string, []string) {
	return arr[len(arr)-1], arr[:len(arr)-1]
}

// ArrayIntBack ... int配列の後尾を切り取る
func ArrayIntBack(arr []int) (int, []int) {
	return arr[len(arr)-1], arr[:len(arr)-1]
}

// ArrayStringFilter ... string配列をフィルタする
func ArrayStringFilter(arr []string, fn func(string) bool) []string {
	ret := []string{}
	for _, v := range arr {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// ArrayIntFilter ... int配列をフィルタする
func ArrayIntFilter(arr []int, fn func(int) bool) []int {
	ret := []int{}
	for _, v := range arr {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
