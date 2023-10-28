package main

import (
	"fmt"
	tokenizer "go-bert-tokenizer"
)

func main() {
	sentences := []string{"Hello world Hello This is ā text!!! 你好呀 "}

	// FullTokenizer
	seqLen, lower := 128, true
	voc, _ := tokenizer.FromFile("../tmp/vocab.txt")
	tkz := tokenizer.NewFullTokenizer(voc, seqLen, lower)
	for _, sentence := range sentences {
		fmt.Println(tkz.Tokenize(sentence).TokenIDs)
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
