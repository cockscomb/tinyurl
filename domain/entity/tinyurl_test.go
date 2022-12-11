package entity

import (
	"github.com/cockscomb/tinyurl/util"
	"gotest.tools/v3/assert"
	"net/url"
	"testing"
)

func TestGenerateTinyURL(t *testing.T) {
	tu := GenerateTinyURL(util.Must(url.Parse("https://example.com")))
	assert.Assert(t, tu.ID != "")
	assert.Assert(t, tu.URL.String() == "https://example.com")
}
