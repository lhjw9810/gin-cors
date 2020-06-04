package bufferpool

import (
	"io"
	"sync"
)

// A BufferPool is an interface for getting and returning temporary
// byte slices for use by io.CopyBuffer.
type BufferPool interface {
	Get() []byte
	Put([]byte)
}

var bufferPool BufferPool

//字节缓存池
type bytePool struct {
	pool *sync.Pool
}

func (c *bytePool) Get() []byte {
	return c.pool.Get().([]byte)
}
func (c *bytePool) Put(bytes []byte) {
	c.pool.Put(bytes)
}

//基于pool 的io.copy
func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	buf := bufferPool.Get()
	defer bufferPool.Put(buf)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

func init() {
	once := &sync.Once{}
	once.Do(func() {
		bytePool := &bytePool{
			pool: &sync.Pool{},
		}
		bytePool.pool.New = func() interface{} {
			return make([]byte, 32*1024)
		}
		bufferPool = bytePool
	})
}
