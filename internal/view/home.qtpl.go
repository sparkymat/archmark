// Code generated by qtc from "home.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/view/home.qtpl:1
package view

//line internal/view/home.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/view/home.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/view/home.qtpl:1
func StreamHome(qw422016 *qt422016.Writer) {
//line internal/view/home.qtpl:1
	qw422016.N().S(`
  <div id="archmark-app"></div>
`)
//line internal/view/home.qtpl:3
}

//line internal/view/home.qtpl:3
func WriteHome(qq422016 qtio422016.Writer) {
//line internal/view/home.qtpl:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/view/home.qtpl:3
	StreamHome(qw422016)
//line internal/view/home.qtpl:3
	qt422016.ReleaseWriter(qw422016)
//line internal/view/home.qtpl:3
}

//line internal/view/home.qtpl:3
func Home() string {
//line internal/view/home.qtpl:3
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/view/home.qtpl:3
	WriteHome(qb422016)
//line internal/view/home.qtpl:3
	qs422016 := string(qb422016.B)
//line internal/view/home.qtpl:3
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/view/home.qtpl:3
	return qs422016
//line internal/view/home.qtpl:3
}
