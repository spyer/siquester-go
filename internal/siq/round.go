package siq

import (
	"encoding/xml"
)

// Round represents a game round containing themes.
type Round struct {
	Name   string   `json:"name" xml:"name,attr"`
	Type   string   `json:"type,omitempty" xml:"type,attr,omitempty"`
	Themes []*Theme `json:"themes"`
	Info   Info     `json:"info,omitempty"`
}

// NewRound creates a new round with the given name.
func NewRound(name string) *Round {
	return &Round{
		Name:   name,
		Type:   RoundTypeStandard,
		Themes: make([]*Theme, 0),
	}
}

// Clone creates a deep copy of Round.
func (r *Round) Clone() *Round {
	clone := &Round{
		Name: r.Name,
		Type: r.Type,
		Info: r.Info.Clone(),
	}
	if r.Themes != nil {
		clone.Themes = make([]*Theme, len(r.Themes))
		for i, t := range r.Themes {
			clone.Themes[i] = t.Clone()
		}
	}
	return clone
}

// MarshalXML implements xml.Marshaler for Round.
func (r *Round) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "round"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "name"}, Value: r.Name},
	}
	if r.Type != "" && r.Type != RoundTypeStandard {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: r.Type})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Info
	if !r.Info.IsEmpty() {
		if err := e.Encode(&r.Info); err != nil {
			return err
		}
	}
	// Themes
	if len(r.Themes) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "themes"}}); err != nil {
			return err
		}
		for _, t := range r.Themes {
			if err := e.Encode(t); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "themes"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Round.
func (r *Round) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	r.Type = RoundTypeStandard
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			r.Name = attr.Value
		case "type":
			r.Type = attr.Value
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
				if err := d.DecodeElement(&r.Info, &tok); err != nil {
					return err
				}
			case "themes":
				// Continue to parse themes
			case "theme":
				var t Theme
				if err := d.DecodeElement(&t, &tok); err != nil {
					return err
				}
				r.Themes = append(r.Themes, &t)
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if tok.Name.Local == "round" {
				return nil
			}
		}
	}
}
