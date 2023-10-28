package go_bert_tokenizer

import (
	"bufio"
	"os"
)

type Vocab struct {
	tokens map[string]int32
}

func FromFile(path string) (*Vocab, error) {
	f, err := os.Open(path)
	if err != nil {
		return &Vocab{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	voc := &Vocab{tokens: map[string]int32{}}
	for scanner.Scan() {
		voc.Add(scanner.Text())
	}
	return voc, nil
}

func New(tokens []string) *Vocab {
	v := make(map[string]int32, len(tokens))
	for i, t := range tokens {
		v[t] = int32(i)
	}
	return &Vocab{tokens: v}
}

func (v *Vocab) Add(token string) {
	v.tokens[token] = int32(v.Size())
}

func (v *Vocab) GetID(token string) int32 {
	id, ok := v.tokens[token]
	if !ok {
		return int32(-1)
	}
	return int32(id)
}

func (v *Vocab) Size() int {
	return len(v.tokens)
}

func (v *Vocab) GetToken() map[string]int32 {
	return v.tokens
}
