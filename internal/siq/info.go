package siq

import (
	"encoding/xml"
)

// Info contains package item information (authors, sources, comments).
type Info struct {
	Authors         []string `json:"authors,omitempty"`
	Sources         []string `json:"sources,omitempty"`
	Comments        string   `json:"comments,omitempty"`
	ShowmanComments string   `json:"showmanComments,omitempty"`
	Extension       string   `json:"extension,omitempty"`
}

// xmlInfo is used for XML serialization of Info.
type xmlInfo struct {
	Authors         *xmlStringList `xml:"authors,omitempty"`
	Sources         *xmlStringList `xml:"sources,omitempty"`
	Comments        string         `xml:"comments,omitempty"`
	ShowmanComments string         `xml:"showmanComments,omitempty"`
	Extension       string         `xml:"extension,omitempty"`
}

type xmlStringList struct {
	Items []string `xml:",any"`
}

type xmlAuthor struct {
	XMLName xml.Name `xml:"author"`
	Value   string   `xml:",chardata"`
}

type xmlSource struct {
	XMLName xml.Name `xml:"source"`
	Value   string   `xml:",chardata"`
}

// MarshalXML implements xml.Marshaler for Info.
func (i *Info) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "info"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Authors
	if len(i.Authors) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "authors"}}); err != nil {
			return err
		}
		for _, author := range i.Authors {
			if err := e.EncodeElement(author, xml.StartElement{Name: xml.Name{Local: "author"}}); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "authors"}}); err != nil {
			return err
		}
	}
	// Sources
	if len(i.Sources) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "sources"}}); err != nil {
			return err
		}
		for _, source := range i.Sources {
			if err := e.EncodeElement(source, xml.StartElement{Name: xml.Name{Local: "source"}}); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "sources"}}); err != nil {
			return err
		}
	}
	// Comments
	if i.Comments != "" {
		if err := e.EncodeElement(i.Comments, xml.StartElement{Name: xml.Name{Local: "comments"}}); err != nil {
			return err
		}
	}
	// Showman Comments
	if i.ShowmanComments != "" {
		if err := e.EncodeElement(i.ShowmanComments, xml.StartElement{Name: xml.Name{Local: "showmanComments"}}); err != nil {
			return err
		}
	}
	// Extension
	if i.Extension != "" {
		if err := e.EncodeElement(i.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Info.
func (i *Info) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "authors":
				if err := i.unmarshalAuthors(d); err != nil {
					return err
				}
			case "sources":
				if err := i.unmarshalSources(d); err != nil {
					return err
				}
			case "comments":
				var s string
				if err := d.DecodeElement(&s, &t); err != nil {
					return err
				}
				i.Comments = s
			case "showmanComments":
				var s string
				if err := d.DecodeElement(&s, &t); err != nil {
					return err
				}
				i.ShowmanComments = s
			case "extension":
				var s string
				if err := d.DecodeElement(&s, &t); err != nil {
					return err
				}
				i.Extension = s
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}

func (i *Info) unmarshalAuthors(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "author" {
				var author string
				if err := d.DecodeElement(&author, &t); err != nil {
					return err
				}
				i.Authors = append(i.Authors, author)
			}
		case xml.EndElement:
			if t.Name.Local == "authors" {
				return nil
			}
		}
	}
}

func (i *Info) unmarshalSources(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "source" {
				var source string
				if err := d.DecodeElement(&source, &t); err != nil {
					return err
				}
				i.Sources = append(i.Sources, source)
			}
		case xml.EndElement:
			if t.Name.Local == "sources" {
				return nil
			}
		}
	}
}

// IsEmpty returns true if the info has no content.
func (i *Info) IsEmpty() bool {
	return len(i.Authors) == 0 && len(i.Sources) == 0 && i.Comments == "" && i.ShowmanComments == "" && i.Extension == ""
}

// Clone creates a deep copy of Info.
func (i *Info) Clone() Info {
	clone := Info{
		Comments:        i.Comments,
		ShowmanComments: i.ShowmanComments,
		Extension:       i.Extension,
	}
	if len(i.Authors) > 0 {
		clone.Authors = make([]string, len(i.Authors))
		copy(clone.Authors, i.Authors)
	}
	if len(i.Sources) > 0 {
		clone.Sources = make([]string, len(i.Sources))
		copy(clone.Sources, i.Sources)
	}
	return clone
}
