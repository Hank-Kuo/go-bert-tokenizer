# go-bert-tokenizer
Implement bert tokenizer by go. 

Document:https://pkg.go.dev/github.com/Hank-Kuo/go-bert-tokenizer

Follow Bert tokenzier algo, reference: [bert repo](https://github.com/google-research/bert)

Valify Hugging face python code. 
```python
from transformers import BertTokenizer

tokenizer = BertTokenizer.from_pretrained("bert-base-uncased")
sentences = [
    "This is ā text!!!",
    "This sentence is: 你好呀!很高興認識你（我是XXX） ",
    "This sentence is: Σήμερα είναι τα γενέθλιά μου",
    "This sentence is: Bugün benim doğum günüm",
    "This sentence is: 今日は私の誕生日です",
    "Сегодня мой день рождения",
    "안녕하세요 오늘은 제 생일이에요 만나서 반가워요",
    "สวัสดี วันนี้เป็นวันเกิดของฉัน ยินดีที่ได้รู้จัก",
    "In documentis nostris, prima sententia est numerus 1º.",
    "سلام، امروز تولد من است، از آشنایی با شما خوشحالم"
]

for i in sentences:
    encoding = tokenizer.encode_plus(i, max_length=128, padding='max_length', truncation=True)
    print(encoding["input_ids"])
```

## Usage

### FullTokenizer
```go
import (
    tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)
sentence := "This is a test case !!!!"

voc, err := tokenizer.FromFile("./tmp/vocab.txt") // load vocab for vocab file 
if err != nil {
    panic(err)
}
tkz := tokenizer.NewFullTokenizer(voc, 128, true) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding.Text)
fmt.Println(encoding.Tokens)
fmt.Println(encoding.TokenIDs)
fmt.Println(encoding.MaskIDs)
fmt.Println(encoding.TypeIDs)
```


### WordpieceTokenizer
```go
import (
    tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)

sentence := "This is a test case !!!!"

voc, err := tokenizer.FromFile("./tmp/vocab.txt") // load vocab for vocab file 
if err != nil {
    panic(err)
}
tkz := tokenizer.NewWordpieceTokenizer(voc, true) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding)

```

### BasicTokenizer
```go
import (
    tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)
sentence := "This is a test case !!!!"
tkz := tokenizer.NewBasicTokenizer(voc, lower) 
encoding := tkz.Tokenize(sentence)
fmt.Println(encoding)
```




## Performance
Go benchmark test: ```go test -bench .````

```bash
goos: darwin
goarch: arm64
pkg: github.com/Hank-Kuo/go-bert-tokenizer
BenchmarkTokenize-8        58741             19350 ns/op
PASS
ok      github.com/Hank-Kuo/go-bert-tokenizer   1.793s
```