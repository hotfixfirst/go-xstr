package xstr

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateSpaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single space",
			input:    " ",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello     world",
			expected: "hello world",
		},
		{
			name:     "leading spaces",
			input:    "   hello world",
			expected: "hello world",
		},
		{
			name:     "trailing spaces",
			input:    "hello world   ",
			expected: "hello world",
		},
		{
			name:     "leading and trailing spaces",
			input:    "   hello world   ",
			expected: "hello world",
		},
		{
			name:     "tabs between words",
			input:    "hello\t\t\tworld",
			expected: "hello world",
		},
		{
			name:     "newlines between words",
			input:    "hello\n\n\nworld",
			expected: "hello world",
		},
		{
			name:     "carriage returns between words",
			input:    "hello\r\r\rworld",
			expected: "hello world",
		},
		{
			name:     "vertical tabs between words",
			input:    "hello\v\v\vworld",
			expected: "hello world",
		},
		{
			name:     "form feeds between words",
			input:    "hello\f\f\fworld",
			expected: "hello world",
		},
		{
			name:     "mixed whitespace characters",
			input:    "hello \t\n\r\v\f world",
			expected: "hello world",
		},
		{
			name:     "zero-width space removal",
			input:    "hello\u200Bworld",
			expected: "helloworld",
		},
		{
			name:     "BOM removal",
			input:    "hello\uFEFFworld",
			expected: "helloworld",
		},
		{
			name:     "word joiner removal",
			input:    "hello\u2060world",
			expected: "helloworld",
		},
		{
			name:     "zero-width joiner removal",
			input:    "hello\u200Dworld",
			expected: "helloworld",
		},
		{
			name:     "LTR mark removal",
			input:    "hello\u200Eworld",
			expected: "helloworld",
		},
		{
			name:     "RTL mark removal",
			input:    "hello\u200Fworld",
			expected: "helloworld",
		},
		{
			name:     "non-breaking space removal",
			input:    "hello\u00A0world",
			expected: "helloworld",
		},
		{
			name:     "multiple zero-width chars",
			input:    "hello\u200B\uFEFF\u2060world",
			expected: "helloworld",
		},
		{
			name:     "zero-width chars with spaces",
			input:    "hello \u200B \uFEFF world",
			expected: "hello world",
		},
		{
			name:     "complex mixed input",
			input:    "  \t hello \u200B \n\n world \uFEFF \r\r test  \t ",
			expected: "hello world test",
		},
		{
			name:     "only whitespace",
			input:    "   \t\n\r\v\f   ",
			expected: "",
		},
		{
			name:     "only zero-width chars",
			input:    "\u200B\uFEFF\u2060\u200D\u200E\u200F\u00A0",
			expected: "",
		},
		{
			name:     "unicode text with spaces",
			input:    "สวัสดี    โลก",
			expected: "สวัสดี โลก",
		},
		{
			name:     "unicode with zero-width chars",
			input:    "สวัสดี\u200Bโลก",
			expected: "สวัสดีโลก",
		},
		{
			name:     "numbers and symbols",
			input:    "123   $%^   456",
			expected: "123 $%^ 456",
		},
		{
			name:     "single character with spaces",
			input:    "   a   ",
			expected: "a",
		},
		{
			name:     "multiple words with mixed separators",
			input:    "one\ttwo\nthree\rfour\vfive\fsix",
			expected: "one two three four five six",
		},
		// JSON format tests
		{
			name:     "json with extra spaces",
			input:    `{  "name"  :   "John"  ,  "age"  :   25  }`,
			expected: `{ "name" : "John" , "age" : 25 }`,
		},
		{
			name:     "json with newlines and tabs",
			input:    "{\n\t\"name\"\t:\t\"John\",\n\t\"age\"\t:\t25\n}",
			expected: `{ "name" : "John", "age" : 25 }`,
		},
		{
			name:     "json array with spaces",
			input:    `[  "item1"  ,   "item2"   ,  "item3"  ]`,
			expected: `[ "item1" , "item2" , "item3" ]`,
		},
		{
			name:     "json nested with zero-width chars",
			input:    "{\"user\"\u200B:\u200B{\"name\"\u200B:\u200B\"John\"}}",
			expected: "{\"user\":{\"name\":\"John\"}}",
		},
		// HTML format tests
		{
			name:     "html with extra spaces",
			input:    `<div   class="container"  >  Hello   World  </div>`,
			expected: `<div class="container" > Hello World </div>`,
		},
		{
			name:     "html with newlines and tabs",
			input:    "<html>\n\t<body>\n\t\t<h1>Title</h1>\n\t</body>\n</html>",
			expected: "<html> <body> <h1>Title</h1> </body> </html>",
		},
		{
			name:     "html attributes with spaces",
			input:    `<img   src="image.jpg"   alt="  description  "   />`,
			expected: `<img src="image.jpg" alt=" description " />`,
		},
		{
			name:     "html with zero-width chars",
			input:    "<div\u200Bclass\u200B=\u200B\"test\"\u200B>content</div>",
			expected: "<divclass=\"test\">content</div>",
		},
		{
			name:     "html script tag with spaces",
			input:    `<script   type="text/javascript"  >  alert(  'hello'  );  </script>`,
			expected: `<script type="text/javascript" > alert( 'hello' ); </script>`,
		},
		// XML format tests
		{
			name:     "xml with extra spaces",
			input:    `<root   xmlns="namespace"  >  <item   id="1"  >  value  </item>  </root>`,
			expected: `<root xmlns="namespace" > <item id="1" > value </item> </root>`,
		},
		{
			name:     "xml with newlines and tabs",
			input:    "<?xml version=\"1.0\"?>\n\t<root>\n\t\t<item>value</item>\n\t</root>",
			expected: `<?xml version="1.0"?> <root> <item>value</item> </root>`,
		},
		{
			name:     "xml with CDATA and spaces",
			input:    `<data>  <![CDATA[  some   content   ]]>  </data>`,
			expected: `<data> <![CDATA[ some content ]]> </data>`,
		},
		{
			name:     "xml with zero-width chars",
			input:    "<item\u200Bid\u200B=\u200B\"123\"\u200B>text</item>",
			expected: "<itemid=\"123\">text</item>",
		},
		{
			name:     "xml with mixed whitespace in content",
			input:    "<message>\n\t  Hello  \r\n  World  \t\n</message>",
			expected: "<message> Hello World </message>",
		},
		// Complex structured data tests
		{
			name:     "json with html content",
			input:    `{  "html"  :  "<div   class='test'  >  content  </div>"  }`,
			expected: `{ "html" : "<div class='test' > content </div>" }`,
		},
		{
			name:     "xml with json-like content",
			input:    `<config>  {  "setting"  :  "value"  }  </config>`,
			expected: `<config> { "setting" : "value" } </config>`,
		},
		{
			name:     "nested xml with attributes",
			input:    "<root>\n  <user   id=\"1\"  name=\"John\"  >\n    <profile   active=\"true\"  />\n  </user>\n</root>",
			expected: `<root> <user id="1" name="John" > <profile active="true" /> </user> </root>`,
		},
		{
			name:     "malformed json with extra spaces",
			input:    `{   "incomplete"  :   `,
			expected: `{ "incomplete" :`,
		},
		{
			name:     "html comments with spaces",
			input:    `<!--   This is a comment   -->  <div>content</div>`,
			expected: `<!-- This is a comment --> <div>content</div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveDuplicateSpaces(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "space character",
			input:    ' ',
			expected: true,
		},
		{
			name:     "tab character",
			input:    '\t',
			expected: true,
		},
		{
			name:     "newline character",
			input:    '\n',
			expected: true,
		},
		{
			name:     "carriage return",
			input:    '\r',
			expected: true,
		},
		{
			name:     "vertical tab",
			input:    '\v',
			expected: true,
		},
		{
			name:     "form feed",
			input:    '\f',
			expected: true,
		},
		{
			name:     "regular letter",
			input:    'a',
			expected: false,
		},
		{
			name:     "number",
			input:    '1',
			expected: false,
		},
		{
			name:     "symbol",
			input:    '$',
			expected: false,
		},
		{
			name:     "unicode character",
			input:    'ก',
			expected: false,
		},
		{
			name:     "zero-width space (not whitespace in isWhitespace)",
			input:    '\u200B',
			expected: false,
		},
		{
			name:     "non-breaking space (not whitespace in isWhitespace)",
			input:    '\u00A0',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isWhitespace(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Edge case tests for performance and memory efficiency
func TestRemoveDuplicateSpaces_EdgeCases(t *testing.T) {
	t.Run("very long string with many spaces", func(t *testing.T) {
		// Create a string with many consecutive spaces
		input := "start" + strings.Repeat(" ", 1000) + "middle" + strings.Repeat("\t", 500) + "end"
		expected := "start middle end"

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("string with only spaces and zero-width chars", func(t *testing.T) {
		input := "   \u200B  \uFEFF  \u2060  "
		expected := ""

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("alternating characters and spaces", func(t *testing.T) {
		input := " a  b  c  d  e "
		expected := "a b c d e"

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("mixed zero-width and whitespace", func(t *testing.T) {
		input := "text\u200B \t\uFEFF\n\u2060 \rmore"
		expected := "text more"

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("large json with formatting", func(t *testing.T) {
		input := `{
			"users"  :  [
				{  "id"  :  1  ,  "name"  :  "John"  },
				{  "id"  :  2  ,  "name"  :  "Jane"  }
			]
		}`
		expected := `{ "users" : [ { "id" : 1 , "name" : "John" }, { "id" : 2 , "name" : "Jane" } ] }`

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("xml with multiple namespaces", func(t *testing.T) {
		input := `<root   xmlns:a="ns1"   xmlns:b="ns2"  >
			<a:item   id="1"  >  content  </a:item>
			<b:item   id="2"  >  content  </b:item>
		</root>`
		expected := `<root xmlns:a="ns1" xmlns:b="ns2" > <a:item id="1" > content </a:item> <b:item id="2" > content </b:item> </root>`

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("html with inline styles and scripts", func(t *testing.T) {
		input := `<div   style="  margin:  10px;  padding:  5px;  "  >
			<script>  var   x  =  "hello   world";  </script>
		</div>`
		expected := `<div style=" margin: 10px; padding: 5px; " > <script> var x = "hello world"; </script> </div>`

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})

	t.Run("mixed formats with zero-width chars", func(t *testing.T) {
		input := "{\"html\"\u200B:\u200B\"<div\u200Bclass\u200B=\u200B'test'\u200B>content</div>\"}"
		expected := "{\"html\":\"<divclass='test'>content</div>\"}"

		result := RemoveDuplicateSpaces(input)
		assert.Equal(t, expected, result)
	})
}

// Benchmark tests
func BenchmarkRemoveDuplicateSpaces_ShortString(b *testing.B) {
	input := "hello    world"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}

func BenchmarkRemoveDuplicateSpaces_LongString(b *testing.B) {
	input := "start" + strings.Repeat(" ", 1000) + "middle" + strings.Repeat("\t", 500) + "end"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}

func BenchmarkRemoveDuplicateSpaces_ZeroWidthChars(b *testing.B) {
	input := "hello\u200B\uFEFF\u2060world\u200D\u200E\u200F\u00A0test"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}

func BenchmarkRemoveDuplicateSpaces_JSON(b *testing.B) {
	input := `{  "users"  :  [  {  "id"  :  1  ,  "name"  :  "John"  }  ,  {  "id"  :  2  ,  "name"  :  "Jane"  }  ]  }`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}

func BenchmarkRemoveDuplicateSpaces_HTML(b *testing.B) {
	input := `<div   class="container"  >  <h1>  Title  </h1>  <p>  Content  with  spaces  </p>  </div>`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}

func BenchmarkRemoveDuplicateSpaces_XML(b *testing.B) {
	input := `<root   xmlns="namespace"  >  <item   id="1"  >  <data>  value  </data>  </item>  </root>`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicateSpaces(input)
	}
}
