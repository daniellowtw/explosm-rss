package explosm

import (
	"testing"
)

var testCases = map[string]string{
	"with sapces": `<li><a class="disabled" title="Latest comic"><img src="/img/nav-button_newest@2x.png"/></a></li> </ul> </div> </div> </div> <div id="comic-container"> <div class="row"> <div class="small-12 medium-12 large-12 columns"> <img id="main-comic" src="//files.explosm.net/comics/Kris/mailman2.png?t=0AB34F" /> </div> </div> </div> <div class="row"> <div class="small-12 medium-6 large-6 columns"> <div class="row collapse">`,
	"with new lines": `<li><a href="/comics/latest/" class="" title="Latest comic"><img src="/img/nav-button_newest@2x.png"/></a></li>
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
<div class="row collapse">`,
}

func TestRegexp(t *testing.T) {
	for name, tc := range testCases {
		res := FindComicURL([]byte(tc))
		if res == "" {
			t.Logf("failed to find image for test case %s", name)
			t.Fail()
		}
	}
}

