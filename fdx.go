// Package fdx supports encoding/decoding fdx (Final Draft XML) formatted screenplays.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
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
	"bytes"
	"encoding/xml"
	"os"
	"strings"
)

/** The constants name the strings Final Draft writes: DocString is the XML
 * declaration, the *Style, *Alignment and *Type constants are the values of the
 * Style, Alignment and Type attributes of Text, Paragraph, ElementSettings and
 * DynamicLabel.
 *
 * Example:
 *   if paragraph.Type == fdx.SceneHeadingType { ... }
 */
const (
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`

	// Style
	UnderlineStyle = "Underline"
	ItalicStyle    = "Italic"
	BoldStyle      = "Bold"
	AllCapsStyle   = "AllCaps"
	Strikethrough  = "Strikethrough"

	// Alignments
	CenterAlignment = "Center"
	LeftAlignment   = "Left"
	RightAlignment  = "Right"

	// Types used in ElementSettings and Paragraph elements
	GeneralType       = "General"
	SceneHeadingType  = "Scene Heading"
	ActionType        = "Action"
	CharacterType     = "Character"
	DialogueType      = "Dialogue"
	ParentheticalType = "Parenthetical"
	TransitionType    = "Transition"
	CastListType      = "Cast List"
	ShotType          = "Shot"
	SingingType       = "Singing"

	// DynamicLabel types
	PageNoType      = "Page #"
	LastRevisedType = "Last Revised"

	// Tabstop types
	RightType = "Right"
	LeftType  = "Left"
)

/** MaxLineWidth is the width in characters of a line, in a monospace font.
 *
 * Example:
 *   fdx.MaxLineWidth = 60
 */
var (
	// MaxLineWidth is the number of characters wide a line can be
	// based on a monospace font.
	MaxLineWidth = 80
)

/** FinalDraft is the root of an fdx document: the XML element <FinalDraft> with its screenplay Content, TitlePage and the document wide settings Final Draft keeps beside them.
 *
 * Example:
 *   doc, _ := fdx.ParseFile("screenplay.fdx")
 *   fmt.Println(doc.Version, len(doc.Content.Paragraph))
 */
type FinalDraft struct {
	XMLName               xml.Name `json:"-" yaml:"-"`
	DocumentType          string   `xml:",attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Template              string   `xml:",attr" json:"template,omitempty" yaml:"template,omitempty"`
	Version               string   `xml:",attr" json:"version,omitempty" yaml:"version,omitempty"`
	Content               *Content
	TitlePage             *TitlePage
	ElementSettings       []*ElementSettings
	HeaderAndFooter       *HeaderAndFooter
	SpellCheckIgnoreLists *SpellCheckIgnoreLists
	PageLayout            *PageLayout
	WindowState           *WindowState
	TextState             *TextState
	ScriptNoteDefinitions *ScriptNoteDefinitions
	SmartType             *SmartType
	MoresAndContinueds    *MoresAndContinueds
	LockedPages           *LockedPages
	Revisions             *Revisions
	SplitState            *SplitState
	Macros                *Macros
	Actors                *Actors
	Cast                  *Cast `xml:"Cast,omitempty" json:"cast,omitempty" yaml:"cast,omitempty"`
	SceneNumberOptions    *SceneNumberOptions
}

/** Content is the <Content> element: the ordered paragraphs of the screenplay body, or of a title page.
 *
 * Example:
 *   n := len(doc.Content.Paragraph)
 */
type Content struct {
	XMLName   xml.Name     `json:"-" yaml:"-"`
	Paragraph []*Paragraph `json:"paragraphs,omitempty" yaml:"paragraphs,omitempty"`
}

/** Paragraph is one <Paragraph> of a screenplay. Type says what it is (see SceneHeadingType, ActionType and the other *Type constants); Number is a scene number; the remaining attributes are layout overrides. Its words are in Text.
 *
 * Example:
 *   p := doc.Content.Paragraph[0]
 *   fmt.Println(p.Type, p.Text[0].InnerText)
 */
type Paragraph struct {
	XMLName         xml.Name `json:"-" yaml:"-"`
	Type            string   `xml:",attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Number          string   `xml:",attr,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	Alignment       string   `xml:",attr,omitempty" json:"alignment,omitempty" yaml:"alignment,omitempty"`
	FirstIndent     string   `xml:",attr,omitempty" json:"first_indent,omitempty" yaml:"first_indent,omitempty"`
	Leading         string   `xml:",attr,omitempty" json:"leading,omitempty" yaml:"leading,omitempty"`
	LeftIndent      string   `xml:",attr,omitempty" json:"left_indent,omitempty" yaml:"left_indent,omitempty"`
	RightIndent     string   `xml:",attr,omitempty" json:"right_indent,omitempty" yaml:"right_indent,omitempty"`
	SpaceBefore     string   `xml:",attr,omitempty" json:"space_before,omitempty" yaml:"space_before,omitempty"`
	Spacing         string   `xml:",attr,omitempty" json:"spacing,omitempty" yaml:"spacing,omitempty"`
	StartsNewPage   string   `xml:",attr,omitempty" json:"starts_new_page,omitempty" yaml:"starts_new_page,omitempty"`
	SceneProperties []*SceneProperties
	DynamicLabel    []*DynamicLabel
	Text            []*Text
}

