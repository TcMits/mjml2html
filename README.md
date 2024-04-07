# mjml2html

## Usage

```go
package main

import "github.com/TcMits/mjml2html"

func main() {
	html, err := mjml2html.ToHTML(`<mjml>
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
</mjml>`)
	if err != nil {
		panic(err)
	}

	println(html)
}
```
