package paragraph

import (
	"github.com/christiandenisi/editorjs-go"
)

// ParagraphData represents the structure of a paragraph block in Editor.js.
type ParagraphData struct {
	Text string `json:"text"`
}

// RenderParagraph renders a paragraph block to HTML.
func RenderParagraph(b editorjs.Block[ParagraphData], ctx *editorjs.Context) (string, error) {
	return "<p>" + ctx.EscapeHTML(b.Data.Text) + "</p>", nil
}
