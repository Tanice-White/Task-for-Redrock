package main
import (
	"fmt"
	"sort"
)
func main() {
	//sort()函数的排序作用（默认是从小到大排序）

	var arr7 = []int{2, 20, 56, 32, 1, 435, 6, 67, 9}
	fmt.Println("\narr7的初始顺序", arr7)
	sort.Ints(arr7) //还有sort.strings()   sort.float64s()的类型
	fmt.Println("arr7被sort.Ints()排序后的顺序", arr7)

	//sort()的降序排序方法
	sort.Sort(sort.Reverse(sort.IntSlice(arr7)))
	fmt.Println("arr7的降序排列", arr7)
	//同理还有
	//sort.Sort(sort.Reverse(sort.Float64Slice(数组名)))
	//sort.Sort(sort.Reserve(sort.StringSlice(数组名)))
}

