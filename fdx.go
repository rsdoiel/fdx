//
// fdx is a package encoding/decoding fdx formatted XML files.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package fdx

import (
	"encoding/xml"
	"strings"
)

const (
	Version   = `v0.0.0-dev`
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`

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

type FinalDraft struct {
	XMLName               xml.Name `json:"-"`
	DocumentType          string   `xml:",attr,omitempty" json:"type,omitempty"`
	Template              string   `xml:",attr" json:"template,omitempty"`
	Version               string   `xml:",attr" json:"version,omitempty"`
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
	Cast                  *Cast `xml:"Cast,omitempty"`
	SceneNumberOptions    *SceneNumberOptions
}

type Content struct {
	XMLName   xml.Name `json:"-"`
	Paragraph []*Paragraph
}

type Paragraph struct {
	XMLName         xml.Name `json:"-"`
	Type            string   `xml:",attr,omitempty"`
	Number          string   `xml:",attr,omitempty"`
	Alignment       string   `xml:",attr,omitempty"`
	FirstIndent     string   `xml:",attr,omitempty"`
	Leading         string   `xml:",attr,omitempty"`
	LeftIndent      string   `xml:",attr,omitempty"`
	RightIndent     string   `xml:",attr,omitempty"`
	SpaceBefore     string   `xml:",attr,omitempty"`
	Spacing         string   `xml:",attr,omitempty"`
	StartsNewPage   string   `xml:",attr,omitempty"`
	SceneProperties []*SceneProperties
	DynamicLabel    []*DynamicLabel
	Text            []*Text
}

type SceneProperties struct {
	XMLName xml.Name `json:"-"`
	Length  string   `xml:",attr,omitempty"`
	Page    string   `xml:",attr,omitempty"`
	Title   string   `xml:",attr,omitempty"`
}

type HeaderAndFooter struct {
	XMLName         xml.Name `json:"-"`
	FooterFirstPage string   `xml:",attr,omitempty"`
	FooterVisible   string   `xml:",attr,omitempty"`
	HeaderFirstPage string   `xml:",attr,omitempty"`
	HeaderVisible   string   `xml:",attr,omitempty"`
	StartingPage    string   `xml:",attr,omitempty"`
	Header          Header
	Footer          Footer
}

type Header struct {
	XMLName   xml.Name `json:"-"`
	Paragraph []Paragraph
}

type DynamicLabel struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:",attr,omitempty"`
}

type Footer struct {
	XMLName   xml.Name `json:"-"`
	Paragraph []Paragraph
}

type Text struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty"`
	Background     string   `xml:",attr,omitempty"`
	Color          string   `xml:",attr,omitempty"`
	Font           string   `xml:",attr,omitempty"`
	RevisionID     string   `xml:",attr,omitempty"`
	Size           string   `xml:",attr,omitempty"`
	Style          string   `xml:",attr,omitempty"`
	InnerText      string   `xml:",chardata"`
}

type TitlePage struct {
	XMLName         xml.Name `json:"-"`
	HeaderAndFooter *HeaderAndFooter
	Content         *Content
}

type Revisions struct {
	XMLName        xml.Name `json:"-"`
	ActiveSet      string   `xml:",attr,omitempty"`
	Location       string   `xml:",attr,omitempty"`
	RevisionMode   string   `xml:",attr,omitempty"`
	RevisionsShown string   `xml:",attr,omitempty"`
	ShowAllMarks   string   `xml:",attr,omitempty"`
	ShowAllSets    string   `xml:",attr,omitempty"`
	Revision       []Revision
}

type Revision struct {
	XMLName      xml.Name `json:"-"`
	Color        string   `xml:",attr,omitempty"`
	FullRevision string   `xml:",attr,omitempty"`
	ID           string   `xml:",attr,omitempty"`
	Mark         string   `xml:",attr,omitempty"`
	Name         string   `xml:",attr,omitempty"`
	Style        string   `xml:",attr,omitempty"`
}

