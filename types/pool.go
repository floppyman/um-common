package types

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	sync.Pool
}

// NewBufferPool creates a new buffer pool
//
//goland:noinspection GoUnusedExportedFunction
func NewBufferPool() *BufferPool {
	return &BufferPool{
		Pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 128))
			b.Reset()
			return b
		}},
	}
}

func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.Pool.Get().(*bytes.Buffer)
}

func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	bp.Pool.Put(b)
}
