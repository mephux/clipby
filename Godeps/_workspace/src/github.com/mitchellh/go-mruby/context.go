package mruby

// #include "gomruby.h"
import "C"

// CompileContext represents a context for code compilation.
//
// CompileContexts keep track of things such as filenames, line numbers,
// as well as some settings for how to parse and execute code.
type CompileContext struct {
	ctx      *C.mrbc_context
	filename string
	mrb      *Mrb
}

func NewCompileContext(m *Mrb) *CompileContext {
	return &CompileContext{
		ctx: C.mrbc_context_new(m.state),
		mrb: m,
	}
}

// Closes the context, freeing any resources associated with it.
//
// This is safe to call once the context has been used for parsing/loading
// any Ruby code.
func (c *CompileContext) Close() {
	C.mrbc_context_free(c.mrb.state, c.ctx)
}

// Filename returns the filename associated with this context.
func (c *CompileContext) Filename() string {
	return C.GoString(c.ctx.filename)
}

// SetFilename sets the filename associated with this compilation context.
//
// Code parsed under this context will be from this file.
func (c *CompileContext) SetFilename(f string) {
	c.filename = f
	c.ctx.filename = C.CString(c.filename)
}