type ElementSettings struct {
	XMLName       xml.Name `json:"-"`
	Type          string   `xml:",attr,omitempty"`
	FontSpec      *FontSpec
	ParagraphSpec *ParagraphSpec
	Behavior      *Behavior
}

type FontSpec struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty"`
	Background     string   `xml:",attr,omitempty"`
	Color          string   `xml:",attr,omitempty"`
	Font           string   `xml:",attr,omitempty"`
	RevisionID     string   `xml:",attr,omitempty"`
	Size           string   `xml:",attr,omitempty"`
	Style          string   `xml:",attr,omitempty"`
}

type ParagraphSpec struct {
	XMLName       xml.Name `json:"-"`
	Alignment     string   `xml:",attr,omitempty"`
	FirstIndent   string   `xml:",attr,omitempty"`
	Leading       string   `xml:",attr,omitempty"`
	LeftIndent    string   `xml:",attr,omitempty"`
	RightIndent   string   `xml:",attr,omitempty"`
	SpaceBefore   string   `xml:",attr,omitempty"`
	Spacing       string   `xml:",attr,omitempty"`
	StartsNewPage string   `xml:",attr,omitempty"`
}

type Behavior struct {
	XMLName    xml.Name `json:"-"`
	PaginateAs string   `xml:",attr,omitempty"`
	ReturnKey  string   `xml:",attr,omitempty"`
	Shortcut   string   `xml:",attr,omitempty"`
}

type SpellCheckIgnoreLists struct {
	XMLName       xml.Name `json:"-"`
	IgnoredRanges *IgnoredRanges
	IgnoredWords  []*IgnoredWords
}

type IgnoredRanges struct {
	XMLName xml.Name `json:"-"`
}

type IgnoredWords struct {
	XMLName xml.Name `json:"-"`
	Word    []*Word
}

type Word struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type PageLayout struct {
	XMLName                           xml.Name `json:"-"`
	BackgroundColor                   string   `xml:",attr,omitempty"`
	BottomMargin                      string   `xml:",attr,omitempty"`
	BreakDialogueAndActionAtSentences string   `xml:",attr,omitempty"`
	DocumentLeading                   string   `xml:",attr,omitempty"`
	FooterMargin                      string   `xml:",attr,omitempty"`
	ForegroundColor                   string   `xml:",attr,omitempty"`
	HeaderMargin                      string   `xml:",attr,omitempty"`
	InvisiblesColor                   string   `xml:",attr,omitempty"`
	TopMargin                         string   `xml:",attr,omitempty"`
	UsesSmartQuotes                   string   `xml:",attr,omitempty"`
	AutoCastList                      *AutoCastList
}

type AutoCastList struct {
	XMLName               xml.Name `json:"-"`
	AddParentheses        string   `xml:",attr,omitempty"`
	AutomaticallyGenerate string   `xml:",attr,omitempty"`
	CastListElement       string   `xml:",attr,omitempty"`
}

type WindowState struct {
	XMLName xml.Name `json:"-"`
	Height  string   `xml:",attr,omitempty"`
	Left    string   `xml:",attr,omitempty"`
	Mode    string   `xml:",attr,omitempty"`
	Top     string   `xml:",attr,omitempty"`
	Width   string   `xml:",attr,omitempty"`
}

type TextState struct {
	XMLName        xml.Name `json:"-"`
	Scaling        string   `xml:",attr,omitempty"`
	Selection      string   `xml:",attr,omitempty"`
	ShowInvisibles string   `xml:",attr,omitempty"`
}

type ScriptNoteDefinitions struct {
	XMLName              xml.Name `json:"-"`
	Active               string   `xml:",attr,omitempty"`
	ScriptNoteDefinition []*ScriptNoteDefinition
}

