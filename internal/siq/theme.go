package siq

import (
	"encoding/xml"
)

// Theme represents a package theme containing questions.
type Theme struct {
	Name      string      `json:"name" xml:"name,attr"`
	Questions []*Question `json:"questions"`
	Info      Info        `json:"info,omitempty"`
}

// NewTheme creates a new theme with the given name.
func NewTheme(name string) *Theme {
	return &Theme{
		Name:      name,
		Questions: make([]*Question, 0),
	}
}

// Clone creates a deep copy of Theme.
func (t *Theme) Clone() *Theme {
	clone := &Theme{
		Name: t.Name,
		Info: t.Info.Clone(),
	}
	if t.Questions != nil {
		clone.Questions = make([]*Question, len(t.Questions))
		for i, q := range t.Questions {
			clone.Questions[i] = q.Clone()
		}
	}
	return clone
}

// MarshalXML implements xml.Marshaler for Theme.
func (t *Theme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "theme"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "name"}, Value: t.Name},
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Info
	if !t.Info.IsEmpty() {
		if err := e.Encode(&t.Info); err != nil {
			return err
		}
	}
	// Questions
	if len(t.Questions) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "questions"}}); err != nil {
			return err
		}
		for _, q := range t.Questions {
			if err := e.Encode(q); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "questions"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Theme.
func (t *Theme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			t.Name = attr.Value
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "info":
				if err := d.DecodeElement(&t.Info, &tok); err != nil {
					return err
				}
			case "questions":
				// Continue to parse questions
			case "question":
				var q Question
				if err := d.DecodeElement(&q, &tok); err != nil {
					return err
				}
				t.Questions = append(t.Questions, &q)
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if tok.Name.Local == "theme" {
				return nil
			}
		}
	}
}
