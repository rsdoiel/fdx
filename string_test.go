// Copyright (C) 2017-2026 R. S. Doiel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package fdx

import (
	"errors"
	"io/fs"
	"strings"
	"testing"

	"github.com/rsdoiel/fountain"
)

// These tests pin what the code does to input it was written for. They do not
// say Final Draft agrees: no Final Draft is available to check against.

/** TestTextString covers every style a Text can carry, the order the markers
 * nest in, upper casing by AllCaps or a Capitals font, and a nil Text.
 */
func TestTextString(t *testing.T) {
	cases := []struct {
		name string
		text *Text
		want string
	}{
		{"nil", nil, ""},
		{"plain", &Text{InnerText: "a"}, "a"},
		{"italic", &Text{Style: ItalicStyle, InnerText: "a"}, "*a*"},
		{"bold", &Text{Style: BoldStyle, InnerText: "a"}, "**a**"},
		{"underline", &Text{Style: UnderlineStyle, InnerText: "a"}, "_a_"},
		{"strikethrough", &Text{Style: Strikethrough, InnerText: "a"}, "~~a~~"},
		{"all caps", &Text{Style: AllCapsStyle, InnerText: "ab"}, "AB"},
		{"capitals font", &Text{Font: "Courier Final Draft+Capitals", InnerText: "ab"}, "AB"},
		{"bold italic", &Text{Style: "Bold+Italic", InnerText: "a"}, "***a***"},
		{"all styles", &Text{Style: "AllCaps+Italic+Bold+Underline+Strikethrough", InnerText: "a"}, "~~_***A***_~~"},
		{"numeric style is plain", &Text{Style: "0", InnerText: "a"}, "a"},
	}
	for _, c := range cases {
		if got := c.text.String(); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

/** TestParagraphString covers the rendering rule of each paragraph Type: which
 * ones are upper cased, which get a trailing blank line, parentheses on a
 * parenthetical, the ~ of singing, centering, the page break and a scene number.
 */
func TestParagraphString(t *testing.T) {
	para := func(typ, text string) *Paragraph {
		return &Paragraph{Type: typ, Text: []*Text{{InnerText: text}}}
	}
	centered := para(ActionType, "THE END")
	centered.Alignment = CenterAlignment
	emptyCentered := para(ActionType, "")
	emptyCentered.Alignment = CenterAlignment
	numbered := para(SceneHeadingType, "int. kitchen")
	numbered.Number = "12A"
	paged := para(ActionType, "Later.")
	paged.StartsNewPage = "Yes"
	cases := []struct {
		name string
		p    *Paragraph
		want string
	}{
		{"nil", nil, ""},
		{"scene heading", para(SceneHeadingType, "int. kitchen - day"), "INT. KITCHEN - DAY\n\n"},
		{"scene heading with number", numbered, "INT. KITCHEN #12A#\n\n"},
		{"action", para(ActionType, "Anna waits."), "Anna waits.\n\n"},
		{"character", para(CharacterType, "anna"), "ANNA\n"},
		{"parenthetical", para(ParentheticalType, "softly"), "(softly)\n"},
		{"parenthetical already wrapped", para(ParentheticalType, "(softly)"), "(softly)\n"},
		{"dialogue", para(DialogueType, "Hello."), "Hello.\n\n"},
		{"transition", para(TransitionType, "cut to:"), "CUT TO:\n\n"},
		{"singing", para(SingingType, "la la"), "~la la\n"},
		{"general", para(GeneralType, "x"), "x\n"},
		{"centered", centered, ">THE END<\n\n"},
		{"centered but empty", emptyCentered, "\n\n"},
		{"starts a new page", paged, "===\n\nLater.\n\n"},
	}
	for _, c := range cases {
		if got := c.p.String(); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}

/** TestStringOfNilAndEmptyParts checks that a missing part renders as nothing
 * and that a document with only one part renders that part.
 */
func TestStringOfNilAndEmptyParts(t *testing.T) {
	var nilDoc *FinalDraft
	var nilContent *Content
	var nilTitle *TitlePage
	for name, got := range map[string]string{
		"nil document":          nilDoc.String(),
		"nil content":           nilContent.String(),
		"empty content":         (&Content{}).String(),
		"nil title page":        nilTitle.String(),
		"title page no content": (&TitlePage{}).String(),
		"empty document":        NewFinalDraft().String(),
	} {
		if got != "" {
			t.Errorf("%s: got %q, want \"\"", name, got)
		}
	}
	body := &Content{Paragraph: []*Paragraph{{Type: ActionType, Text: []*Text{{InnerText: "Hi."}}}}}
	if got := (&FinalDraft{Content: body}).String(); got != "Hi.\n\n" {
		t.Errorf("content only: got %q", got)
	}
	if got := (&FinalDraft{TitlePage: &TitlePage{Content: body}}).String(); got != "Hi.\n\n\n" {
		t.Errorf("title page only: got %q", got)
	}
}

/** TestParseAndParseFile covers reading: a good file, a file that does not
 * exist (the error says so and the document is nil), and bytes that are not XML.
 */
func TestParseAndParseFile(t *testing.T) {
	doc, err := ParseFile("testdata/sample-01.fdx")
	if err != nil {
		t.Fatal(err)
	}
	if doc.DocumentType != "Script" || doc.Content == nil || len(doc.Content.Paragraph) == 0 {
		t.Errorf("sample-01.fdx: type %q, content %v", doc.DocumentType, doc.Content)
	}
	doc, err = ParseFile("testdata/no-such-file.fdx")
	if !errors.Is(err, fs.ErrNotExist) || doc != nil {
		t.Errorf("missing file: got %v, %v; want nil and fs.ErrNotExist", doc, err)
	}
	if _, err := Parse([]byte("<FinalDraft><Content>")); err == nil {
		t.Error("unterminated XML: want an error")
	}
	if _, err := Parse([]byte("plain text, not XML")); err == nil {
		t.Error("plain text: want an error")
	}
}

/** TestToXMLRoundTrip renders a document made from Fountain and reads it back:
 * the XML starts with the root element, has the empty elements self-closing,
 * and the paragraphs survive.
 */
func TestToXMLRoundTrip(t *testing.T) {
	screenplay, err := fountain.Parse([]byte("INT. KITCHEN - DAY\n\nAnna waits.\n"))
	if err != nil {
		t.Fatal(err)
	}
	doc := NewFinalDraft()
	doc.FromFountain(screenplay)
	src, err := doc.ToXML()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(string(src), "<FinalDraft") {
		t.Errorf("XML starts %.30q", src)
	}
	back, err := Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	if back.String() != doc.String() {
		t.Errorf("round trip changed the text:\n%q\n%q", doc.String(), back.String())
	}
}

/** TestFromFountainTitlePage checks that each title page field the converter
 * knows (title, credit, author, copyright, draft date, contact, source, story
 * by) becomes a title page paragraph, in order, whatever the case of its name,
 * and that a script with no title page leaves the title page nil.
 */
func TestFromFountainTitlePage(t *testing.T) {
	src := "Title: The Kitchen\nCredit: Written by\nAuthor: A. Writer\nCopyright: (c) 2026\nDRAFT DATE: 2026-11-18\nContact: a@example.org\nSource: A story\nStory by: B. Teller\n\nINT. KITCHEN - DAY\n\nAnna waits.\n"
	screenplay, err := fountain.Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	doc := NewFinalDraft()
	doc.FromFountain(screenplay)
	want := []string{"The Kitchen", "Written by", "A. Writer", "(c) 2026", "2026-11-18", "a@example.org", "A story", "B. Teller"}
	if doc.TitlePage == nil || doc.TitlePage.Content == nil {
		t.Fatal("no title page")
	}
	var got []string
	for _, p := range doc.TitlePage.Content.Paragraph {
		got = append(got, strings.TrimSpace(p.Text[0].InnerText))
	}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("title page %q, want %q", got, want)
	}

	plain, err := fountain.Parse([]byte("INT. KITCHEN - DAY\n\nAnna waits.\n"))
	if err != nil {
		t.Fatal(err)
	}
	doc = NewFinalDraft()
	doc.FromFountain(plain)
	if doc.TitlePage != nil {
		t.Errorf("no title page in the script, got %v", doc.TitlePage)
	}
}

/** TestFromFountainKeepsExistingParts checks that FromFountain adds to a
 * document that already has content rather than replacing it.
 */
func TestFromFountainKeepsExistingParts(t *testing.T) {
	doc := NewFinalDraft()
	doc.Content = &Content{Paragraph: []*Paragraph{{Type: ActionType, Text: StringToTextArray("First.")}}}
	screenplay, err := fountain.Parse([]byte("INT. A - DAY\n\nSecond.\n"))
	if err != nil {
		t.Fatal(err)
	}
	doc.FromFountain(screenplay)
	if len(doc.Content.Paragraph) < 3 || doc.Content.Paragraph[0].Text[0].InnerText != "First." {
		t.Errorf("existing paragraph lost: %d paragraphs", len(doc.Content.Paragraph))
	}
}

/** TestFmtHelp checks the four tokens and that text without tokens is unchanged. */
func TestFmtHelp(t *testing.T) {
	got := FmtHelp("{app_name} {version} {release_hash} {release_date} {app_name}", "x", "1.2.3", "2026-01-02", "abc1234")
	if got != "x 1.2.3 abc1234 2026-01-02 x" {
		t.Errorf("got %q", got)
	}
	if got := FmtHelp("no tokens", "x", "1", "d", "h"); got != "no tokens" {
		t.Errorf("got %q", got)
	}
}
