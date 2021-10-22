// Code generated by qtc from "api_tokens_index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/api_tokens_index.qtpl:1
package view

//line view/api_tokens_index.qtpl:1
import "github.com/sparkymat/archmark/presenter"

//line view/api_tokens_index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/api_tokens_index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/api_tokens_index.qtpl:3
func StreamApiTokensIndex(qw422016 *qt422016.Writer, tokens []presenter.APIToken) {
//line view/api_tokens_index.qtpl:3
	qw422016.N().S(`
  <div class="container mx-auto">
    <div class="flex flex-col mt-4">
      <div class="flex flex-row justify-end">
        <form action="/tokens" method="POST">
          <input type="submit" value="Create new token" class="text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-4 py-2">
        </form>
      </div>
      <ul>
        `)
//line view/api_tokens_index.qtpl:12
	for _, token := range tokens {
//line view/api_tokens_index.qtpl:12
		qw422016.N().S(`
          <li class="md:mx-0 mx-4 p-4 mt-4 border-2 border-dashed flex flex-row justify-between items-center rounded">
            <span class="text-md italic break-all">`)
//line view/api_tokens_index.qtpl:14
		qw422016.E().S(token.Token)
//line view/api_tokens_index.qtpl:14
		qw422016.N().S(`</span>
            <form action="/tokens/`)
//line view/api_tokens_index.qtpl:15
		qw422016.E().S(token.ID)
//line view/api_tokens_index.qtpl:15
		qw422016.N().S(`/destroy" method="POST">
              <input type="submit" value="Delete" class="text-l text-white bg-red-700 hover:bg-red-900 rounded shadow-md px-4 py-2">
            </form>
          </li>
        `)
//line view/api_tokens_index.qtpl:19
	}
//line view/api_tokens_index.qtpl:19
	qw422016.N().S(`
      </ul>
    </div>
  </div>
`)
//line view/api_tokens_index.qtpl:23
}

//line view/api_tokens_index.qtpl:23
func WriteApiTokensIndex(qq422016 qtio422016.Writer, tokens []presenter.APIToken) {
//line view/api_tokens_index.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/api_tokens_index.qtpl:23
	StreamApiTokensIndex(qw422016, tokens)
//line view/api_tokens_index.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line view/api_tokens_index.qtpl:23
}

//line view/api_tokens_index.qtpl:23
func ApiTokensIndex(tokens []presenter.APIToken) string {
//line view/api_tokens_index.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line view/api_tokens_index.qtpl:23
	WriteApiTokensIndex(qb422016, tokens)
//line view/api_tokens_index.qtpl:23
	qs422016 := string(qb422016.B)
//line view/api_tokens_index.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line view/api_tokens_index.qtpl:23
	return qs422016
//line view/api_tokens_index.qtpl:23
}
