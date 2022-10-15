package sonikMath

import (
	"fmt"
	"math"
	"strconv"
)

func getVariance(arr []string, avg float64, n int) float64 {
	v := 0.0
	for i := 0; i < len(arr); i++ {
		num, err := strconv.ParseFloat(arr[i], 64)
		if err != nil {
			fmt.Println(arr[i] + "is mot number")
			return 0
		}
		v += math.Pow(num-avg, 2)
	}
	return v / float64(n)
}

func getMedian(a, b string, n int) float64 {
	c, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(a + " is not number")
		return 0
	}
	if n%2 != 0 {
		return float64(c)
	}

	d, err2 := strconv.ParseFloat(b, 64)

	if err2 != nil {
		fmt.Println(b + " is not number")
		return 0
	}
	return (c + d) / float64(2)
}

func getAvg(sum, n int) float64 {
	return math.Round(float64(sum) / float64(n))
}

func getSD(v float64) float64 {
	return math.Sqrt(v)
}

func Sum(arr []string) (int, error) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}
