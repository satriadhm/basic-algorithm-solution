package main

import "fmt"

func tebakan(player rune, nilai int) int {
	var jawaban int
	i := 1
	fmt.Printf("%c - masukkan angka tebakan ke-%d: ", player, i)
	fmt.Scan(&jawaban)
	for i < 3 && jawaban != nilai {
		i++
		fmt.Printf("%c - masukkan angka tebakan ke-%d: ", player, i)
		fmt.Scan(&jawaban)
	}
	return jawaban
}

func tukerpemenang(benar bool, winner, player *rune) {
	if benar {
		var temp rune = *winner
		*winner = *player
		*player = temp
	}
}

func mulaironde(ronde int, winner rune, nilai *int) {
	fmt.Println("Ronde", ronde, ":")
	fmt.Printf("%c - masukkan angka rahasia: ", winner)
	fmt.Scan(nilai)
}

func main() {
	var winner, player rune
	var ronde, nilai, answer int

	winner = 'A'
	player = 'B'
	ronde = 1

	mulaironde(ronde, winner, &nilai)
	for nilai != -101 {
		answer = tebakan(player, nilai)
		tukerpemenang(answer == nilai, &winner, &player)
		fmt.Printf("%c adalah pemenang!\n", winner)
		ronde = ronde + 1
		mulaironde(ronde, winner, &nilai)
	}
	fmt.Println("Permainan Selesai!")
}
