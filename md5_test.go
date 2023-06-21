package gox

import (
	"testing"
)

func TestMD5(t *testing.T) {
	if `1790c0449c346a45cbe511a76de9f7a2` == MD5("6c:0d:c4:0f:7a:15") {
		t.Log(`ok`)
	}
}
