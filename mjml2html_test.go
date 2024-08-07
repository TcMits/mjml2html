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
    <mj-body background-color="#F4F4F4" color="#55575d" font-family="Arial, sans-serif">
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top">
            <mj-column>
                <mj-image align="center" border="none" padding-bottom="30px" padding="10px 25px" src="http://5vph.mj.am/img/5vph/b/1g86w/0g67t.png" target="_blank" title="" width="180px"></mj-image>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="0px" padding-top="0px" padding="10px 25px">
                    <p style="line-height: 18px; margin: 10px 0; text-align: center;font-size:14px;color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">WOMEN&nbsp; &nbsp; &nbsp; &nbsp;| &nbsp; &nbsp; &nbsp; MEN&nbsp; &nbsp; &nbsp; &nbsp;| &nbsp; &nbsp; &nbsp; KIDS</p>
                </mj-text>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding="0 0 0 0">
            <mj-column>
                <mj-image align="center" border="none" padding-bottom="0px" padding-left="0px" padding-right="0px" padding="0px 25px" src="http://5vph.mj.am/img/5vph/b/1g86w/0696u.jpeg" target="_blank" title="" width="780px"></mj-image>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding="0 0 0 0">
            <mj-column>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="5px" padding-top="25px" padding="10px 25px">
                    <p style="line-height: 60px; text-align: center; margin: 10px 0;font-size:55px;color:#fcfcfc;font-family:'Times New Roman',Helvetica,Arial,sans-serif"><b>Black Friday</b></p>
                </mj-text>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="20px" padding-top="0px" padding="10px 25px">
                    <p style="line-height: 30px; text-align: center; margin: 10px 0;color:#f5f5f5;font-size:25px;font-family:'Times New Roman',Helvetica,Arial,sans-serif"><b>Take an&nbsp; extra 50% off</b><br><span style="color:#ffffff;font-size:18px;font-family:'Times New Roman',Helvetica,Arial,sans-serif">Use code SALEONSALE* at checkout</span></p>
                </mj-text>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding-bottom="40px" padding="0 0 0 0">
            <mj-column>
                <mj-button background-color="#ffffff" border-radius="3px" font-family="Times New Roman, Helvetica, Arial, sans-serif" font-size="18px" font-weight="normal" inner-padding="10px 25px" padding-bottom="30px" padding="10px 25px"><span style="color:#212020">Shop Now</span></mj-button>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="0px" padding-top="5px" padding="10px 25px">
                    <p style="line-height: 16px; text-align: center; margin: 10px 0;font-size:12px;color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">* Offer valid on Allura purchases on 17/29/11 at 11:59 pm. No price adjustments on previous&nbsp;<br><span style="color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">purchases, offer limited to stock. Cannot be combined with any offer or promotion other than free.</span></p>
                </mj-text>
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
    <mj-body background-color="#F4F4F4" color="#55575d" font-family="Arial, sans-serif">
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top">
            <mj-column>
                <mj-image align="center" border="none" padding-bottom="30px" padding="10px 25px" src="http://5vph.mj.am/img/5vph/b/1g86w/0g67t.png" target="_blank" title="" width="180px"></mj-image>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="0px" padding-top="0px" padding="10px 25px">
                    <p style="line-height: 18px; margin: 10px 0; text-align: center;font-size:14px;color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">WOMEN&nbsp; &nbsp; &nbsp; &nbsp;| &nbsp; &nbsp; &nbsp; MEN&nbsp; &nbsp; &nbsp; &nbsp;| &nbsp; &nbsp; &nbsp; KIDS</p>
                </mj-text>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding="0 0 0 0">
            <mj-column>
                <mj-image align="center" border="none" padding-bottom="0px" padding-left="0px" padding-right="0px" padding="0px 25px" src="http://5vph.mj.am/img/5vph/b/1g86w/0696u.jpeg" target="_blank" title="" width="780px"></mj-image>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding="0 0 0 0">
            <mj-column>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="5px" padding-top="25px" padding="10px 25px">
                    <p style="line-height: 60px; text-align: center; margin: 10px 0;font-size:55px;color:#fcfcfc;font-family:'Times New Roman',Helvetica,Arial,sans-serif"><b>Black Friday</b></p>
                </mj-text>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="20px" padding-top="0px" padding="10px 25px">
                    <p style="line-height: 30px; text-align: center; margin: 10px 0;color:#f5f5f5;font-size:25px;font-family:'Times New Roman',Helvetica,Arial,sans-serif"><b>Take an&nbsp; extra 50% off</b><br><span style="color:#ffffff;font-size:18px;font-family:'Times New Roman',Helvetica,Arial,sans-serif">Use code SALEONSALE* at checkout</span></p>
                </mj-text>
            </mj-column>
        </mj-section>
        <mj-section background-color="#000000" background-repeat="no-repeat" text-align="center" vertical-align="top" padding-bottom="40px" padding="0 0 0 0">
            <mj-column>
                <mj-button background-color="#ffffff" border-radius="3px" font-family="Times New Roman, Helvetica, Arial, sans-serif" font-size="18px" font-weight="normal" inner-padding="10px 25px" padding-bottom="30px" padding="10px 25px"><span style="color:#212020">Shop Now</span></mj-button>
                <mj-text align="left" color="#55575d" font-family="Arial, sans-serif" font-size="13px" line-height="22px" padding-bottom="0px" padding-top="5px" padding="10px 25px">
                    <p style="line-height: 16px; text-align: center; margin: 10px 0;font-size:12px;color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">* Offer valid on Allura purchases on 17/29/11 at 11:59 pm. No price adjustments on previous&nbsp;<br><span style="color:#ffffff;font-family:'Times New Roman',Helvetica,Arial,sans-serif">purchases, offer limited to stock. Cannot be combined with any offer or promotion other than free.</span></p>
                </mj-text>
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
