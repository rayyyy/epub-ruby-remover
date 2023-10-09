package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveRuby(t *testing.T) {
	testCases := []struct {
		html     string
		expected string
	}{
		// no ruby
		{
			html:     `<p>test</p>`,
			expected: `<p>test</p>`,
		},
		// ruby
		{
			html:     `<p><ruby>漢<rt>かん</rt></ruby></p>`,
			expected: `<p>漢</p>`,
		},
		// multi ruby
		{
			html:     `<p><ruby>漢<rt>かん</rt></ruby><ruby>脳<rt>のう</rt>裏<rt>り</rt></ruby></p>`,
			expected: `<p>漢脳裏</p>`,
		},
		// full ruby
		{
			html:     `<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" xml:lang="ja" class="vrtl"><head><meta charset="UTF-8"/><title>title</title><link rel="stylesheet" type="text/css" href="../Styles/x.css"/></head><body class="p-text"><div class="main"><p>test</p><p><ruby>漢<rt>かん</rt></ruby></p><p><ruby>漢<rt>かん</rt></ruby><ruby>脳<rt>のう</rt>裏<rt>り</rt></ruby></p><p>test</p></div></body></html>`,
			expected: `<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" xml:lang="ja" class="vrtl"><head><meta charset="UTF-8"/><title>title</title><link rel="stylesheet" type="text/css" href="../Styles/x.css"/></head><body class="p-text"><div class="main"><p>test</p><p>漢</p><p>漢脳裏</p><p>test</p></div></body></html>`,
		},
		// no ruby kobo
		{
			html:     `<p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">test</span></p>`,
			expected: `<p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">test</span></p>`,
		},
		// ruby kobo
		{
			html:     `<p><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">漢</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.3.1">かん</span></rt></ruby></p>`,
			expected: `<p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">漢</span></p>`,
		},
		// multi ruby kobo
		{
			html:     `<p><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">漢</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.2.1">かん</span></rt></ruby><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.3.1">脳</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.4.1">のう</span></rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.5.1">裏</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.6.1">り</span></rt></ruby></p>`,
			expected: `<p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">漢</span><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.3.1">脳</span><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.5.1">裏</span></p>`,
		},
		// full ruby kobo
		{
			html:     `<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" xml:lang="ja" class="vrtl"><head><meta charset="UTF-8" /><title>title</title><link rel="stylesheet" type="text/css" href="../Styles/x.css" /></head><body class="p-text"><div class="main"><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">test</span></p><p><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.2.1">漢</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.3.1">かん</span></rt></ruby></p><p><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.4.1">漢</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.5.1">かん</span></rt></ruby><ruby><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.6.1">脳</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.7.1">のう</span></rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.8.1">裏</span><rt><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.9.1">り</span></rt></ruby></p><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.10.1">test</span></p></div></body></html>`,
			expected: `<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" xml:lang="ja" class="vrtl"><head><meta charset="UTF-8" /><title>title</title><link rel="stylesheet" type="text/css" href="../Styles/x.css" /></head><body class="p-text"><div class="main"><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">test</span></p><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.2.1">漢</span></p><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.4.1">漢</span><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.6.1">脳</span><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.8.1">裏</span></p><p><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.10.1">test</span></p></div></body></html>`,
		},
	}

	for _, testCase := range testCases {
		actual := RemoveRuby(testCase.html)
		assert.Equal(t, testCase.expected, actual)
	}
}
