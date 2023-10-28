package go_bert_tokenizer

import (
	"testing"
)

func TestTokenize(t *testing.T) {

	type fields struct {
		Text string
	}

	tests := []struct {
		name   string
		fields fields
		want   []int32
	}{
		{
			"Test case: 1",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 2",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 3",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 4",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 5",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 6",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 7",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 8",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 9",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
		{
			"Test case: 10",
			fields{
				Text: "Hello world Hello This is ā text!!! 你好呀",
			},
			[]int32{101, 7592, 2088, 7592, 2023, 2003, 1037, 3793, 999, 999, 999, 100, 100, 100, 102},
		},
	}

	voc, _ := FromFile("./tmp/vocab.txt")
	tkz := NewFullTokenizer(voc, 15, true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encode := tkz.Tokenize(tt.fields.Text)
			if len(encode.TokenIDs) != len(tt.want) {
				t.Errorf("Tokenize() length should be %d", len(tt.want))
			}
			for i, v := range tt.want {
				if encode.TokenIDs[i] != v {
					t.Errorf("Tokenize() idx(%d) should be %d", i, v)
				}
			}
		})
	}
}

func BenchmarkTokenize(b *testing.B) {
	b.ResetTimer()
	voc, _ := FromFile("./tmp/vocab.txt")
	tkz := NewFullTokenizer(voc, 128, true)
	text := "Hello world Hello This is ā text!!! 你好呀This is ā text!!! 你好呀This is ā text!!! 你好呀"

	for i := 0; i < b.N; i++ {
		tkz.Tokenize(text)
	}

	b.StopTimer()
}

/*
go test -bench .
go test -race -coverpkg=. .
*/
