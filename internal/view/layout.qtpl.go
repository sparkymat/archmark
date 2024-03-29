// Code generated by qtc from "layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/layout.qtpl:1
package view

//line view/layout.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/layout.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/layout.qtpl:1
func StreamLayout(qw422016 *qt422016.Writer, title string, csrfToken string, content string) {
//line view/layout.qtpl:1
	qw422016.N().S(`
  <!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width,initial-scale=1">
      <meta name="csrf-token" content="`)
//line view/layout.qtpl:7
	qw422016.E().S(csrfToken)
//line view/layout.qtpl:7
	qw422016.N().S(`">
      <title>`)
//line view/layout.qtpl:8
	qw422016.E().S(title)
//line view/layout.qtpl:8
	qw422016.N().S(`</title>
      <link rel="stylesheet" type="text/css" href="/css/uikit.min.css">
      <link rel="stylesheet" type="text/css" href="/css/fonts.css">
      <link rel="stylesheet" type="text/css" href="/css/style.css">
    </head>
    <body>
      `)
//line view/layout.qtpl:14
	qw422016.N().S(content)
//line view/layout.qtpl:14
	qw422016.N().S(`
      <script src="/js/uikit.min.js"></script>
      <script src="/js/uikit-icons.min.js"></script>
      <script src="/js/app/index.js"></script>
    </body>
  </html>
`)
//line view/layout.qtpl:20
}

//line view/layout.qtpl:20
func WriteLayout(qq422016 qtio422016.Writer, title string, csrfToken string, content string) {
//line view/layout.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/layout.qtpl:20
	StreamLayout(qw422016, title, csrfToken, content)
//line view/layout.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line view/layout.qtpl:20
}

//line view/layout.qtpl:20
func Layout(title string, csrfToken string, content string) string {
//line view/layout.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line view/layout.qtpl:20
	WriteLayout(qb422016, title, csrfToken, content)
//line view/layout.qtpl:20
	qs422016 := string(qb422016.B)
//line view/layout.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line view/layout.qtpl:20
	return qs422016
//line view/layout.qtpl:20
}