/** SceneProperties is the <SceneProperties> element of a paragraph: the length, page and title Final Draft records for a scene.
 *
 * Example:
 *   sp := &fdx.SceneProperties{Title: "Kitchen"}
 */
type SceneProperties struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Length  string   `xml:",attr,omitempty" json:"length,omitempty" yaml:"length,omitempty"`
	Page    string   `xml:",attr,omitempty" json:"page,omitempty" yaml:"page,omitempty"`
	Title   string   `xml:",attr,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

/** HeaderAndFooter is the <HeaderAndFooter> element: whether the page header and footer show, on the first page too, the starting page number, and the Header and Footer content.
 *
 * Example:
 *   visible := doc.HeaderAndFooter.HeaderVisible == "Yes"
 */
type HeaderAndFooter struct {
	XMLName         xml.Name `json:"-" yaml:"-"`
	FooterFirstPage string   `xml:",attr,omitempty" json:"footer_first_page,omitempty" yaml:"footer_first_page,omitempty"`
	FooterVisible   string   `xml:",attr,omitempty" json:"footer_visible,omitempty" yaml:"footer_visible,omitempty"`
	HeaderFirstPage string   `xml:",attr,omitempty" json:"header_first_page,omitempty" yaml:"header_first_page,omitempty"`
	HeaderVisible   string   `xml:",attr,omitempty" json:"header_visible,omitempty" yaml:"header_visible,omitempty"`
	StartingPage    string   `xml:",attr,omitempty" json:"starting_page,omitempty" yaml:"starting_page,omitempty"`
	Header          Header
	Footer          Footer
}

/** Header is the <Header> element of HeaderAndFooter: the paragraphs printed at the top of every page.
 *
 * Example:
 *   n := len(doc.HeaderAndFooter.Header.Paragraph)
 */
type Header struct {
	XMLName   xml.Name    `json:"-" yaml:"-"`
	Paragraph []Paragraph `json:"paragraphs,omitempty" yaml:"paragraphs,omitempty"`
}

/** DynamicLabel is a <DynamicLabel> in a header or footer: text Final Draft fills in when it prints, such as the page number (PageNoType) or the date last revised (LastRevisedType).
 *
 * Example:
 *   label := &fdx.DynamicLabel{Type: fdx.PageNoType}
 */
type DynamicLabel struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Type    string   `xml:",attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

/** Footer is the <Footer> element of HeaderAndFooter: the paragraphs printed at the bottom of every page.
 *
 * Example:
 *   n := len(doc.HeaderAndFooter.Footer.Paragraph)
 */
type Footer struct {
	XMLName   xml.Name    `json:"-" yaml:"-"`
	Paragraph []Paragraph `json:"paragraphs,omitempty" yaml:"paragraphs,omitempty"`
}

/** Text is a <Text> run in a paragraph: InnerText with the character styling that applies to all of it (Style holds style names such as Bold, Italic, Underline, AllCaps and Strikethrough).
 *
 * Example:
 *   t := &fdx.Text{Style: fdx.ItalicStyle, InnerText: "softly"}
 *   fmt.Println(t.String()) // *softly*
 */
