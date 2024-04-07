var mjml2html = require('mjml-browser')

export default function toHTML(mjml, opts) {
  return mjml2html(mjml, opts);
};
