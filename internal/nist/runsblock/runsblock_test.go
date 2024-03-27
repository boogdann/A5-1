package runsblock

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRuns(t *testing.T) {
	type args struct {
		bits []byte
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				bits: strToBits("11001100000101010110110001001100111000000000001001001101010100010001001111010110100000001101011111001100111001101101100010110010"),
			},
			want: 0.18060966766949693,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := New(tt.args.bits)
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