type Text struct {
	XMLName        xml.Name `json:"-" yaml:"-"`
	AdornmentStyle string   `xml:",attr,omitempty" json:"adornment_type,omitempty" yaml:"adornment_type,omitempty"`
	Background     string   `xml:",attr,omitempty" json:"background,omitempty" yaml:"background,omitempty"`
	Color          string   `xml:",attr,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	Font           string   `xml:",attr,omitempty" json:"font,omitempty" yaml:"font,omitempty"`
	RevisionID     string   `xml:",attr,omitempty" json:"revision_id,omitempty" yaml:"revision_id,omitempty"`
	Size           string   `xml:",attr,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	Style          string   `xml:",attr,omitempty" json:"style,omitempty" yaml:"style,omitempty"`
	InnerText      string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** TitlePage is the <TitlePage> element: the Content of the title page and the header and footer used on it.
 *
 * Example:
 *   fmt.Print(doc.TitlePage.String())
 */
type TitlePage struct {
	XMLName         xml.Name `json:"-" yaml:"-"`
	HeaderAndFooter *HeaderAndFooter
	Content         *Content
}

/** Revisions is the <Revisions> element: the revision sets (colors) of the script and how Final Draft shows them.
 *
 * Example:
 *   for _, r := range doc.Revisions.Revision { fmt.Println(r.Name, r.Color) }
 */
type Revisions struct {
	XMLName        xml.Name `json:"-" yaml:"-"`
	ActiveSet      string   `xml:",attr,omitempty" json:"active_set,omitempty" yaml:"active_set,omitempty"`
	Location       string   `xml:",attr,omitempty" json:"location,omitempty" yaml:"location,omitempty"`
	RevisionMode   string   `xml:",attr,omitempty" json:"revision_mode,omitempty" yaml:"revision_mode,omitempty"`
	RevisionsShown string   `xml:",attr,omitempty" json:"revisions_shown,omitempty" yaml:"revisions_shown,omitempty"`
	ShowAllMarks   string   `xml:",attr,omitempty" json:"show_all_marks,omitempty" yaml:"show_all_marks,omitempty"`
	ShowAllSets    string   `xml:",attr,omitempty" json:"show_all_sets,omitempty" yaml:"show_all_sets,omitempty"`
	Revision       []Revision
}

/** Revision is one <Revision> in Revisions: a named revision color with the mark and style used to flag changed text.
 *
 * Example:
 *   r := doc.Revisions.Revision[0]
 *   fmt.Println(r.ID, r.Name)
 */
type Revision struct {
	XMLName      xml.Name `json:"-" yaml:"-"`
	Color        string   `xml:",attr,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	FullRevision string   `xml:",attr,omitempty" json:"full_revision,omitempty" yaml:"full_revision,omitempty"`
	ID           string   `xml:",attr,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Mark         string   `xml:",attr,omitempty" json:"mark,omitempty" yaml:"mark,omitempty"`
	Name         string   `xml:",attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Style        string   `xml:",attr,omitempty" json:"style,omitempty" yaml:"style,omitempty"`
}

/** ElementSettings is an <ElementSettings> element: the font, paragraph layout and behavior Final Draft applies to one paragraph Type, such as Action or Dialogue.
 *
 * Example:
 *   for _, e := range doc.ElementSettings { fmt.Println(e.Type) }
 */
type ElementSettings struct {
	XMLName       xml.Name `json:"-" yaml:"-"`
	Type          string   `xml:",attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	FontSpec      *FontSpec
	ParagraphSpec *ParagraphSpec
	Behavior      *Behavior
}

/** FontSpec is a <FontSpec> element: a font, size, color and style, used by element settings and by the Mores and Continueds, script panel and scene number options.
 *
 * Example:
 *   f := doc.ElementSettings[0].FontSpec
 *   fmt.Println(f.Font, f.Size)
 */
type FontSpec struct {
	XMLName        xml.Name `json:"-" yaml:"-"`
	AdornmentStyle string   `xml:",attr,omitempty" json:"adornment_style,omitempty" yaml:"adornment_type,omitempty"`
	Background     string   `xml:",attr,omitempty" json:"background,omitempty" yaml:"background,omitempty"`
	Color          string   `xml:",attr,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	Font           string   `xml:",attr,omitempty" json:"font,omitempty" yaml:"font,omitempty"`
	RevisionID     string   `xml:",attr,omitempty" json:"revision_id,omitempty" yaml:"revision_id,omitempty"`
	Size           string   `xml:",attr,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	Style          string   `xml:",attr,omitempty" json:"style,omitempty" yaml:"style,omitempty"`
}

/** ParagraphSpec is a <ParagraphSpec> element: the alignment, indents, spacing and page break of a paragraph Type.
 *
 * Example:
 *   p := doc.ElementSettings[0].ParagraphSpec
 *   fmt.Println(p.Alignment, p.LeftIndent)
 */
type ParagraphSpec struct {
	XMLName       xml.Name `json:"-"`
	Alignment     string   `xml:",attr,omitempty" json:"alignment,omitempty" yaml:"alignment,omitempty"`
	FirstIndent   string   `xml:",attr,omitempty" json:"first_indent,omitempty" yaml:"first_indent,omitempty"`
	Leading       string   `xml:",attr,omitempty" json:"leading,omitempty" yaml:"leading,omitempty"`
	LeftIndent    string   `xml:",attr,omitempty" json:"left_indent,omitempty" yaml:"left_indent,omitempty"`
	RightIndent   string   `xml:",attr,omitempty" json:"right_indent,omitempty" yaml:"right_indent,omitempty"`
	SpaceBefore   string   `xml:",attr,omitempty" json:"space_before,omitempty" yaml:"space_before,omitempty"`
	Spacing       string   `xml:",attr,omitempty" json:"spacing,omitempty" yaml:"spacing,omitempty"`
	StartsNewPage string   `xml:",attr,omitempty" json:"starts_new_page,omitempty" yaml:"starts_new_page,omitempty"`
}

/** Behavior is a <Behavior> element: how the editor treats a paragraph Type, the type it paginates as, what the Return key starts next and its keyboard shortcut.
 *
 * Example:
 *   b := doc.ElementSettings[0].Behavior
 *   fmt.Println(b.ReturnKey)
 */
type Behavior struct {
	XMLName    xml.Name `json:"-" yaml:"-"`
	PaginateAs string   `xml:",attr,omitempty" json:"paginate_as,omitempty" yaml:"paginate_as,omitempty"`
	ReturnKey  string   `xml:",attr,omitempty" json:"return_key,omitempty" yaml:"return_key,omitempty"`
	Shortcut   string   `xml:",attr,omitempty" json:"shortcut,omitempty" yaml:"shortcut,omitempty"`
}

/** SpellCheckIgnoreLists is the <SpellCheckIgnoreLists> element: the ranges and words the spell checker skips.
 *
 * Example:
 *   n := len(doc.SpellCheckIgnoreLists.IgnoredWords)
 */
type SpellCheckIgnoreLists struct {
	XMLName       xml.Name `json:"-" yaml:"-"`
	IgnoredRanges *IgnoredRanges
	IgnoredWords  []*IgnoredWords
}

/** IgnoredRanges is the <IgnoredRanges> element of SpellCheckIgnoreLists. It has no content here; Final Draft writes it empty.
 *
 * Example:
 *   ranges := &fdx.IgnoredRanges{}
 */
type IgnoredRanges struct {
	XMLName xml.Name `json:"-" yaml:"-"`
}

/** IgnoredWords is an <IgnoredWords> element: a list of Word the spell checker skips.
 *
 * Example:
 *   for _, w := range doc.SpellCheckIgnoreLists.IgnoredWords[0].Word { fmt.Println(w.InnerText) }
 */
type IgnoredWords struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Word    []*Word
}

