package hash

import (
	"testing"

	"github.com/go4x/got"
)

func TestSHA256(t *testing.T) {
	tests := []got.Case{
		got.NewCase("empty string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", false, nil),
		got.NewCase("a word", "hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", false, nil),
		got.NewCase("a sentence", "The quick brown fox jumps over the lazy dog", "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592", false, nil),
		got.NewCase("a sentence with a period", "The quick brown fox jumps over the lazy dog.", "ef537f25c895bfa782526529a9b63d97aa631564d5d789c2b765448c8635fb6c", false, nil),
		got.NewCase("a single character", "a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb", false, nil),
		got.NewCase("a number string", "1234567890", "c775e7b757ede630cd0aa1113bd102661ab38829ca52a6422ab782862f268646", false, nil),
		got.NewCase("a space", " ", "36a9e7f1c95b82ffb99743e0c5c4ce95d83c9a430aac59f84ef3cbfab6145068", false, nil),
		got.NewCase("a newline symbol", "\n", "01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b", false, nil),
	}

	r := got.New(t, "SHA256")
	r.Cases(tests, func(c got.Case, tt *testing.T) {
		result := SHA256(c.Input().(string))
		if result != c.Want().(string) {
			r.Fail("SHA256(%q) = %q; want %q", c.Input().(string), result, c.Want().(string))
		} else {
			r.Pass("SHA256(%q) = %q; want %q", c.Input().(string), result, c.Want().(string))
		}
	})

	// Test long string
	r.Case("long string")
	longInput := ""
	for i := 0; i < 10000; i++ {
		longInput += "a"
	}
	expectedLong := "27dd1f61b867b6a0f6e9d8a41c43231de52107e53ae424de8f847b821db4b711"
	resultLong := SHA256(longInput)
	r.Logf("resultLong: %s", resultLong)
	if resultLong != expectedLong {
		t.Errorf("SHA256(longInput) = %q; want %q", resultLong, expectedLong)
	}

	// Test unicode
	r.Case("unicode string")
	unicodeInput := "你好，世界"
	expectedUnicode := "46932f1e6ea5216e77f58b1908d72ec9322ed129318c6d4bd4450b5eaab9d7e7"
	resultUnicode := SHA256(unicodeInput)
	r.Logf("resultUnicode: %s", resultUnicode)
	if resultUnicode != expectedUnicode {
		t.Errorf("SHA256(%q) = %q; want %q", unicodeInput, resultUnicode, expectedUnicode)
	}
}
