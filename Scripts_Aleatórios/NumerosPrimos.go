package main

// Cria uma Lista de NÃºmeros Primos

// FunÃ§Ã£o Privada Minuscula
// FunÃ§Ã£o Publica Maiuscula

import (
	system "os"
	"strconv"
	"strings"
)

func main() {
	lista_primos := Numeros_Primos(900)
	Grava_Resultado(toString(lista_primos))
	//fmt.Println(lista_primos)
	//fmt.Println(len(lista_primos))
}

func Numeros_Primos(limite int) (lista_primos []int) {

	var prime_numbers []int // primos

	prime_numbers = []int{2}

	for numero := 2; numero < limite+1; numero++ {
		spread := numero % 2
		if spread != 0 {
			prime_numbers = append(prime_numbers, numero)
		}
	}

	for numb1 := 2; numb1 < len(prime_numbers)+1; numb1++ {
		for numb2 := 3; numb2 < len(prime_numbers)+1; numb2++ {
			spread := numb1 * numb2
			if existe(prime_numbers, spread) {
				prime_numbers = remove(prime_numbers, spread)
			}
		}
	}
	return (prime_numbers)
}

func existe(elemento []int, lista int) bool {
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

func Grava_Resultado(result []string) {
	caminho, _ := system.Getwd()
	arquivo, _ := system.Create(caminho + "\\NumerosMega.txt")
	slice := lineBound(len(result))
	for i := 2; i < len(result)+1; i++ {
		if i%slice == 0 {
			arquivo.WriteString(strings.Join(result[i-slice:i], " ") + "\n")
		}
	}
	arquivo.Close()
}

func lineBound(lenght int) (divisor int) {
	for i := 5; i < lenght; i++ {
		divisor := lenght % i
		if divisor == 0 {
			return i
		}
	}
	return
}
