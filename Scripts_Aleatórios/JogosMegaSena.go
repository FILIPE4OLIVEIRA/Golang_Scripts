package main

import (
	"math/rand"
	system "os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	AllNumbs := Numeros_Mega(650)
	Grava_Resultado(AllNumbs)
	//fmt.Println(AllNumbs)
	//fmt.Println(len(AllNumbs))
}

func Numeros_Mega(numero_de_jogos int) []string {

	var lista_numeros []int
	var lista_jogos []string
	var lista_string []string
	var new_lista_string string

	for i := 0; i < numero_de_jogos; i++ {
		lista_numeros = nil
		for len(lista_numeros) < 6 {
			numero_sorteado := AleatNumber() //rand.Intn(60)
			if exist(lista_numeros, numero_sorteado) || numero_sorteado == 0 {
				lista_numeros = remove(lista_numeros, numero_sorteado)
			} else {
				lista_numeros = append(lista_numeros, numero_sorteado)
			}
		}

		sort.Ints(lista_numeros)
		lista_string = toString(lista_numeros)
		new_lista_string = strings.Join(lista_string, " ")
		lista_jogos = append(lista_jogos, "[")
		lista_jogos = append(lista_jogos, new_lista_string)
		lista_jogos = append(lista_jogos, "] ")
	}
	return (lista_jogos)
}

func exist(elemento []int, lista int) bool {
	for _, v := range elemento {
		if v == lista {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	position := search(slice, s)
	return append(slice[:position], slice[position+1:]...)
}

func search(list []int, j int) int {
	var aux int
	for i := 0; i < len(list); i++ {
		if list[i] == j {
			aux = i
			return aux
		}
	}
	return aux
}

func toString(lista []int) (lista_string []string) {
	zero := "0"
	for j := 0; j <= len(lista)-1; j++ {
		if lista[j] < 10 {
			lista_string = append(lista_string, zero+strconv.Itoa(lista[j]))
		} else {
			lista_string = append(lista_string, strconv.Itoa(lista[j]))
		}
	}
	return
}

func AleatNumber() (aleatnumber int) {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 60
	numb := rand.Intn(max-min+1) + min

	return numb
}

func Grava_Resultado(result []string) {
	caminho, _ := system.Getwd()
	arquivo, _ := system.Create(caminho + "\\NumerosMega.txt")
	for i := 0; i <= len(result)-1; i++ {
		if result[i] == "] " {
			arquivo.WriteString(strings.Join(result[i-2:i+1], "") + "\n")
		}
	}
	arquivo.Close()
}
