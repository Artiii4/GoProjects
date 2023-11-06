package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Artiii4/contains/pkg"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type List struct {
	Name string `yaml:"Name"`
}

func makeList(data []byte, list *[]List) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	for {
		var bo List
		if err := decoder.Decode(&bo); err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return
			}
			break
		}
		*list = append(*list, bo)
	}
}

func main() {
	fmt.Print("Give the yml file you want to check\n")
	reader := bufio.NewReader(os.Stdin)
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(strings.TrimSpace(fileName))
	if err != nil {
		log.Fatal(err)
		return
	}
	var list []List
	makeList(data, &list)
	fmt.Print("Give the txt file you want to check\n")
	read := bufio.NewReader(os.Stdin)
	name, err := read.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	flag := false
	for _, order := range list {
		if order.Name == strings.TrimSpace(name) {
			flag = true
			fmt.Println("Found file!\nWhich substring do you want?")
			reader := bufio.NewReader(os.Stdin)
			substring, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			result, err := pkg.Contains(order.Name, strings.TrimSpace(substring))
			if err != nil {
				log.Fatal(err)
				return
			} else if result {
				log.Info("We have it!")
			} else {
				log.Info("We dont have this substring!")
			}
		}
	}
	if !flag {
		log.Info("We dont have such a file! Try another one")
	}
}
