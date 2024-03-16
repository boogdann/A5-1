package ciphering

import "2/internal/a51"

type Cipher struct {
	a51 a51.A51
}

func New(a51 a51.A51) *Cipher {
	return &Cipher{
		a51: a51,
	}
}

func (c *Cipher) Encrypt(data []byte) ([]byte, []byte) {
	key := c.a51.GenerateKeyStream(len(data))

	return c.crypt(data, key), key
}

func (c *Cipher) crypt(data []byte, key []byte) []byte {
	newData := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		newData[i] = data[i] ^ key[i]
	}

	return newData
}

func (c *Cipher) Decrypt(data []byte, key []byte) []byte {
	return c.crypt(data, key)
}
