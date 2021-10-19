// Code generated by qtc from "bookmarks_new.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/bookmarks_new.qtpl:1
package view

//line view/bookmarks_new.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/bookmarks_new.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/bookmarks_new.qtpl:1
func StreamBookmarksNew(qw422016 *qt422016.Writer) {
//line view/bookmarks_new.qtpl:1
	qw422016.N().S(`
  <div class="container mx-auto">
    <form action="/bookmarks" method="POST">
      <div class="flex flex-row mt-4">
        <input type="text" name="url" class="text-xl flex-grow p-2 border outline-white" placeholder="Paste URL here..." autofocus>
        <input type="submit" value="Add" class="text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-8 py-2 ml-4">
      </div>
    </form>
    <p class="text-l p-4 mt-4 border-2 border-gray-400 text-gray-400 border-dashed italic">Please note that it might take a while for the archive to be generated. In the meantime, you can use the 'original' link.</p>
  </div>
`)
//line view/bookmarks_new.qtpl:11
}

//line view/bookmarks_new.qtpl:11
func WriteBookmarksNew(qq422016 qtio422016.Writer) {
//line view/bookmarks_new.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/bookmarks_new.qtpl:11
	StreamBookmarksNew(qw422016)
//line view/bookmarks_new.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line view/bookmarks_new.qtpl:11
}

//line view/bookmarks_new.qtpl:11
func BookmarksNew() string {
//line view/bookmarks_new.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
//line view/bookmarks_new.qtpl:11
	WriteBookmarksNew(qb422016)
//line view/bookmarks_new.qtpl:11
	qs422016 := string(qb422016.B)
//line view/bookmarks_new.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
//line view/bookmarks_new.qtpl:11
	return qs422016
//line view/bookmarks_new.qtpl:11
}
