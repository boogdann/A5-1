package files

import "os"

func Save(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	saveData := make([]byte, len(data)/8)
	for i, bit := range data {
		byteIndex := i / 8
		bitIndex := (8 - 1) - uint(i%8)

		if bit == 1 {
			saveData[byteIndex] |= 1 << bitIndex
		}
	}

	_, err = file.Write(saveData)

	file, err = os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.Write(saveData)

	return err
}
