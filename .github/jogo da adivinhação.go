package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numeroescolhido int
	for {
		NA := rand.Intn(100) + 1
		t := 0
		fmt.Println("Seja bem-vindo ao jogo da advinhação!!!!!! ")
		fmt.Println("Neste jogo sera gerado um número aleatório e o jogador deve acertar este número")
		fmt.Print("Digite um número entre 1 a 100: ")
		fmt.Scan(&numeroescolhido)

		for {
			if NA > numeroescolhido {
				fmt.Printf("O número é maior que  %d \n ", numeroescolhido)
				fmt.Print("Digite um número entre 1 a 100: ")
				fmt.Scan(&numeroescolhido)
				t++
			} else if NA < numeroescolhido {
				fmt.Printf("O número é menor que %d \n ", numeroescolhido)
				fmt.Print("Digite um número entre 1 a 100: ")
				fmt.Scan(&numeroescolhido)
				t++
			} else {
				fmt.Println("Parabéns você acertou!!")
				break
			}
		}
		TN := []int{}
		TN = append(TN, t)
		fmt.Printf("Número de tentativas  %d \n", t)
		var denovo string
		fmt.Print("Você deseja jogar novamente ? (Sim/Não): ")
		fmt.Scan(&denovo)
		if denovo == "Não" {
			for i := 0; i < len(TN); i++ {
				fmt.Println("você tentou %d  na tentativa %d", TN[i], i+1)
			}
			break
		}
	}
}
