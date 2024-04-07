package mjml2html

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"io"
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func init() {
	v := pool.Get()
	defer pool.Put(v)
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

//go:embed dist/main.js.gz
var source []byte

var registry = require.NewRegistryWithLoader(func(path string) ([]byte, error) {
	if source == nil {
		panic("something went wrong")
	}

	if path != "node_modules/mjml2html" {
		panic("unexpected path: " + path)
	}

	r := must(gzip.NewReader(bytes.NewReader(source)))
	defer func() {
		must(0, r.Close())
		source = nil // free up memory
	}()
	return must(io.ReadAll(r)), nil
})

type poolEntry struct {
	rt            *goja.Runtime
	requiredValue goja.Value
}

var pool = sync.Pool{
	New: func() any {
		rt := goja.New()
		rt.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
		req := registry.Enable(rt)
		return &poolEntry{
			rt:            rt,
			requiredValue: must(req.Require("mjml2html")).ToObject(rt).Get("default"),
		}
	},
}

type config map[string]any

type option func(*config)

func Fonts(f map[string]string) option {
	return func(c *config) {
		(*c)["fonts"] = f
	}
}

func KeepComments(b bool) option {
	return func(c *config) {
		(*c)["keepComments"] = b
	}
}

func ValidationLevel(l string) option {
	return func(c *config) {
		(*c)["validationLevel"] = l
	}
}

type MJMLNode struct {
	TagName    string            `json:"tagName"`
	Attributes map[string]string `json:"attributes"`
	Content    string            `json:"content,omitempty"`
	Children   []MJMLNode        `json:"children,omitempty"`
}

func ToHTML[T ~string | *MJMLNode](v T, opts ...option) (string, error) {
	var empty T
	if v == empty {
		return "", nil
	}

	entry := pool.Get().(*poolEntry)
	defer pool.Put(entry)

	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}

	inner, _ := goja.AssertFunction(entry.requiredValue)
	result, err := inner(goja.Undefined(), entry.rt.ToValue(v), entry.rt.ToValue(cfg))
	if err != nil {
		return "", err
	}

	ret := ""
	if err := entry.rt.ExportTo(result.ToObject(entry.rt).Get("html"), &ret); err != nil {
		return "", err
	}

	return ret, nil
}

func ToJSON(v string, opts ...option) (*MJMLNode, error) {
	rootNode := MJMLNode{TagName: "mjml"}
	if v == "" {
		return &rootNode, nil
	}

	entry := pool.Get().(*poolEntry)
	defer pool.Put(entry)

	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}

	inner, _ := goja.AssertFunction(entry.requiredValue)
	result, err := inner(goja.Undefined(), entry.rt.ToValue(v), entry.rt.ToValue(cfg))
	if err != nil {
		return nil, err
	}

	if err := entry.rt.ExportTo(result.ToObject(entry.rt).Get("json"), &rootNode); err != nil {
		return nil, err
	}

	return &rootNode, nil
}
