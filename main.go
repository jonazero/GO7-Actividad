package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	ss := []string{}
	ascendente := []string{}
	descendente := []string{}
	fmt.Println("¿Desea agregar un string [y/n]?")
	var letra string
	var str string
	fmt.Scan(&letra)
	for letra == "y" {
		fmt.Print("Ingresa el string: ")
		fmt.Scan(&str)
		ss = append(ss, str)
		ascendente = append(ascendente, str)
		descendente = append(descendente, str)
		fmt.Println("Desea agregar un string [y/n]")
		fmt.Scan(&letra)
	}
	sort.Strings(ascendente)
	sort.Sort(sort.Reverse(sort.StringSlice(descendente)))
	writeascending(ascendente)
	writedescending(descendente)

}
func writeascending(ascendente []string) {
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for _, v := range ascendente {
		file.WriteString(v)
		file.WriteString("\n")

	}
}
func writedescending(descendente []string) {
	file, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for _, v := range descendente {
		file.WriteString(v)
		file.WriteString("\n")

	}
}
