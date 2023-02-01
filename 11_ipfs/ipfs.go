package _1_ipfs

import (
	"fmt"
	"github.com/alanshaw/go-carbites"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	bigCar, _ := os.Open("big.car")
	targetSize := 1024 * 1024 // 1MiB chunks
	strategy := carbites.Treewalk
	spltr, _ := carbites.Split(bigCar, targetSize, strategy)

	var i int
	for {
		car, err := spltr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		b, _ := ioutil.ReadAll(car)
		ioutil.WriteFile(fmt.Sprintf("chunk-%d.car", i), b, 0644)
		i++
	}
}
