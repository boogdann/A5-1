package rank

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRank(t *testing.T) {
	type args struct {
		bits []byte
		m    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				bits: strToBits(readTestData("../../../data.txt")),
				m:    32,
				q:    32,
			},
			want: 0.5320686208924519,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := New(tt.args.bits, tt.args.m, tt.args.q)
			if got, _ := test.Run(); got != tt.want {
				t.Errorf("frequency.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func strToBits(s string) []byte {
	data := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		data = append(data, s[i]-'0')
	}
	return data
}

func readTestData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", data)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")

	fmt.Printf("str: %d\n", len(str))

	return str[:100000]
}
