package freqblock

import "testing"

func TestBlockFrequency(t *testing.T) {
	type args struct {
		bits      []byte
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				bits:      strToBits("1011010101"),
				blockSize: 3,
			},
			want: 0.801251989289459,
		},
		{
			name: "Test 2",
			args: args{
				bits:      strToBits("1100100100001111110110101010001000100001011010001100001000110100110001001100011001100010100010111000"),
				blockSize: 10,
			},
			want: 0.7064384496412819,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := New(tt.args.bits, tt.args.blockSize)
			if got := test.Run(); got != tt.want {
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