type ScriptNoteDefinition struct {
	XMLName xml.Name `json:"-"`
	Color   string   `xml:",attr,omitempty"`
	ID      string   `xml:",attr,omitempty"`
	Marker  string   `xml:",attr,omitempty"`
	Name    string   `xml:",attr,omitempty"`
}

type SmartType struct {
	XMLName     xml.Name `json:"-"`
	Characters  *Characters
	Extensions  *Extensions
	SceneIntros *SceneIntros
	Locations   *Locations
	TimesOfDay  *TimesOfDay
	Transitions *Transitions
}

type Characters struct {
	XMLName   xml.Name `json:"-"`
	Character []*Character
}

type Character struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type Extensions struct {
	XMLName   xml.Name `json:"-"`
	Extension []*Extension
}

type Extension struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type SceneIntros struct {
	XMLName    xml.Name `json:"-"`
	SceneIntro []*SceneIntro
}

type SceneIntro struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type Locations struct {
	XMLName  xml.Name `json:"-"`
	Location []*Location
}

type Location struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type TimesOfDay struct {
	XMLName   xml.Name `json:"-"`
	Separator string   `xml:",attr,omitempty"`
	TimeOfDay []*TimeOfDay
}

type TimeOfDay struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type Transitions struct {
	XMLName    xml.Name `json:"-"`
	Transition []*Transition
}

type Transition struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata"`
}

type MoresAndContinueds struct {
	XMLName        xml.Name `json:"-"`
	FontSpec       *FontSpec
	DialogueBreaks *DialogueBreaks
	SceneBreaks    *SceneBreaks
}

type DialogueBreaks struct {
	XMLNAme        xml.Name `json:"-"`
	BottomOfPage   string   `xml:",attr,omitempty"`
	DialogueBottom string   `xml:",attr,omitempty"`
	DialogueTop    string   `xml:",attr,omitempty"`
	TopOfNext      string   `xml:",attr,omitempty"`
}

type SceneBreaks struct {
	XMLName           xml.Name `json:"-"`
	ContinuedNumber   string   `xml:",attr,omitempty"`
	SceneBottom       string   `xml:",attr,omitempty"`
	SceneBottomOfPage string   `xml:",attr,omitempty"`
	SceneTop          string   `xml:",attr,omitempty"`
	SceneTopOfNext    string   `xml:",attr,omitempty"`
}

type LockedPages struct {
	XMLName xml.Name `json:"-"`
}

type Macros struct {
	XMLName xml.Name `json:"-"`
	Macro   []*Macro
}

type Macro struct {
	XMLName    xml.Name `json:"-"`
	Element    string   `xml:",attr,omitempty"`
	Name       string   `xml:",attr,omitempty"`
	Shortcut   string   `xml:",attr,omitempty"`
	Text       string   `xml:",attr,omitempty"`
	Transition string   `xml:",attr,omitempty"`
	Alias      []*Alias
}

type Alias struct {
	XMLName      xml.Name `json:"-"`
	Confirm      string   `xml:",attr,omitempty"`
	MatchCase    string   `xml:",attr,omitempty"`
	SmartReplace string   `xml:",attr,omitempty"`
	Text         string   `xml:",attr,omitempty"`
	WordOnly     string   `xml:",attr,omitempty"`
	ActivateIn   []*ActivateIn
}

type ActivateIn struct {
	XMLName xml.Name `json:"-"`
	Element string   `xml:",attr,omitempty"`
}

type Actors struct {
	XMLName xml.Name `json:"-"`
	Actor   []*Actor
}

type Actor struct {
	XMLName  xml.Name `json:"-"`
	MacVoice string   `xml:",attr,omitempty"`
	Name     string   `xml:",attr,omitempty"`
	Pitch    string   `xml:",attr,omitempty"`
	Speed    string   `xml:",attr,omitempty"`
	WinVoice string   `xml:",attr,omitempty"`
}

