package gost

// #include "gosthash.h"
import "C"
import (
	"hash"
	"unsafe"
)

func init() {
	C.gosthash_init()
}

type gostHash struct {
	ctx C.GostHashCtx
}

func (g *gostHash) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		C.gosthash_update(&g.ctx, (*C.uchar)(&p[0]), (C.size_t)(len(p)))
	} else {
		C.gosthash_update(&g.ctx, nil, 0)
	}
	return len(p), nil
}

func (g *gostHash) Sum(b []byte) []byte {
	var digest [32]C.uchar
	newCtx := g.ctx
	C.gosthash_final(&newCtx, &digest[0])
	return append(b, C.GoBytes(unsafe.Pointer(&digest[0]), 32)...)
}

func (g *gostHash) Reset() {
	C.gosthash_reset(&g.ctx)
}

func (g *gostHash) Size() int {
	return 32
}

func (g *gostHash) BlockSize() int {
	return 32
}

func New() hash.Hash {
	h := new(gostHash)
	h.Reset()
	return h
}
