package mjml2html_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/TcMits/mjml2html"
	"golang.org/x/net/html"
)

func Test_ToHTML(t *testing.T) {
	input := `<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-image width="100px" src="/assets/img/logo-small.png"></mj-image>
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">{{ if eq 1 2 }}3{{ end }}</mj-text>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">Hello World</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`

	result, err := mjml2html.ToHTML(input)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := html.Parse(strings.NewReader(result)); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(result, "{{ if eq 1 2 }}") {
		t.Fatal("Expected to contain '{{ if eq 1 2 }}'")
	}

	if !strings.Contains(result, "{{ end }}") {
		t.Fatal("Expected to contain '{{ end }}'")
	}
}

func Test_ToHTML_JSON(t *testing.T) {
	input := `{"tagName":"mjml","attributes":{},"children":[{"tagName":"mj-body","attributes":{},"children":[{"tagName":"mj-section","attributes":{},"children":[{"tagName":"mj-column","attributes":{},"children":[{"tagName":"mj-image","attributes":{"width":"100px","src":"/assets/img/logo-small.png"}},{"tagName":"mj-divider","attributes":{"border-color":"#F46E43"}},{"tagName":"mj-text","attributes":{"font-size":"20px","color":"#F45E43","font-family":"Helvetica"},"content":"{{ if eq 1 2 }}3{{ end }}"},{"tagName":"mj-text","attributes":{"font-size":"20px","color":"#F45E43","font-family":"Helvetica"},"content":"Hello World"}]}]}]}]}`

	node := mjml2html.MJMLNode{}
	if err := json.Unmarshal([]byte(input), &node); err != nil {
		t.Fatal(err)
	}

	result, err := mjml2html.ToHTML(&node)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := html.Parse(strings.NewReader(result)); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(result, "{{ if eq 1 2 }}") {
		t.Fatal("Expected to contain '{{ if eq 1 2 }}'")
	}

	if !strings.Contains(result, "{{ end }}") {
		t.Fatal("Expected to contain '{{ end }}'")
	}
}

func Benchmark_ToHTML(b *testing.B) {
	input := `<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-image width="100px" src="/assets/img/logo-small.png"></mj-image>
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">Hello World</mj-text>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">{{ if eq 1 2 }}3{{ end }}</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`

	for i := 0; i < b.N; i++ {
		_, err := mjml2html.ToHTML(input)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_ToHTML_Parallel(b *testing.B) {
	input := `<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-image width="100px" src="/assets/img/logo-small.png"></mj-image>
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">Hello World</mj-text>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">{{ if eq 1 2 }}3{{ end }}</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := mjml2html.ToHTML(input)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Test_ToJSON(t *testing.T) {
	input := `<mjml>
  <mj-body>
    <mj-section>
      <mj-column>
        <mj-image src="/assets/img/logo-small.png" width="100px" ></mj-image>
        <mj-divider border-color="#F45E43"></mj-divider>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">{{ if eq 1 2 }}3{{ end }}</mj-text>
        <mj-text font-size="20px" color="#F45E43" font-family="helvetica">Hello World</mj-text>
      </mj-column>
    </mj-section>
  </mj-body>
</mjml>`
	expect := `{"tagName":"mjml","attributes":{},"children":[{"tagName":"mj-body","attributes":{},"children":[{"tagName":"mj-section","attributes":{},"children":[{"tagName":"mj-column","attributes":{},"children":[{"tagName":"mj-image","attributes":{"src":"/assets/img/logo-small.png","width":"100px"}},{"tagName":"mj-divider","attributes":{"border-color":"#F45E43"}},{"tagName":"mj-text","attributes":{"color":"#F45E43","font-family":"helvetica","font-size":"20px"},"content":"{{ if eq 1 2 }}3{{ end }}"},{"tagName":"mj-text","attributes":{"color":"#F45E43","font-family":"helvetica","font-size":"20px"},"content":"Hello World"}]}]}]}]}`

	result, err := mjml2html.ToJSON(input)
	if err != nil {
		t.Fatal(err)
	}

	resultB, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}

	if string(resultB) != expect {
		t.Fatal(string(resultB))
	}
}
