package explosm

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRegexp(t *testing.T) {
	var validResponse = `<li><a href="/comics/latest/" class="" title="Latest comic"><img src="/img/nav-button_newest@2x.png"/></a></li>
</ul>
</div>
</div>
</div>
<div id="comic-container">
<div class="row">
<div class="small-12 medium-12 large-12 columns">
<a href="http://explosm.net/show/episode/201/eden">
<img id="main-comic" src="//files.explosm.net/comics/Dave/Eden_Thumbnail_728x410.png?t=FDF100"/>
</a>
</div>
</div>
</div>
<div class="row">
<div class="small-12 medium-6 large-6 columns">
<div class="row collapse">`
	res := r.FindSubmatch([]byte(validResponse))
	spew.Dump(res)
}
