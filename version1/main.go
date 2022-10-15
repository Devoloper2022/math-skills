package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	check := os.Args
	if len(check) == 2 {
		from := check[1]

		// get from file data
		content, err := ioutil.ReadFile(from)
		if err != nil {
			log.Fatal(err)
		}

		// split data
		array := strings.Fields(string(content))
		var numArray []float64
		var sum float64
		sum = 0.0
		n := len(array)

		// convert sting to float and calculation SUM
		for i := 0; i < len(array); i++ {
			num, numerr := strconv.ParseFloat(array[i], 64)
			if numerr == nil {
				sum += num
				numArray = append(numArray, float64(num))
			}
		}

		// calculate Average
		avg := sum / float64(n)

		// calculate Variance

		var v float64
		v = 0.0
		for i := 0; i < len(numArray); i++ {
			v += math.Pow(numArray[i]-avg, 2)
		}

		v = v / float64(n)
		// calculate Standard Deviation
		sd := math.Sqrt(v)

		// calculate Median
		var mid float64
		s := n / 2
		if n%2 == 0 {
			mid = (numArray[s-1] + numArray[s]) / 2
		} else {
			mid = numArray[s]
		}

		// Result Print
		ansnum := int(math.Round(avg))
		ans := strconv.Itoa(ansnum)
		fmt.Printf(" Average: %s\n ", ans)

		ansnum = int(math.Round(mid))
		ans = strconv.Itoa(ansnum)
		fmt.Printf("Median: %s\n ", ans)

		ansnum = int(math.Round(v))
		ans = strconv.Itoa(ansnum)
		fmt.Printf("Variance: %s\n ", ans)

		ansnum = int(math.Round(sd))
		ans = strconv.Itoa(ansnum)
		fmt.Printf("Standard Deviation: %s\n ", ans)

	} else {
		fmt.Println("Please enter arguments: asource.txt")
	}
}
