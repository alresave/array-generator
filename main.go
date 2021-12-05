package main
import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)
func main() {
	aLen := flag.Int("len", 100, "Length of the array to generate")
	aMin := flag.Int("min", 0, "Initial or starting number. For random numbers specifies the minimum. For linear arrays it specifies the starting number.")
	aMax := flag.Int("max", 100, "Max element for random arrays. Will be ignored for linear arrays because the end will be determined by the length.")
	aRand := flag.Bool("rnd", false, "Specifies if the array should be made of random numners or linear")
	aSep := flag.String("sep", ",", "Separator for the array elements")
	aDel := flag.String("enc", "c", "Array enclosing character. Valid values are 'c' for curly braces, 'b' for brackets or 'p' for parenthesis.")
	flag.Parse()
	var arr []string
	str := "{"
	end := "}"
	if *aDel == "b" {
		str = "["
		end = "]"
	} else if *aDel == "p"{
		str = "("
		end = ")"
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *aLen; i++ {
		if *aRand {
			arr = append(arr, strconv.FormatInt(int64(*aMin + rand.Intn(*aMax-*aMin+1)), 10))
		}else {
			arr = append(arr, strconv.Itoa(i + *aMin))
		}
	}
	fmt.Println(str + strings.Join(arr, *aSep) + end)
}