/** Word is a <Word> element: one word in an IgnoredWords list.
 *
 * Example:
 *   w := &fdx.Word{InnerText: "Gwyneth"}
 */
type Word struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** PageLayout is the <PageLayout> element: page margins, colors, leading and quote handling for the whole document.
 *
 * Example:
 *   fmt.Println(doc.PageLayout.TopMargin, doc.PageLayout.BottomMargin)
 */
type PageLayout struct {
	XMLName                           xml.Name `json:"-" yaml:"-"`
	BackgroundColor                   string   `xml:",attr,omitempty" json:"background_color,omitempty" yaml:"background_color,omitempty"`
	BottomMargin                      string   `xml:",attr,omitempty" json:"bottom_margin,omitempty" yaml:"bottom_margin,omitempty"`
	BreakDialogueAndActionAtSentences string   `xml:",attr,omitempty" json:"break_dialogue_and_action_at_sentences,omitempty" yaml:"break_dialogue_and_action_at_sentences,omitempty"`
	DocumentLeading                   string   `xml:",attr,omitempty" json:"document_leading,omitempty" yaml:"document_leading,omitempty"`
	FooterMargin                      string   `xml:",attr,omitempty" json:"footer_margin,omitempty" yaml:"footer_margin,omitempty"`
	ForegroundColor                   string   `xml:",attr,omitempty" json:"foreground_color,omitempty" yaml:"foreground_color,omitempty"`
	HeaderMargin                      string   `xml:",attr,omitempty" json:"header_margin,omitempty" yaml:"header_margin,omitempty"`
	InvisiblesColor                   string   `xml:",attr,omitempty" json:"invisible_colors,omitempty" yaml:"invisible_colors,omitempty"`
	TopMargin                         string   `xml:",attr,omitempty" json:"top_margin,omitempty" yaml:"top_margin,omitempty"`
	UsesSmartQuotes                   string   `xml:",attr,omitempty" json:"uses_smart_quotes,omitempty" yaml:"uses_smart_quotes,omitempty"`
	AutoCastList                      *AutoCastList
}

/** AutoCastList is the <AutoCastList> element of PageLayout: whether Final Draft generates the cast list and in which element type.
 *
 * Example:
 *   auto := doc.PageLayout.AutoCastList.AutomaticallyGenerate == "Yes"
 */
type AutoCastList struct {
	XMLName               xml.Name `json:"-" yaml:"-"`
	AddParentheses        string   `xml:",attr,omitempty" json:"add_parentheses,omitempty" yaml:"add_parentheses,omitempty"`
	AutomaticallyGenerate string   `xml:",attr,omitempty" json:"automatically_generate,omitempty" yaml:"automatically_generate,omitempty"`
	CastListElement       string   `xml:",attr,omitempty" json:"cast_list_element,omitempty" yaml:"cast_list_element,omitempty"`
}

/** WindowState is the <WindowState> element: the editor window position, size and mode saved with the document.
 *
 * Example:
 *   fmt.Println(doc.WindowState.Width, doc.WindowState.Height)
 */
