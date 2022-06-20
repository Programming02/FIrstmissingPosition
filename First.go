//Firstmissing position ----> arrayning ichida eng birinchi unutilgan eng kichik son

package main

import "fmt"

func FirstMissingPositive(nums []int) int {
	hashMap := make(map[int]bool)
	firstmissingPositive := 1

	for _, num := range nums {
		hashMap[num] = true
	}
	for {
		_, ok := hashMap[firstmissingPositive]
		if !ok {
			return firstmissingPositive
		}
		firstmissingPositive += 1
	}
}

func main() {
	nums := []int{2, 4, 1, 3, 65, 123, 87}
	fmt.Println(FirstMissingPositive(nums))
}
