package paragraph

import (
	"testing"

	"github.com/christiandenisi/editorjs-go"
)

func TestRenderParagraph(t *testing.T) {
	json := []byte(`{
		"time": 1752781597903,
		"blocks": [
			{
				"id": "abc123",
				"type": "paragraph",
				"data": {
					"text": "Hello <script>alert(1)</script>"
				}
			}
		],
		"version": "2.27.0"
	}`)

	converter := editorjs.New()
	editorjs.RegisterTyped(converter, "paragraph", RenderParagraph)

	output, err := converter.Convert(json)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<p>Hello &lt;script&gt;alert(1)&lt;/script&gt;</p>"
	if output != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
