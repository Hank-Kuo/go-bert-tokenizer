package go_bert_tokenizer

import (
	"testing"
)

func TestIsPunctuation(t *testing.T) {

	type fields struct {
		Char string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test case: 1",
			fields{
				Char: "r",
			},
			false,
		},
		{
			"Test case: 2",
			fields{
				Char: "郭",
			},
			false,
		},
		{
			"Test case: 3",
			fields{
				Char: ",",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			want := isPunctuation([]rune(tt.fields.Char)[0])

			if want != tt.want {
				t.Errorf("isPunctuation() result should be %t", tt.want)
			}
		})
	}
}

func TestIsControl(t *testing.T) {

	type fields struct {
		Char string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test case: 1",
			fields{
				Char: "r",
			},
			false,
		},
		{
			"Test case: 2",
			fields{
				Char: "郭",
			},
			false,
		},
		{
			"Test case: 3",
			fields{
				Char: ",",
			},
			false,
		},
		{
			"Test case: 4",
			fields{
				Char: "\u0000",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			want := isControl([]rune(tt.fields.Char)[0])

			if want != tt.want {
				t.Errorf("isControl() result should be %t", tt.want)
			}
		})
	}
}

func TestIsWhitespace(t *testing.T) {

	type fields struct {
		Char string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test case: 1",
			fields{
				Char: "r",
			},
			false,
		},
		{
			"Test case: 2",
			fields{
				Char: " ",
			},
			true,
		},
		{
			"Test case: 3",
			fields{
				Char: "\r",
			},
			true,
		},
		{
			"Test case: 4",
			fields{
				Char: "\n",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			want := isWhitespace([]rune(tt.fields.Char)[0])

			if want != tt.want {
				t.Errorf("isWhitespace() result should be %t", tt.want)
			}
		})
	}
}

func TestIsChinese(t *testing.T) {

	type fields struct {
		Char string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test case: 1",
			fields{
				Char: "r",
			},
			false,
		},
		{
			"Test case: 2",
			fields{
				Char: "郭",
			},
			true,
		},
		{
			"Test case: 3",
			fields{
				Char: "你",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			want := isChinese([]rune(tt.fields.Char)[0])

			if want != tt.want {
				t.Errorf("isChinese() result should be %t", tt.want)
			}
		})
	}
}