type WindowState struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Height  string   `xml:",attr,omitempty" json:"height,omitempty" yaml:"height,omitempty"`
	Left    string   `xml:",attr,omitempty" json:"left,omitempty" yaml:"left,omitempty"`
	Mode    string   `xml:",attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Top     string   `xml:",attr,omitempty" json:"top,omitempty" yaml:"top,omitempty"`
	Width   string   `xml:",attr,omitempty" json:"width,omitempty" yaml:"width,omitempty"`
}

/** TextState is the <TextState> element: the editor scaling, selection and whether invisible characters show.
 *
 * Example:
 *   fmt.Println(doc.TextState.Scaling)
 */
type TextState struct {
	XMLName        xml.Name `json:"-" yaml:"-"`
	Scaling        string   `xml:",attr,omitempty" json:"scaling,omitempty" yaml:"scaling,omitempty"`
	Selection      string   `xml:",attr,omitempty" json:"selection,omitempty" yaml:"selection,omitempty"`
	ShowInvisibles string   `xml:",attr,omitempty" json:"show_invisibles,omitempty" yaml:"show_invisibles,omitempty"`
}

/** ScriptNoteDefinitions is the <ScriptNoteDefinitions> element: the kinds of script note available in the document.
 *
 * Example:
 *   for _, d := range doc.ScriptNoteDefinitions.ScriptNoteDefinition { fmt.Println(d.Name) }
 */
type ScriptNoteDefinitions struct {
	XMLName              xml.Name `json:"-" yaml:"-"`
	Active               string   `xml:",attr,omitempty" json:"active,omitempty" yaml:"active,omitempty"`
	ScriptNoteDefinition []*ScriptNoteDefinition
}

/** ScriptNoteDefinition is one <ScriptNoteDefinition>: a script note kind with its name, color and marker.
 *
 * Example:
 *   d := doc.ScriptNoteDefinitions.ScriptNoteDefinition[0]
 *   fmt.Println(d.Name, d.Color)
 */
type ScriptNoteDefinition struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Color   string   `xml:",attr,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	ID      string   `xml:",attr,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Marker  string   `xml:",attr,omitempty" json:"marker,omitempty" yaml:"marker,omitempty"`
	Name    string   `xml:",attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

/** SmartType is the <SmartType> element: the lists Final Draft uses to complete character names, extensions, scene introductions, locations, times of day and transitions as you type.
 *
 * Example:
 *   for _, c := range doc.SmartType.Characters.Character { fmt.Println(c.InnerText) }
 */
type SmartType struct {
	XMLName     xml.Name `json:"-" yaml:"-"`
	Characters  *Characters
	Extensions  *Extensions
	SceneIntros *SceneIntros
	Locations   *Locations
	TimesOfDay  *TimesOfDay
	Transitions *Transitions
}

/** Characters is the <Characters> element of SmartType: the character names to complete.
 *
 * Example:
 *   n := len(doc.SmartType.Characters.Character)
 */
type Characters struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	Character []*Character
}

/** Character is a <Character> element in SmartType: one character name, in InnerText.
 *
 * Example:
 *   c := &fdx.Character{InnerText: "ANNA"}
 */
type Character struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** Extensions is the <Extensions> element of SmartType: the character name extensions to complete, such as (V.O.).
 *
 * Example:
 *   n := len(doc.SmartType.Extensions.Extension)
 */
type Extensions struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	Extension []*Extension
}

/** Extension is an <Extension> element in SmartType: one character extension, in InnerText.
 *
 * Example:
 *   e := &fdx.Extension{InnerText: "(V.O.)"}
 */
type Extension struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** SceneIntros is the <SceneIntros> element of SmartType: the scene heading introductions to complete, such as INT. and EXT.
 *
 * Example:
 *   n := len(doc.SmartType.SceneIntros.SceneIntro)
 */
type SceneIntros struct {
	XMLName    xml.Name `json:"-" yaml:"-"`
	SceneIntro []*SceneIntro
}

/** SceneIntro is a <SceneIntro> element in SmartType: one scene introduction, in InnerText.
 *
 * Example:
 *   s := &fdx.SceneIntro{InnerText: "INT."}
 */
type SceneIntro struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** Locations is the <Locations> element of SmartType: the scene locations to complete.
 *
 * Example:
 *   n := len(doc.SmartType.Locations.Location)
 */
type Locations struct {
	XMLName  xml.Name `json:"-" yaml:"-"`
	Location []*Location
}

/** Location is a <Location> element in SmartType: one scene location, in InnerText.
 *
 * Example:
 *   l := &fdx.Location{InnerText: "KITCHEN"}
 */
type Location struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** TimesOfDay is the <TimesOfDay> element of SmartType: the times of day to complete and the Separator written before them in a scene heading.
 *
 * Example:
 *   sep := doc.SmartType.TimesOfDay.Separator
 */
type TimesOfDay struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	Separator string   `xml:",attr,omitempty" json:"times_of_day,omitempty" yaml:"times_of_day,omitempty"`
	TimeOfDay []*TimeOfDay
}

/** TimeOfDay is a <TimeOfDay> element in SmartType: one time of day, in InnerText.
 *
 * Example:
 *   t := &fdx.TimeOfDay{InnerText: "NIGHT"}
 */
type TimeOfDay struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** Transitions is the <Transitions> element of SmartType: the transitions to complete, such as CUT TO:.
 *
 * Example:
 *   n := len(doc.SmartType.Transitions.Transition)
 */
type Transitions struct {
	XMLName    xml.Name `json:"-" yaml:"-"`
	Transition []*Transition
}

/** Transition is a <Transition> element in SmartType: one transition, in InnerText.
 *
 * Example:
 *   t := &fdx.Transition{InnerText: "CUT TO:"}
 */
type Transition struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty" yaml:"text,omitempty"`
}

/** MoresAndContinueds is the <MoresAndContinueds> element: the font and the text used where dialogue or a scene breaks across a page.
 *
 * Example:
 *   fmt.Println(doc.MoresAndContinueds.DialogueBreaks.DialogueBottom)
 */
type MoresAndContinueds struct {
	XMLName        xml.Name `json:"-" yaml:"-"`
	FontSpec       *FontSpec
	DialogueBreaks *DialogueBreaks
	SceneBreaks    *SceneBreaks
}

/** DialogueBreaks is the <DialogueBreaks> element: the (MORE) and (CONT'D) text written where dialogue breaks across a page.
 *
 * Example:
 *   fmt.Println(doc.MoresAndContinueds.DialogueBreaks.DialogueTop)
 */
type DialogueBreaks struct {
	XMLNAme        xml.Name `json:"-" yaml:"-"`
	BottomOfPage   string   `xml:",attr,omitempty" json:"bottom_of_page,omitempty" yaml:"bottom_of_page,omitempty"`
	DialogueBottom string   `xml:",attr,omitempty" json:"dialogue_bottom,omitempty" yaml:"dialogue_bottom,omitempty"`
	DialogueTop    string   `xml:",attr,omitempty" json:"dialogue_top,omitempty" yaml:"dialogue_top,omitempty"`
	TopOfNext      string   `xml:",attr,omitempty" json:"top_of_next,omitempty" yaml:"top_of_next,omitempty"`
}

/** SceneBreaks is the <SceneBreaks> element: the CONTINUED text written where a scene breaks across a page.
 *
 * Example:
 *   fmt.Println(doc.MoresAndContinueds.SceneBreaks.SceneBottom)
 */
