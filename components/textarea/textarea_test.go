package textarea_test

import (
	"strings"
	"testing"
	g "maragu.dev/gomponents"
	"github.com/rizome-dev/shadcn-gomponents/components/textarea"
)

func TestTextarea(t *testing.T) {
	tests := []struct {
		name     string
		textarea g.Node
		contains []string
	}{
		{
			name:     "default textarea",
			textarea: textarea.Default(),
			contains: []string{
				`<textarea`,
				`class="flex min-h-[80px] w-full rounded-md border`,
				`border-input bg-background`,
				`px-3 py-2`,
				`text-sm`,
				`focus-visible:outline-none focus-visible:ring-2`,
			},
		},
		{
			name:     "textarea with placeholder",
			textarea: textarea.WithPlaceholder("Enter your message..."),
			contains: []string{
				`placeholder="Enter your message..."`,
			},
		},
		{
			name: "textarea with all props",
			textarea: textarea.New(textarea.Props{
				ID:          "message",
				Name:        "message",
				Value:       "Initial text",
				Placeholder: "Type here...",
				Rows:        5,
				Cols:        40,
				MaxLength:   500,
				MinLength:   10,
				Disabled:    true,
				Required:    true,
				ReadOnly:    true,
				OnChange:    "handleChange()",
			}),
			contains: []string{
				`id="message"`,
				`name="message"`,
				`>Initial text</textarea>`,
				`placeholder="Type here..."`,
				`rows="5"`,
				`cols="40"`,
				`maxlength="500"`,
				`minlength="10"`,
				`disabled`,
				`required`,
				`readonly`,
				`onchange="handleChange()"`,
			},
		},
		{
			name:     "auto-resize textarea",
			textarea: textarea.AutoResize(),
			contains: []string{
				`field-sizing-content`,
				`resize-none`,
			},
		},
		{
			name:     "no-resize textarea",
			textarea: textarea.NoResize(),
			contains: []string{
				`resize-none`,
			},
		},
		{
			name: "textarea with different resize options",
			textarea: textarea.New(textarea.Props{
				Resize: "both",
			}),
			contains: []string{
				`resize`,
			},
		},
		{
			name: "form field textarea",
			textarea: textarea.FormField(
				textarea.Props{
					Name:     "description",
					Required: true,
				},
				"Description",
				"Please provide a detailed description",
			),
			contains: []string{
				`<label for="textarea-description"`,
				`Description`,
				`<span class="text-destructive ml-1">*</span>`,
				`<p class="text-sm text-muted-foreground">Please provide a detailed description</p>`,
			},
		},
		{
			name: "textarea with character count",
			textarea: textarea.WithCharacterCount(
				textarea.Props{
					Name:      "bio",
					MaxLength: 200,
				},
				"Bio",
			),
			contains: []string{
				`<label for="textarea-count-bio"`,
				`Bio</label>`,
				`<span id="textarea-count-bio-count">0</span>/200`,
				`oninput="document.getElementById('textarea-count-bio-count').textContent = this.value.length"`,
			},
		},
		{
			name: "message textarea",
			textarea: textarea.Message("comment", 4),
			contains: []string{
				`name="comment"`,
				`placeholder="Type your message here..."`,
				`rows="4"`,
				`resize-none`,
			},
		},
		{
			name: "bio textarea",
			textarea: textarea.Bio("user-bio", 500),
			contains: []string{
				`name="user-bio"`,
				`placeholder="Tell us a little bit about yourself"`,
				`rows="4"`,
				`maxlength="500"`,
				`field-sizing-content`, // auto-resize
				`>Bio</label>`,
				`<span id="textarea-count-user-bio-count">0</span>/500`,
			},
		},
		{
			name: "textarea with custom class",
			textarea: textarea.New(textarea.Props{
				Class: "custom-textarea",
			}),
			contains: []string{
				`custom-textarea`,
			},
		},
		{
			name: "textarea with onInput handler",
			textarea: textarea.New(textarea.Props{
				OnInput: "updatePreview(this.value)",
			}),
			contains: []string{
				`oninput="updatePreview(this.value)"`,
			},
		},
		{
			name: "horizontal resize textarea",
			textarea: textarea.New(textarea.Props{
				Resize: "horizontal",
			}),
			contains: []string{
				`resize-x`,
			},
		},
		{
			name: "vertical resize textarea",
			textarea: textarea.New(textarea.Props{
				Resize: "vertical",
			}),
			contains: []string{
				`resize-y`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.textarea.String()
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("expected result to contain %q, but it didn't.\nGot: %s", expected, result)
				}
			}
		})
	}
}