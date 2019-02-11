package json_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/toolsparty/mvc"
	"github.com/valyala/fasthttp"
)

// implements abstract controller
type JSONController struct {
	*mvc.BaseController

	view mvc.View
}

func (b *JSONController) Init() error {
	return b.BaseController.Init()
}

// Context returns *fasthttp.RequestCtx (from interface{}) or error
func (JSONController) Context(v interface{}) (*fasthttp.RequestCtx, error) {
	ctx, ok := v.(*fasthttp.RequestCtx)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%T is not *fasthttp.RequestCtx", v))
	}

	return ctx, nil
}

// Render encode interface{} to JSON and write to writer
func (b JSONController) Render(w io.Writer, v interface{}) error {
	b.SetHeaders(w)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	if err != nil {
		return errors.Wrap(err, "encoding to json failed")
	}

	return nil
}

// Decode struct from bytes
func (JSONController) Decode(b []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(v)
	if err != nil {
		return errors.Wrap(err, "decoding json failed")
	}

	return nil
}

// Set content-type header for json
func (JSONController) SetHeaders(w io.Writer) {
	wr, ok := w.(*fasthttp.RequestCtx)
	if ok {
		wr.SetContentType("application/json")
	}
}
