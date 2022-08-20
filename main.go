package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)


type game_struct struct {
	amount_try	uint
	word		string
	guess_word	[]rune
	len_word	uint
}


func main () {
	var game game_struct

	game.get_word()
	fmt.Println("numbers of try:")
	fmt.Scan(&game.amount_try)
	game.lets_play()
}

func (game *game_struct) lets_play() {
	var sym string

	for i := 0; i < int(game.amount_try); i++ {

		fmt.Println("amount of try: ", int(game.amount_try) - i)
		fmt.Println(string(game.guess_word))
		
		fmt.Scan(&sym)
		if find_syms(game, rune(sym[0])) == 0 {
			i-- 
		}

		if game.len_word == uint(0) {
			fmt.Println("Congratulations!")
			break
		}
	}
	if game.len_word != uint(0) {
		fmt.Println("lose!")
	}

}

func find_syms(game *game_struct, sym rune) int {
	count := strings.Count(game.word, string(sym))

	if count == 0 {fmt.Println("symbol not found :("); return -1}

	game.len_word -= uint(count)

	for  i := 0; count > 0; i++ {
		if game.word[i] == byte(sym) {
			count--
			game.guess_word[i] = rune(sym)
		}
	}

	return 0
}


func (game *game_struct) get_word() {
	var words []string

	file, err := os.Open("words.txt")

	if err != nil { panic(err) }

	scanner := bufio.NewScanner(file)

	var count_str uint

	for scanner.Scan() {
		words = append(words, scanner.Text())
		count_str++
	}
	file.Close()

	rand.Seed(time.Now().UnixMicro())
	game.word = words[rand.Intn(int(count_str))]
	game.len_word = uint(utf8.RuneCountInString(game.word))
	game.guess_word = make([]rune, int(game.len_word))
	for i := range game.guess_word {
		game.guess_word[i] = '*'
	}
}



