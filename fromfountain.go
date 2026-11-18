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
	"strings"

	// My Packages
	"github.com/rsdoiel/fountain"
)

/** StringToTextArray returns the string as a one element slice of Text. Embedded formatting is not split into styled runs yet.
 *
 * Parameters:
 *   s (string) — the text of a paragraph
 *
 * Returns:
 *   []*Text — one Text whose InnerText is s
 *
 * Example:
 *   paragraph.Text = fdx.StringToTextArray("Anna waits.")
 */
func StringToTextArray(s string) []*Text {
	var a []*Text

	//FIXME: if string has embedded formatting split into parts and set format attributes
	// (e.g. text.Bold, text.Italic) as needed.
	text := new(Text)
	text.InnerText = s
	a = append(a, text)
	return a
}

/** FromFountain fills the document from a parsed Fountain screenplay: the title page fields become title page paragraphs, and every element becomes a Paragraph of the element's type. A scene heading's scene number (#12A#) is moved from the text into the Paragraph's Number attribute.
 *
 * Parameters:
 *   screenplay (*fountain.Fountain) — a parsed Fountain screenplay
 *
 * Example:
 *   screenplay, _ := fountain.Parse(src)
 *   document := fdx.NewFinalDraft()
 *   document.FromFountain(screenplay)
 */
func (document *FinalDraft) FromFountain(screenplay *fountain.Fountain) {
	if screenplay.TitlePage != nil {
		// Build the Info section
		if document.TitlePage == nil {
			document.TitlePage = new(TitlePage)
		}
		if document.TitlePage.Content == nil {
			document.TitlePage.Content = new(Content)
		}

		// NOTE: build a map of elements that will belong to Info.
		for _, elem := range screenplay.TitlePage {
			switch strings.ToLower(elem.Name) {
			case "title":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "credit":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "author":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "copyright":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "draft date":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "contact":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "source":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "story by":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			}
		}
	}

	if screenplay.Elements != nil {
		// Populate the Paragraphs array
		if document.Content == nil {
			document.Content = new(Content)
		}
		for _, elem := range screenplay.Elements {
			text := new(Text)
			text.InnerText = elem.Content
			paragraph := new(Paragraph)
			paragraph.Type = elem.TypeName()
			if elem.Type == fountain.SceneHeadingType && elem.SceneNumber != "" {
				// Final Draft keeps a scene number in the Number attribute,
				// not in the heading's text.
				paragraph.Number = elem.SceneNumber
				text.InnerText = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(elem.Content), "#"+elem.SceneNumber+"#"))
			}
			paragraph.Text = append(paragraph.Text, text)
			document.Content.Paragraph = append(document.Content.Paragraph, paragraph)
		}
	}
}
