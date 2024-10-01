package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_1",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "test_2",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "test_3",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits)
			assert.Equal(t, got, tt.want)
		})
	}
}
