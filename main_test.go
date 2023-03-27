package dutf

import (
	"bytes"
	"testing"
)

const utf8code = "\u0041\u2262\u0391\u002E"

var dutfcode = []byte{0x41, 0xE2, 0x44, 0xF3, 0x43, 0x2E}

func TestEncode(t *testing.T) {
	result := EncodeString(utf8code)

	if !bytes.Equal(result, dutfcode) {
		t.Fatalf("expect %v, but %v", dutfcode, result)
	}
}

func TestDecode(t *testing.T) {
	result, undecoded := DecodeString(dutfcode)
	if undecoded != nil {
		t.Fatalf("all DUTF code is not decoded. The rest is %v", undecoded)
	}
	if result != utf8code {
		t.Fatalf("expect %v, but %v", utf8code, result)
	}
}
