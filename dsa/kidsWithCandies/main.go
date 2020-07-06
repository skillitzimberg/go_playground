package main

import "fmt"

func main() {
	kidsWithCandies2([]int{12, 1, 12}, 10)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := 0
	result := make([]bool, len(candies))
	j := len(candies) - 1

	for i := 0; i < len(candies); i++ {
		if i != j && candies[i] > candies[j] {
			greatest = candies[i]
			fmt.Println("i", i, candies[i])
		}

		if i != j && candies[i] < candies[j] {
			greatest = candies[j]
			fmt.Println("j", j, candies[j])
		}

		if i == j && candies[i] > greatest {
			greatest = candies[i]
			break
		}
		j--
	}

	fmt.Println(greatest)

	for i, v := range candies {
		if v+extraCandies >= greatest {
			result[i] = true
		}
	}
	fmt.Println(result)
	return result
}

func kidsWithCandies2(candies []int, extraCandies int) []bool {
	greatest := 0
	res := make([]bool, len(candies))
	for _, v := range candies {
		if v > greatest {
			greatest = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= greatest {
			res[i] = true
		}
	}
	fmt.Println(res)
	return res
}
