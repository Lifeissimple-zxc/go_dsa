package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abc and pqr",
			args: args{
				word1: "abc",
				word2: "pqr",
			},
			want: "apbqcr",
		},
		{
			name: "ab and pqrs",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "abcd and pq",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MergeAlternately(tt.args.word1, tt.args.word2))
		})
	}
}

func TestGCDOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ABABABAB and ABAB",
			args: args{
				str1: "ABABABAB",
				str2: "ABAB",
			},
			want: "ABAB",
		},
		{
			name: "ABCABC and ABC",
			args: args{
				str1: "ABCABC",
				str2: "ABC",
			},
			want: "ABC",
		},
		{
			name: "ABABAB and ABAB",
			args: args{
				str1: "ABABAB",
				str2: "ABAB",
			},
			want: "AB",
		},
		{
			name: "no divisor",
			args: args{
				str1: "LEET",
				str2: "CODE",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GCDOfStrings(tt.args.str1, tt.args.str2))
		})
	}

}
