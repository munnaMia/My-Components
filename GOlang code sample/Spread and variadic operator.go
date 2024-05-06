
func main() {
	values := []int{2,23433,34,3234,3,2234343,5,3223434,234,234,23423456,2323423,2344,23456}
	fmt.Println(sums(values...)) // Spread operator here
}

// Variadic function -> which can take multiple values as parameter and use then as slices
func sums(nums ...int) int{
	sums:=0
	for _,v:=range nums {
		sums+=v
	}
	return sums
}
