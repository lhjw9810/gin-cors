package bufferpool

import (
	"github.com/valyala/bytebufferpool"
	"io"
)

/**
改进bufferpool中的copy方法，不使用sync.pool,使用bytebufferpool
*/
func Copy2(dst io.Writer, src io.Reader) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	bytebuffer := bytebufferpool.Get()
	defer bytebufferpool.Put(bytebuffer)
	if n, err := bytebuffer.ReadFrom(src); err != nil {
		return n, err
	}
	return bytebuffer.WriteTo(dst)
}
