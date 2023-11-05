package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Artiii4/calculations/for3Lab/pkg"
	log "github.com/sirupsen/logrus"
)

func main() {
	useFlag := flag.Bool("flag", false, "flag have to be false or true")
	flag.Parse()
	if len(os.Args[1:]) < 1 {
		fmt.Println("Необходимо указать значение числа в аргументах командной строки.")
		return
	}
	fir, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Неверное значение числа в аргументе командной строки.")
		return
	}
	n := int64(fir)
	log.Info(pkg.Calculate(n, *useFlag))
	pkg.Calculate(n, false)
}