type SceneBreaks struct {
	XMLName           xml.Name `json:"-" yaml:"-"`
	ContinuedNumber   string   `xml:",attr,omitempty" json:"continued_number,omitempty" yaml:"continued_number,omitempty"`
	SceneBottom       string   `xml:",attr,omitempty" json:"scene_bottom,omitempty" yaml:"scene_buttom,omitempty"`
	SceneBottomOfPage string   `xml:",attr,omitempty" json:"scene_bottom_of_page,omitempty" yaml:"scene_bottom_of_page,omitempty"`
	SceneTop          string   `xml:",attr,omitempty" json:"scene_top,omitempty" yaml:"scene_top,omitempty"`
	SceneTopOfNext    string   `xml:",attr,omitempty" json:"scene_top_of_next,omitempty" yaml:"scene_top_of_next,omitempty"`
}

/** LockedPages is the <LockedPages> element: the pages locked for production. Final Draft writes it empty here.
 *
 * Example:
 *   locked := &fdx.LockedPages{}
 */
type LockedPages struct {
	XMLName xml.Name `json:"-" yaml:"-"`
}

/** Macros is the <Macros> element: the text macros defined in the document.
 *
 * Example:
 *   for _, m := range doc.Macros.Macro { fmt.Println(m.Name) }
 */
type Macros struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Macro   []*Macro
}

/** Macro is one <Macro>: a named shortcut that inserts Text, or an element Alias, optionally followed by a transition.
 *
 * Example:
 *   m := doc.Macros.Macro[0]
 *   fmt.Println(m.Name, m.Shortcut)
 */
type Macro struct {
	XMLName    xml.Name `json:"-" yaml:"-"`
	Element    string   `xml:",attr,omitempty" json:"element,omitempty" yaml:"element,omitempty"`
	Name       string   `xml:",attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Shortcut   string   `xml:",attr,omitempty" json:"shortcut,omitempty" yaml:"shortcut,omitempty"`
	Text       string   `xml:",attr,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	Transition string   `xml:",attr,omitempty" json:"transition,omitempty" yaml:"transition,omitempty"`
	Alias      []*Alias
}

/** Alias is an <Alias> of a Macro: a replacement Text with the options for matching and where it applies.
 *
 * Example:
 *   a := doc.Macros.Macro[0].Alias[0]
 *   fmt.Println(a.Text)
 */
type Alias struct {
	XMLName      xml.Name `json:"-" yaml:"-"`
	Confirm      string   `xml:",attr,omitempty" json:"confirm,omitempty" yaml:"confirm,omitempty"`
	MatchCase    string   `xml:",attr,omitempty" json:"match_case,omitempty" yaml:"match_case,omitempty"`
	SmartReplace string   `xml:",attr,omitempty" json:"smart_replace,omitempty" yaml:"smart_replace,omitempty"`
	Text         string   `xml:",attr,omitempty" json:"text,omitempty" yaml:"text,omitempty"`
	WordOnly     string   `xml:",attr,omitempty" json:"word_only,omitempty" yaml:"word_only,omitempty"`
	ActivateIn   []*ActivateIn
}

/** ActivateIn is an <ActivateIn> element of an Alias: an element type in which the alias applies.
 *
 * Example:
 *   e := doc.Macros.Macro[0].Alias[0].ActivateIn[0].Element
 */
type ActivateIn struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Element string   `xml:",attr,omitempty" json:"element,omitempty" yaml:"element,omitempty"`
}

/** Actors is the <Actors> element: the voices Final Draft uses to read the script aloud.
 *
 * Example:
 *   for _, a := range doc.Actors.Actor { fmt.Println(a.Name) }
 */
type Actors struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Actor   []*Actor
}

/** Actor is one <Actor>: a named voice with the Mac and Windows voice, pitch and speed.
 *
 * Example:
 *   a := doc.Actors.Actor[0]
 *   fmt.Println(a.Name, a.Speed)
 */
