// Code generated by qtc from "layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/layout.qtpl:1
package view

//line view/layout.qtpl:1
import "github.com/sparkymat/archmark/localize"

//line view/layout.qtpl:2
import "github.com/sparkymat/archmark/style"

//line view/layout.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/layout.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/layout.qtpl:4
func StreamLayout(qw422016 *qt422016.Writer, styler *style.Service, localizer *localize.Service, lang localize.Language, title string, content string) {
//line view/layout.qtpl:4
	qw422016.N().S(`
  <!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width,initial-scale=1">
      <title>`)
//line view/layout.qtpl:10
	qw422016.E().S(title)
//line view/layout.qtpl:10
	qw422016.N().S(`</title>
      <link rel="stylesheet" type="text/css" href="/css/tailwind-2.min.css">
      <link rel="stylesheet" type="text/css" href="/css/style.css">
    </head>
    <body>
      <nav class="bg-gray-800">
        <div class="flex flex-row align-center">
          <span class="text-white self-center px-4">archmark!</span>
          <a href="/" class="text-white px-4 py-2 rounded-md text-sm font-medium">`)
//line view/layout.qtpl:18
	qw422016.E().S(localizer.Lookup(lang, localize.Bookmarks))
//line view/layout.qtpl:18
	qw422016.N().S(`</a>
          <a href="/add" class="text-white px-4 py-2 rounded-md text-sm font-medium">`)
//line view/layout.qtpl:19
	qw422016.E().S(localizer.Lookup(lang, localize.AddNew))
//line view/layout.qtpl:19
	qw422016.N().S(`</a>
          <a href="/tokens" class="text-white px-4 py-2 rounded-md text-sm font-medium">`)
//line view/layout.qtpl:20
	qw422016.E().S(localizer.Lookup(lang, localize.APITokens))
//line view/layout.qtpl:20
	qw422016.N().S(`</a>
          <div class="flex-grow"></div>
          <a href="/settings" class="text-white px-4 py-2 rounded-md text-sm font-medium">`)
//line view/layout.qtpl:22
	qw422016.E().S(localizer.Lookup(lang, localize.Settings))
//line view/layout.qtpl:22
	qw422016.N().S(`</a>
        </div>
      </nav>
      `)
//line view/layout.qtpl:25
	qw422016.N().S(content)
//line view/layout.qtpl:25
	qw422016.N().S(`
      <script src="/javascript/index.js"></script>
    </body>
  </html>
`)
//line view/layout.qtpl:29
}

//line view/layout.qtpl:29
func WriteLayout(qq422016 qtio422016.Writer, styler *style.Service, localizer *localize.Service, lang localize.Language, title string, content string) {
//line view/layout.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/layout.qtpl:29
	StreamLayout(qw422016, styler, localizer, lang, title, content)
//line view/layout.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line view/layout.qtpl:29
}

//line view/layout.qtpl:29
func Layout(styler *style.Service, localizer *localize.Service, lang localize.Language, title string, content string) string {
//line view/layout.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line view/layout.qtpl:29
	WriteLayout(qb422016, styler, localizer, lang, title, content)
//line view/layout.qtpl:29
	qs422016 := string(qb422016.B)
//line view/layout.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line view/layout.qtpl:29
	return qs422016
//line view/layout.qtpl:29
}
