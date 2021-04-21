package ch1

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(strings.Join(os.Args, ""))
}

func echo2() {
	for i, para := range os.Args {
		fmt.Printf("%d, %v \n", i, para)
	}
}

func echo3() {

}
