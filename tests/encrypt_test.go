package tests

import (
	a51 "2/internal/a51/v2"
	"2/internal/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		method        int
		filename      string
		key           uint64
		pathTmplt     string
		pathExelTmplt string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				method:        a51.Method1,
				filename:      "cipher/textfile1.method_1.save.txt",
				key:           1232312712189,
				pathTmplt:     "plain/%s.method_%d.save",
				pathExelTmplt: "exelplain/%s.method_%d.save",
			},
		},
		{
			name: "Test 2",
			args: args{
				method:        a51.Method2,
				filename:      "cipher/textfile1.method_2.save.txt",
				key:           984348343478934,
				pathTmplt:     "plain/%s.method_%d.save",
				pathExelTmplt: "exelplain/%s.method_%d.save",
			},
		},
		{
			name: "Test 3",
			args: args{
				method:        a51.Method1,
				filename:      "cipher/img.method_1.save.txt",
				key:           126217892231,
				pathTmplt:     "plain/%s.method_%d.save",
				pathExelTmplt: "exelplain/%s.method_%d.save",
			},
		},
		{
			name: "Test 4",
			args: args{
				method:        a51.Method2,
				filename:      "cipher/img.method_2.save.txt",
				key:           437893983487323,
				pathTmplt:     "plain/%s.method_%d.save",
				pathExelTmplt: "exelplain/%s.method_%d.save",
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := app.New(tt.args.method, tt.args.filename, tt.args.key)
			a.Run()
			err := a.Save(tt.args.pathTmplt, tt.args.pathExelTmplt)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}
