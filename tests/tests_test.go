package tests

import (
	"2/internal/nist/discrete"
	"2/internal/nist/freqblock"
	"2/internal/nist/frequency"
	"2/internal/nist/rank"
	"2/internal/nist/runs"
	"2/internal/nist/runsblock"
	"fmt"
	"os"
	"testing"
)

var (
	file *os.File
)

func init() {
	var err error
	file, err = os.Create("tests.txt")
	if err != nil {
		panic(err)
	}
}

func TestDiscrete(t *testing.T) {
	type args struct {
		bits []byte
		path string
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				path: "cipher/textfile1.method_1.save.txt",
				bits: readTestData("cipher/textfile1.method_1.save.txt"),
			},
			want: 0.17776838723061217,
		},
		{
			name: "Test 2",
			args: args{
				path: "cipher/textfile1.method_2.save.txt",
				bits: readTestData("cipher/textfile1.method_2.save.txt"),
			},
			want: 0.30274624213903545,
		},
		{
			name: "Test 3",
			args: args{
				path: "cipher/img.method_1.save.txt",
				bits: readTestData("cipher/img.method_1.save.txt"),
			},
			want: 0.47029107121273206,
		},
		{
			name: "Test 4",
			args: args{
				path: "cipher/img.method_2.save.txt",
				bits: readTestData("cipher/img.method_2.save.txt"),
			},
			want: 0.8906159401689214,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := discrete.New(tt.args.bits)
			got := test.Run()
			fmt.Fprintf(file, `DiscreteTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)
			if got != tt.want {
				t.Errorf("discrete.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockFrequency(t *testing.T) {
	type args struct {
		path      string
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
				path:      "cipher/textfile1.method_1.save.txt",
				bits:      readTestData("cipher/textfile1.method_1.save.txt"),
				blockSize: 8,
			},
			want: 0.30154407845140996,
		},
		{
			name: "Test 2",
			args: args{
				path:      "cipher/textfile1.method_2.save.txt",
				bits:      readTestData("cipher/textfile1.method_2.save.txt"),
				blockSize: 20,
			},
			want: 0.6666250272126062,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := freqblock.New(tt.args.bits, tt.args.blockSize)
			got := test.Run()
			fmt.Fprintf(file, `BlockFrequencyTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)

			if got != tt.want {
				t.Errorf("blockfrequency.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequency(t *testing.T) {
	type args struct {
		bits []byte
		path string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				path: "cipher/textfile1.method_1.save.txt",
				bits: readTestData("cipher/textfile1.method_1.save.txt"),
			},
			want: 0.665789228018142,
		},
		{
			name: "Test 2",
			args: args{
				path: "cipher/textfile1.method_2.save.txt",
				bits: readTestData("cipher/textfile1.method_2.save.txt"),
			},
			want: 0.7296829044432491,
		},
		{
			name: "Test 3",
			args: args{
				path: "cipher/img.method_1.save.txt",
				bits: readTestData("cipher/img.method_1.save.txt"),
			},
			want: 0.8769319077637521,
		},
		{
			name: "Test 4",
			args: args{
				path: "cipher/img.method_2.save.txt",
				bits: readTestData("cipher/img.method_2.save.txt"),
			},
			want: 0.3929796645778726,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := frequency.New(tt.args.bits)
			got := test.Run()
			fmt.Fprintf(file, `FrequencyTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)

			if got != tt.want {
				t.Errorf("frequency.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank(t *testing.T) {
	type args struct {
		path string
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
			name: "Test 3",
			args: args{
				path: "cipher/img.method_1.save.txt",
				bits: readTestData("cipher/img.method_1.save.txt"),
				m:    32,
				q:    32,
			},
			want: 0.7072488849188965,
		},
		{
			name: "Test 4",
			args: args{
				path: "cipher/img.method_2.save.txt",
				bits: readTestData("cipher/img.method_2.save.txt"),
				m:    32,
				q:    32,
			},
			want: 0.6428812482599868,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := rank.New(tt.args.bits, tt.args.m, tt.args.q)
			got, _ := test.Run()
			fmt.Fprintf(file, `RankTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)

			if got != tt.want {
				t.Errorf("rank.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuns(t *testing.T) {
	type args struct {
		path string
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
				path: "cipher/textfile1.method_1.save.txt",
				bits: readTestData("cipher/textfile1.method_1.save.txt"),
			},
			want: 0.5398897401501386,
		},
		{
			name: "Test 2",
			args: args{
				path: "cipher/textfile1.method_2.save.txt",
				bits: readTestData("cipher/textfile1.method_2.save.txt"),
			},
			want: 0.6005601775487122,
		},
		{
			name: "Test 3",
			args: args{
				path: "cipher/img.method_1.save.txt",
				bits: readTestData("cipher/img.method_1.save.txt"),
			},
			want: 0.7284294508359047,
		},
		{
			name: "Test 4",
			args: args{
				path: "cipher/img.method_2.save.txt",
				bits: readTestData("cipher/img.method_2.save.txt"),
			},
			want: 0.821428590050775,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := runs.New(tt.args.bits)
			got := test.Run()
			fmt.Fprintf(file, `RunsTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)

			if got != tt.want {
				t.Errorf("runs.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunsBlock(t *testing.T) {
	type args struct {
		path string
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
				path: "cipher/textfile1.method_1.save.txt",
				bits: readTestData("cipher/textfile1.method_1.save.txt"),
			},
			want: 0.6151057927425478,
		},
		{
			name: "Test 2",
			args: args{
				path: "cipher/textfile1.method_2.save.txt",
				bits: readTestData("cipher/textfile1.method_2.save.txt"),
			},
			want: 0.2071921779389111,
		},
		{
			name: "Test 3",
			args: args{
				path: "cipher/img.method_1.save.txt",
				bits: readTestData("cipher/img.method_1.save.txt"),
			},
			want: 0.5483105435140134,
		},
		{
			name: "Test 4",
			args: args{
				path: "cipher/img.method_2.save.txt",
				bits: readTestData("cipher/img.method_2.save.txt"),
			},
			want: 0.23721878182220824,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := runsblock.New(tt.args.bits)
			got, _ := test.Run()
			fmt.Fprintf(file, `RunsBlockTest:
	Name:     %s
	FilePath: %s
	Val:      %f

`, tt.name, tt.args.path, got)

			if got != tt.want {
				t.Errorf("runsblock.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readTestData(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	newData := make([]byte, 0, len(data)*8)
	for _, num := range data {
		for i := 7; i >= 0; i-- {
			newData = append(newData, (num>>i)&1)
		}
	}

	return newData
}
