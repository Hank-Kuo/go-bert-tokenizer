package go_bert_tokenizer

import (
	"strings"
)

const (
	ClassToken        = "[CLS]"
	SeparatorToken    = "[SEP]"
	SequenceSeparator = " ||| "
)

type Encode struct {
	Text     string
	Tokens   []string
	TokenIDs []int32
	MaskIDs  []int32
	TypeIDs  []int32
}

type FullTokenizer struct {
	Basic     *BasicTokenizer
	Wordpiece *WordpieceTokenizer
	SeqLen    int
}

func NewFullTokenizer(voc *Vocab, seqLen int, lower bool) *FullTokenizer {
	tkz := &FullTokenizer{
		Basic:     NewBasicTokenizer(lower),
		Wordpiece: NewWordpieceTokenizer(voc),
		SeqLen:    seqLen,
	}
	return tkz
}

func (tkz *FullTokenizer) Tokenize(text string) *Encode {
	encode := &Encode{
		Text:     text,
		Tokens:   make([]string, tkz.SeqLen),
		TokenIDs: make([]int32, tkz.SeqLen),
		MaskIDs:  make([]int32, tkz.SeqLen),
		TypeIDs:  make([]int32, tkz.SeqLen),
	}
	parts := strings.Split(text, SequenceSeparator)
	seqs := make([][]string, len(parts))

	for i, part := range parts {
		seqs[i] = tkz.tokenize(part)
	}

	truncate(seqs, tkz.SeqLen-len(seqs)-1)

	voc := tkz.Wordpiece.Vocab

	var s int
	encode.Tokens[s] = ClassToken
	encode.TokenIDs[s] = voc.GetID(ClassToken)
	encode.TypeIDs[s] = 0
	encode.MaskIDs[s] = 1
	s++
	for sid, seq := range seqs {
		for _, tok := range seq {
			encode.Tokens[s] = tok
			encode.TokenIDs[s] = voc.GetID(tok)
			encode.TypeIDs[s] = int32(sid)
			encode.MaskIDs[s] = 1
			s++

		}
		encode.Tokens[s] = SeparatorToken
		encode.TokenIDs[s] = voc.GetID(SeparatorToken)
		encode.TypeIDs[s] = int32(sid)
		encode.MaskIDs[s] = 1
		s++
	}
	return encode
}

func (tkz *FullTokenizer) tokenize(text string) []string {
	var toks []string
	for _, tok := range tkz.Basic.Tokenize(text) {
		toks = append(toks, tkz.Wordpiece.Tokenize(tok)...)
	}
	return toks
}

func truncate(seqs [][]string, maxlen int) {
	var seqlen int
	for i := range seqs {
		seqlen += len(seqs[i])

	}
	for slen := seqlen; slen > maxlen; slen-- {
		var mi, mv int
		for i := len(seqs) - 1; i >= 0; i-- {
			seq := seqs[i]
			if len(seq) > mv {
				mi = i
				mv = len(seq)
			}
		}
		if mv <= 0 {
			return
		}

		seqs[mi] = seqs[mi][:len(seqs[mi])-1]
	}
}
