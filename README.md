# go-bert-tokenizer
Implement bert tokenizer by go. 
Follow this Bert tokenzier 

## Usage

### FullTokenizer
```go
import (
    tokenizer "go-bert-tokenizer"
)

voc, _ := tokenizer.FromFile("./vocab.txt") // load vocab for vocab file 

tkz := tokenizer.NewFullTokenizer(voc1, 128, true) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding.Text)
fmt.Println(encoding.Tokens)
fmt.Println(encoding.TokenIDs)
fmt.Println(encoding.MaskIDs)
fmt.Println(encoding.TypeIDs)
```


### WordpieceTokenizer
```go
import "go-bert-tokenizer/tokenizer"

voc, _ := tokenizer.FromFile("./vocab.txt") // load vocab for vocab file 

tkz := tokenizer.NewFullTokenizer(voc1, 128, true) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding.Text)
fmt.Println(encoding.Tokens)
fmt.Println(encoding.TokenIDs)
fmt.Println(encoding.MaskIDs)
fmt.Println(encoding.TypeIDs)
```

### BasicTokenizer
```go
import "go-bert-tokenizer/tokenizer"

voc, _ := tokenizer.FromFile("./vocab.txt") // load vocab for vocab file 
seqLen, lower := 128, true
tkz := tokenizer.NewFullTokenizer(voc1, seqLen, lower) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding.Text)
fmt.Println(encoding.Tokens)
fmt.Println(encoding.TokenIDs)
fmt.Println(encoding.MaskIDs)
fmt.Println(encoding.TypeIDs)
```


