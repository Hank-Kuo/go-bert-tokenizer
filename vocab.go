package go_bert_tokenizer

import (
	"bufio"
	"os"
)

type ID int32

func (id ID) Int32() int32 {
	return int32(id)
}

type Vocab struct {
	tokens map[string]ID
}

func FromFile(path string) (*Vocab, error) {
	f, err := os.Open(path)
	if err != nil {
		return &Vocab{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	voc := &Vocab{tokens: map[string]ID{}}
	for scanner.Scan() {
		voc.Add(scanner.Text())
	}
	return voc, nil
}

func New(tokens []string) *Vocab {
	v := make(map[string]ID, len(tokens))
	for i, t := range tokens {
		v[t] = ID(i)
	}
	return &Vocab{tokens: v}
}

func (v *Vocab) Add(token string) {
	v.tokens[token] = ID(v.Size())
}

func (v *Vocab) GetID(token string) ID {
	id, ok := v.tokens[token]
	if !ok {
		return ID(-1)
	}
	return ID(id)
}

func (v *Vocab) Size() int {
	return len(v.tokens)
}

func (v *Vocab) GetToken() map[string]ID {
	return v.tokens
}