type Actor struct {
	XMLName  xml.Name `json:"-" yaml:"-"`
	MacVoice string   `xml:",attr,omitempty" json:"mac_voice,omitempty" yaml:"mac_voice,omitempty"`
	Name     string   `xml:",attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Pitch    string   `xml:",attr,omitempty" json:"pitch,omitempty" yaml:"pitch,omitempty"`
	Speed    string   `xml:",attr,omitempty" json:"speed,omitempty" yaml:"speed,omitempty"`
	WinVoice string   `xml:",attr,omitempty" json:"win_voice,omitempty" yaml:"win_voice,omitempty"`
}

/** Cast is the <Cast> element: who reads which part when Final Draft reads the script aloud, the Narrator and each Member.
 *
 * Example:
 *   for _, m := range doc.Cast.Member { fmt.Println(m.Character, m.Actor) }
 */
type Cast struct {
	XMLName  xml.Name `json:"-" yaml:"-"`
	Narrator *Narrator
	Member   []*Member
}

/** Narrator is the <Narrator> element of Cast: the element types the narrator reads.
 *
 * Example:
 *   n := len(doc.Cast.Narrator.Element)
 */
type Narrator struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Element []*Element
}

/** Element is an <Element> of a Narrator: one paragraph Type the narrator reads.
 *
 * Example:
 *   e := doc.Cast.Narrator.Element[0].Type
 */
type Element struct {
	XMLName xml.Name `json:"-" yaml:"-"`
	Type    string   `xml:",attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

/** Member is a <Member> of Cast: the Actor voice that reads a Character.
 *
 * Example:
 *   m := &fdx.Member{Actor: "Alex", Character: "ANNA"}
 */
type Member struct {
	XMLName   xml.Name `json:"-" yaml:"-"`
	Actor     string   `xml:",attr,omitempty" json:"actor,omitempty" yaml:"actor,omitempty"`
	Character string   `xml:",attr,omitempty" json:"character,omitempty" yaml:"character,omitempty"`
}

/** SplitState is the <SplitState> element: how the editor window is split into panels, saved with the document.
 *
 * Example:
 *   fmt.Println(doc.SplitState.SplitMode)
 */
type SplitState struct {
	XMLName          xml.Name `json:"-" yaml:"-"`
	ActivePanel      string   `xml:",attr,omitempty" json:"active_panel,omitempty" yaml:"active_panel,omitempty"`
	SplitMode        string   `xml:",attr,omitempty" json:"split_mode,omitempty" yaml:"split_mode,omitempty"`
	SplitterPosition string   `xml:",attr,omitempty" json:"splitter_position,omitempty" yaml:"splitter_position,omitempty"`
	ScriptPanel      *ScriptPanel
}

/** ScriptPanel is the <ScriptPanel> element of SplitState: the display mode and font of the script panel.
 *
 * Example:
 *   fmt.Println(doc.SplitState.ScriptPanel.DisplayMode)
 */
type ScriptPanel struct {
	XMLName     xml.Name `json:"-" yaml:"-"`
	DisplayMode string   `xml:",attr,omitempty" json:"display_mode,omitempty" yaml:"display_mode,omitempty"`
	FontSpec    *FontSpec
}

/** SceneNumberOptions is the <SceneNumberOptions> element: where scene numbers print (left, right) and in which font.
 *
 * Example:
 *   left := doc.SceneNumberOptions.ShowNumbersOnLeft == "Yes"
 */
type SceneNumberOptions struct {
	XMLName            xml.Name `json:"-" yaml:"-"`
	LeftLocation       string   `xml:",attr,omitempty" json:"left_location,omitempty" yaml:"left_location,omitempty"`
	RightLocation      string   `xml:",attr,omitempty" json:"right_location,omitempty" yaml:"right_location,omitempty"`
	ShowNumbersOnLeft  string   `xml:",attr,omitempty" json:"show_numbers_on_left,omitempty" yaml:"show_numbers_on_left,omitempty"`
	ShowNumbersOnRight string   `xml:",attr,omitempty" json:"show_numbers_on_right,omitempty" yaml:"show_numbers_on_right,omitempty"`
	FontSpec           *FontSpec
}

/** NewFinalDraft returns a new, empty FinalDraft with the document type "Script" and version "1", ready to be filled in or to receive a screenplay through FromFountain.
 *
 * Returns:
 *   *FinalDraft — a new document
 *
 * Example:
 *   document := fdx.NewFinalDraft()
 *   document.FromFountain(screenplay)
 */
func NewFinalDraft() *FinalDraft {
	document := new(FinalDraft)
	document.DocumentType = "Script"
	document.Version = "1"
	return document
}

/** String (of Text) returns the text of one run in Fountain markup: its styles become *italic*, **bold**, _underline_ and ~~strikethrough~~, and AllCaps upper cases it. A nil Text returns "".
 *
 * Returns:
 *   string — the run as Fountain text
 *
 * Example:
 *   t := &fdx.Text{Style: fdx.BoldStyle, InnerText: "stop"}
 *   fmt.Println(t.String()) // **stop**
 */
func (text *Text) String() string {
	if text != nil {
		src := text.InnerText
		if strings.Contains(text.Style, AllCapsStyle) || strings.Contains(text.Font, "Capitals") {
			src = strings.ToUpper(src)
		}
		if strings.Contains(text.Style, ItalicStyle) {
			src = "*" + src + "*"
		}
		if strings.Contains(text.Style, BoldStyle) {
			src = "**" + src + "**"
		}
		if strings.Contains(text.Style, UnderlineStyle) {
			src = "_" + src + "_"
		}
		if strings.Contains(text.Style, Strikethrough) {
			src = "~~" + src + "~~"
		}
		return src
	}
	return ""
}

/** String (of Paragraph) returns one paragraph as Fountain text, with the trailing blank line its Type needs. A scene heading with a Number ends in #NUMBER#, a centered paragraph is wrapped in > <, and a page break prefixes ===. A nil Paragraph returns "".
 *
 * Returns:
 *   string — the paragraph as Fountain text
 *
 * Example:
 *   fmt.Print(doc.Content.Paragraph[0].String())
 */
func (paragraph *Paragraph) String() string {
	if paragraph != nil {
		src := []string{}
		if paragraph.StartsNewPage == "Yes" {
			src = append(src, "===\n\n")
		}
		for i, text := range paragraph.Text {
			s := text.String()
			switch paragraph.Type {
			case SceneHeadingType:
				s = strings.ToUpper(s)
				if paragraph.Number != "" && i == len(paragraph.Text)-1 {
					// Final Draft's Number attribute is Fountain's #12A#.
					s += " #" + paragraph.Number + "#"
				}
			case CharacterType:
				s = strings.ToUpper(s)
			case TransitionType:
				s = strings.ToUpper(s)
			case SingingType:
				s = "~" + s
			case ParentheticalType:
				if strings.HasPrefix(s, "(") == false &&
					strings.HasSuffix(s, ")") == false {
					s = "(" + s + ")"
				}
			}
			if len(s) > 0 {
				switch paragraph.Alignment {
				case CenterAlignment:
					s = ">" + s + "<"
				}
			}
			src = append(src, s, "\n")
		}
		switch paragraph.Type {
		/*
		   case GeneralType:
		       //src = append(src, "\n")
		*/
		case SceneHeadingType:
			src = append(src, "\n")
		case ActionType:
			src = append(src, "\n")
			/*
			   case CharacterType:
			       //src = append(src, "\n")
			   case ParentheticalType:
			       //src = append(src, "\n")
			*/
		case DialogueType:
			src = append(src, "\n")
		case TransitionType:
			src = append(src, "\n")
			/*
			   case ShotType:
			       //src = append(src, "\n")
			   default:
			       //src = append(src, "\n")
			*/
		}
		return strings.Join(src, "")
	}
	return ""
}

/** String (of Content) returns all the paragraphs as Fountain text. A nil or empty Content returns "".
 *
 * Returns:
 *   string — the content as Fountain text
 *
 * Example:
 *   fmt.Print(doc.Content.String())
 */
func (c *Content) String() string {
	if c != nil && c.Paragraph != nil && len(c.Paragraph) > 0 {
		src := []string{}
		for _, p := range c.Paragraph {
			s := p.String()
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

/** String (of TitlePage) returns the title page content as plain Fountain text, without title page fields. A nil TitlePage returns "".
 *
 * Returns:
 *   string — the title page as text
 *
 * Example:
 *   fmt.Print(doc.TitlePage.String())
 */
func (tp *TitlePage) String() string {
	if tp != nil && tp.Content != nil {
		// Move through Title Page content and render the plain text.
		return tp.Content.String()
	}
	return ""
}

/** String (of FinalDraft) returns the whole screenplay, title page then content, as Fountain text. This is what fdx2txt writes. A nil FinalDraft returns "".
 *
 * Returns:
 *   string — the screenplay as Fountain text
 *
 * Example:
 *   doc, _ := fdx.ParseFile("screenplay.fdx")
 *   fmt.Print(doc.String())
 */
func (doc *FinalDraft) String() string {
	if doc != nil {
		src := []string{}
		if doc.TitlePage != nil {
			s := doc.TitlePage.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s, "\n")
		}
		if doc.Content != nil {
			s := doc.Content.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

/** Parse reads the bytes of an fdx file into a FinalDraft.
 *
 * Parameters:
 *   src ([]byte) — the content of an fdx file
 *
 * Returns:
 *   *FinalDraft — the document (never nil; partly filled when err is not nil)
 *   error — the XML decoding error, nil on success
 *
 * Example:
 *   doc, err := fdx.Parse(src)
 */
func Parse(src []byte) (*FinalDraft, error) {
	document := new(FinalDraft)
	err := xml.Unmarshal(src, &document)
	return document, err
}

/** ParseFile reads the named fdx file into a FinalDraft.
 *
 * Parameters:
 *   fname (string) — the path of an fdx file
 *
 * Returns:
 *   *FinalDraft — the document, nil when the file cannot be read
 *   error — the read or XML decoding error, nil on success
 *
 * Example:
 *   doc, err := fdx.ParseFile("screenplay.fdx")
 */
func ParseFile(fname string) (*FinalDraft, error) {
	src, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return Parse(src)
}

/** CleanupSelfClosingElements rewrites the empty elements Final Draft writes in short form, such as <Actor ...></Actor>, back to <Actor .../>, because Go's XML encoder writes them out in full.
 *
 * Parameters:
 *   src ([]byte) — rendered XML
 *
 * Returns:
 *   []byte — the XML with those elements self-closing
 *
 * Example:
 *   src = fdx.CleanupSelfClosingElements(src)
 */
func CleanupSelfClosingElements(src []byte) []byte {
	selfClosing := []string{
		"SceneProperties",
		"Member",
		"FontSpec",
		"DynamicLabel",
		"IgnoredRanges",
		"AutoCastList",
		"WindowState",
		"TextState",
		"ParagraphSpec",
		"Behavior",
		"ScriptNoteDefinition",
		"DialogueBreaks",
		"SceneBreaks",
		"LockedPages",
		"Revision",
		"ActivateIn",
		"Actor",
		"Element",
	}
	for _, elem := range selfClosing {
		src = bytes.Replace(src, []byte("></"+elem+">"), []byte("/>"), -1)
	}
	return src
}

/** ToXML renders the document as indented fdx XML, with the empty elements self-closing (CleanupSelfClosingElements). The XML declaration is not included; DocString holds it.
 *
 * Returns:
 *   []byte — the XML
 *   error — the encoding error, nil on success
 *
 * Example:
 *   src, err := doc.ToXML()
 */
func (document *FinalDraft) ToXML() ([]byte, error) {
	src, err := xml.MarshalIndent(document, "", "    ")
	if err != nil {
		return nil, err
	}
	return CleanupSelfClosingElements(src), nil
}
