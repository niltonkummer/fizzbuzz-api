package repository

import "github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"

type CacheFizzbuzzNoOp struct{}

func NewCacheFizzbuzzNoOp() adapters.CacheFizzbuzz {
	return &CacheFizzbuzzNoOp{}
}

func (c *CacheFizzbuzzNoOp) Get(key string) (string, error) {
	return "", nil
}
func (c *CacheFizzbuzzNoOp) Set(key string, value string) error {
	return nil
}
