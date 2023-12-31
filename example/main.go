package main

import (
	"fmt"

	tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)

func main() {
	sentences := []string{"Hello world, This is ā text!!! 你好呀 "}

	// FullTokenizer
	seqLen, lower := 128, true
	voc, err := tokenizer.FromFile("./tmp/vocab.txt")
	if err != nil {
		panic(err)
	}
	tkz := tokenizer.NewFullTokenizer(voc, seqLen, lower)
	for _, sentence := range sentences {
		fmt.Println(tkz.Tokenize(sentence).TokenIDs)
		fmt.Println(tkz.Tokenize(sentence).MaskIDs)
		fmt.Println(tkz.Tokenize(sentence).TypeIDs)
	}

	// Basic Tokenizer
	tkz1 := tokenizer.NewBasicTokenizer(lower)
	for _, sentence := range sentences {
		fmt.Println(tkz1.Tokenize(sentence))
	}

	// Wordpiece Tokenizer
	tkz2 := tokenizer.NewWordpieceTokenizer(voc)
	for _, sentence := range sentences {
		fmt.Println(tkz2.Tokenize(sentence))
	}
}
