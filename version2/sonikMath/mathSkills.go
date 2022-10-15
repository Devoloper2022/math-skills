package sonikMath

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func MathSkills() {
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
		// preparation
		sort.Strings(array)
		n := len(array)
		s := n / 2 // for Median
		sum, dataErr := Sum(array)
		if dataErr != nil {
			log.Fatal(dataErr)
			return
		}
		// calculation
		avg := getAvg(sum, n)
		v := getVariance(array, avg, n)
		sd := math.Sqrt(v)
		mid := getMedian(array[s], array[s-1], n)
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