type Cast struct {
	XMLName  xml.Name `json:"-"`
	Narrator *Narrator
	Member   []*Member
}

type Narrator struct {
	XMLName xml.Name `json:"-"`
	Element []*Element
}

type Element struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:",attr,omitempty"`
}

type Member struct {
	XMLName   xml.Name `json:"-"`
	Actor     string   `xml:",attr,omitempty"`
	Character string   `xml:",attr,omitempty"`
}

type SplitState struct {
	XMLName          xml.Name `json:"-"`
	ActivePanel      string   `xml:",attr,omitempty"`
	SplitMode        string   `xml:",attr,omitempty"`
	SplitterPosition string   `xml:",attr,omitempty"`
	ScriptPanel      *ScriptPanel
}

type ScriptPanel struct {
	XMLName     xml.Name `json:"-"`
	DisplayMode string   `xml:",attr,omitempty"`
	FontSpec    *FontSpec
}

type SceneNumberOptions struct {
	XMLName            xml.Name `json:"-"`
	LeftLocation       string   `xml:",attr,omitempty"`
	RightLocation      string   `xml:",attr,omitempty"`
	ShowNumbersOnLeft  string   `xml:",attr,omitempty"`
	ShowNumbersOnRight string   `xml:",attr,omitempty"`
	FontSpec           *FontSpec
}

// String (of FinalDraft) returns a plan text in Fountain format for FinalDraft
func (doc *FinalDraft) String() string {
	if doc != nil {
		src := []string{}
		if doc.TitlePage != nil {
			s := doc.TitlePage.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s)
		}
		if doc.Content != nil {
			s := doc.Content.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s)
		}
		return strings.Join(src, "\n")
	}
	return ""
}

// String (of Text) returns plain text in the Fountain format for a single text element
func (text *Text) String() string {
	if text != nil {
		src := text.InnerText
		//FIXME: Apply attribute formatting instructions here
		if strings.Contains(text.Style, "AllCaps") || strings.Contains(text.Font, "Capitals") {
			src = strings.ToUpper(src)
		}
		if strings.Contains(text.Style, "Italic") {
			src = "*" + src + "*"
		}
		if strings.Contains(text.Style, "Bold") {
			src = "**" + src + "**"
		}
		if strings.Contains(text.Style, "Underline") {
			src = "_" + src + "_"
		}
		return src
	}
	return ""
}

// String (of Paragraph) returns plain text in Fountain format for a single paragraph
func (paragraph *Paragraph) String() string {
	if paragraph != nil && len(paragraph.Text) > 0 {
		src := []string{}
		for _, text := range paragraph.Text {
			s := text.String()
			//FIXME: Apply attribute formatting instructions here
			src = append(src, s)
		}
		switch paragraph.Type {
		case GeneralType:
			src = append(src, "\n")
		case SceneHeadingType:
			src = append(src, "\n\n")
		case ActionType:
			src = append(src, "\n\n")
		case CharacterType:
			src = append(src, "\n")
		case ParentheticalType:
			src = append(src, "\n")
		case DialogueType:
			src = append(src, "\n\n")
		case TransitionType:
			src = append(src, "\n\n")
		case ShotType:
			src = append(src, "\n")
		}
		//FIXME: Make sure I am joining with the correct space characters
		return strings.Join(src, "")
	}
	return ""
}

// String (of Content) returns plain text in Fountain format for Content
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

// String (of TitlePage) returns a plain text in Fountain format (unfielded) for TitlePage
func (tp *TitlePage) String() string {
	if tp != nil && tp.Content != nil && len(tp.Content.Paragraph) > 0 {
		// Move through Title Page content and render the plain text.
		return tp.Content.String()
	}
	return ""
}

// Parse takes []byte and returns a FinalDraft struct and error
func Parse(src []byte) (*FinalDraft, error) {
	document := new(FinalDraft)
	err := xml.Unmarshal(src, &document)
	return document, err
}
