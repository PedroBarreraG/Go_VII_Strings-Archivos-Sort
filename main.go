package main

import (
	"fmt"
	"os"
	"sort"
)

type Lista struct {
	cadenas string
}

type BycadenasA []Lista

func (a BycadenasA) Len() int           { return len(a) }
func (a BycadenasA) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BycadenasA) Less(i, j int) bool { return a[i].cadenas < a[j].cadenas }

type BycadenasD []Lista

func (a BycadenasD) Len() int           { return len(a) }
func (a BycadenasD) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BycadenasD) Less(i, j int) bool { return a[i].cadenas > a[j].cadenas }

func main() {
	var lista []Lista

	op := 1
	for op != 2 {
		fmt.Print("1.Agregar String\n2.Generar Archivos\nOpcion: ")
		fmt.Scan(&op)
		if op == 1 {
			var nuevoString string
			fmt.Print("String: ")
			fmt.Scan(&nuevoString)
			lista = append(lista, Lista{cadenas: nuevoString})
		}
	}

	listAscendente := lista

	sort.Sort(BycadenasA(listAscendente))
	//fmt.Println(listAscendente)

	fileA, err := os.Create("asecendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileA.Close()

	for i, _ := range listAscendente {
		fileA.WriteString(listAscendente[i].cadenas)
		fileA.WriteString("\n")
	}

	listDescendente := lista

	sort.Sort(BycadenasD(listDescendente))
	//fmt.Println(listDescendente)

	fileD, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileD.Close()

	for i, _ := range listDescendente {
		fileD.WriteString(listDescendente[i].cadenas)
		fileD.WriteString("\n")
	}
}
