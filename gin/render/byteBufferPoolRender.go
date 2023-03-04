// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"io"
	"net/http"
	"strconv"

	"github.com/lhjw9810/go-infra/v2/bufferpool"
)

// BufferPoolRender contains the IO reader and its length, and custom ContentType and other headers.
type ByteBufferPoolRender struct {
	ContentType   string
	ContentLength int64
	Reader        io.Reader
	Headers       map[string]string
}

// Render (BufferPoolRender) writes data with custom ContentType and headers.
func (r ByteBufferPoolRender) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	if r.ContentLength >= 0 {
		if r.Headers == nil {
			r.Headers = map[string]string{}
		}
		r.Headers["Content-Length"] = strconv.FormatInt(r.ContentLength, 10)
	}
	r.writeHeaders(w, r.Headers)
	_, err = bufferpool.Copy2(w, r.Reader)
	return
}

// WriteContentType (BufferPoolRender) writes custom ContentType.
func (r ByteBufferPoolRender) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, []string{r.ContentType})
}

// writeHeaders writes custom Header.
func (r ByteBufferPoolRender) writeHeaders(w http.ResponseWriter, headers map[string]string) {
	header := w.Header()
	for k, v := range headers {
		if header.Get(k) == "" {
			header.Set(k, v)
		}
	}
}
