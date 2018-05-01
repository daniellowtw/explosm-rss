package explosm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test2018(t *testing.T) {
	var response = `<div id="main-left">
<section id="comic-area">
<div id="comic-wrap">
<img id="main-comic" src="//files.explosm.net/comics/Dave/genieweek3.png?t=06538A">
</div>
<div id="comic-under">
`
	assert.True(t, imgRegexp.MatchString(response), "Regex has changed")
}

