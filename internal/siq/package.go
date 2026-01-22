package siq

import (
	"encoding/xml"
	"strconv"
)

// AuthorInfo represents detailed author information.
type AuthorInfo struct {
	ID         string `json:"id" xml:"id,attr"`
	Name       string `json:"name,omitempty"`
	SecondName string `json:"secondName,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
}

// SourceInfo represents detailed source information.
type SourceInfo struct {
	ID      string `json:"id" xml:"id,attr"`
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Year    string `json:"year,omitempty"`
	Publish string `json:"publish,omitempty"`
	City    string `json:"city,omitempty"`
}

// GlobalData contains global package data (authors and sources collections).
type GlobalData struct {
	Authors []AuthorInfo `json:"authors,omitempty"`
	Sources []SourceInfo `json:"sources,omitempty"`
}

// Package represents a SIGame package.
type Package struct {
	Name       string     `json:"name" xml:"name,attr"`
	Version    float64    `json:"version" xml:"version,attr"`
	ID         string     `json:"id,omitempty" xml:"id,attr,omitempty"`
	Restriction string    `json:"restriction,omitempty" xml:"restriction,attr,omitempty"`
	Date       string     `json:"date,omitempty" xml:"date,attr,omitempty"`
	Publisher  string     `json:"publisher,omitempty" xml:"publisher,attr,omitempty"`
	ContactURI string     `json:"contactUri,omitempty" xml:"contactUri,attr,omitempty"`
	Difficulty int        `json:"difficulty,omitempty" xml:"difficulty,attr,omitempty"`
	Logo       string     `json:"logo,omitempty" xml:"logo,attr,omitempty"`
	Language   string     `json:"language,omitempty" xml:"language,attr,omitempty"`
	Tags       []string   `json:"tags,omitempty"`
	Global     GlobalData `json:"global,omitempty"`
	Rounds     []*Round   `json:"rounds"`
	Info       Info       `json:"info,omitempty"`
}

// NewPackage creates a new package with default values.
func NewPackage(name, author string) *Package {
	return &Package{
		Name:    name,
		Version: PackageVersion,
		Rounds:  make([]*Round, 0),
		Info: Info{
			Authors: []string{author},
		},
	}
}

// Clone creates a deep copy of Package.
func (p *Package) Clone() *Package {
	clone := &Package{
		Name:        p.Name,
		Version:     p.Version,
		ID:          p.ID,
		Restriction: p.Restriction,
		Date:        p.Date,
		Publisher:   p.Publisher,
		ContactURI:  p.ContactURI,
		Difficulty:  p.Difficulty,
		Logo:        p.Logo,
		Language:    p.Language,
		Info:        p.Info.Clone(),
	}
	if p.Tags != nil {
		clone.Tags = make([]string, len(p.Tags))
		copy(clone.Tags, p.Tags)
	}
	if p.Rounds != nil {
		clone.Rounds = make([]*Round, len(p.Rounds))
		for i, r := range p.Rounds {
			clone.Rounds[i] = r.Clone()
		}
	}
	return clone
}

// MarshalXML implements xml.Marshaler for Package.
func (p *Package) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "package"
	start.Name.Space = "https://github.com/VladimirKhil/SI/blob/master/assets/siq_5.xsd"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "name"}, Value: p.Name},
		{Name: xml.Name{Local: "version"}, Value: strconv.FormatFloat(p.Version, 'f', -1, 64)},
	}
	if p.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"}, Value: p.ID})
	}
	if p.Restriction != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "restriction"}, Value: p.Restriction})
	}
	if p.Date != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "date"}, Value: p.Date})
	}
	if p.Publisher != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "publisher"}, Value: p.Publisher})
	}
	if p.ContactURI != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "contactUri"}, Value: p.ContactURI})
	}
	if p.Difficulty > 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "difficulty"}, Value: strconv.Itoa(p.Difficulty)})
	}
	if p.Logo != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "logo"}, Value: p.Logo})
	}
	if p.Language != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "language"}, Value: p.Language})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Tags
	if len(p.Tags) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "tags"}}); err != nil {
			return err
		}
		for _, tag := range p.Tags {
			if err := e.EncodeElement(tag, xml.StartElement{Name: xml.Name{Local: "tag"}}); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "tags"}}); err != nil {
			return err
		}
	}
	// Global
	if len(p.Global.Authors) > 0 || len(p.Global.Sources) > 0 {
		if err := encodeGlobalData(e, p.Global); err != nil {
			return err
		}
	}
	// Info
	if !p.Info.IsEmpty() {
		if err := e.Encode(&p.Info); err != nil {
			return err
		}
	}
	// Rounds
	if len(p.Rounds) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "rounds"}}); err != nil {
			return err
		}
		for _, r := range p.Rounds {
			if err := e.Encode(r); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "rounds"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func encodeGlobalData(e *xml.Encoder, g GlobalData) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "global"}}); err != nil {
		return err
	}
	// Authors
	if len(g.Authors) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "Authors"}}); err != nil {
			return err
		}
		for _, author := range g.Authors {
			if err := encodeAuthorInfo(e, author); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Authors"}}); err != nil {
			return err
		}
	}
	// Sources
	if len(g.Sources) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "Sources"}}); err != nil {
			return err
		}
		for _, source := range g.Sources {
			if err := encodeSourceInfo(e, source); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "Sources"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "global"}})
}

func encodeAuthorInfo(e *xml.Encoder, a AuthorInfo) error {
	start := xml.StartElement{
		Name: xml.Name{Local: "Author"},
		Attr: []xml.Attr{{Name: xml.Name{Local: "id"}, Value: a.ID}},
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if a.Name != "" {
		if err := e.EncodeElement(a.Name, xml.StartElement{Name: xml.Name{Local: "Name"}}); err != nil {
			return err
		}
	}
	if a.SecondName != "" {
		if err := e.EncodeElement(a.SecondName, xml.StartElement{Name: xml.Name{Local: "SecondName"}}); err != nil {
			return err
		}
	}
	if a.Surname != "" {
		if err := e.EncodeElement(a.Surname, xml.StartElement{Name: xml.Name{Local: "Surname"}}); err != nil {
			return err
		}
	}
	if a.Country != "" {
		if err := e.EncodeElement(a.Country, xml.StartElement{Name: xml.Name{Local: "Country"}}); err != nil {
			return err
		}
	}
	if a.City != "" {
		if err := e.EncodeElement(a.City, xml.StartElement{Name: xml.Name{Local: "City"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func encodeSourceInfo(e *xml.Encoder, s SourceInfo) error {
	start := xml.StartElement{
		Name: xml.Name{Local: "Source"},
		Attr: []xml.Attr{{Name: xml.Name{Local: "id"}, Value: s.ID}},
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if s.Author != "" {
		if err := e.EncodeElement(s.Author, xml.StartElement{Name: xml.Name{Local: "Author"}}); err != nil {
			return err
		}
	}
	if s.Title != "" {
		if err := e.EncodeElement(s.Title, xml.StartElement{Name: xml.Name{Local: "Title"}}); err != nil {
			return err
		}
	}
	if s.Year != "" {
		if err := e.EncodeElement(s.Year, xml.StartElement{Name: xml.Name{Local: "Year"}}); err != nil {
			return err
		}
	}
	if s.Publish != "" {
		if err := e.EncodeElement(s.Publish, xml.StartElement{Name: xml.Name{Local: "Publish"}}); err != nil {
			return err
		}
	}
	if s.City != "" {
		if err := e.EncodeElement(s.City, xml.StartElement{Name: xml.Name{Local: "City"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// UnmarshalXML implements xml.Unmarshaler for Package.
func (p *Package) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Version = PackageVersion
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			p.Name = attr.Value
		case "version":
			v, _ := strconv.ParseFloat(attr.Value, 64)
			p.Version = v
		case "id":
			p.ID = attr.Value
		case "restriction":
			p.Restriction = attr.Value
		case "date":
			p.Date = attr.Value
		case "publisher":
			p.Publisher = attr.Value
		case "contactUri":
			p.ContactURI = attr.Value
		case "difficulty":
			d, _ := strconv.Atoi(attr.Value)
			p.Difficulty = d
		case "logo":
			p.Logo = attr.Value
		case "language":
			p.Language = attr.Value
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
			case "tags":
				if err := p.decodeTags(d); err != nil {
					return err
				}
			case "global":
				if err := p.decodeGlobal(d); err != nil {
					return err
				}
			case "info":
				if err := d.DecodeElement(&p.Info, &tok); err != nil {
					return err
				}
			case "rounds":
				// Continue to parse rounds
			case "round":
				var r Round
				if err := d.DecodeElement(&r, &tok); err != nil {
					return err
				}
				p.Rounds = append(p.Rounds, &r)
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if tok.Name.Local == "package" {
				return nil
			}
		}
	}
}

func (p *Package) decodeTags(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			if tok.Name.Local == "tag" {
				var tag string
				if err := d.DecodeElement(&tag, &tok); err != nil {
					return err
				}
				p.Tags = append(p.Tags, tag)
			}
		case xml.EndElement:
			if tok.Name.Local == "tags" {
				return nil
			}
		}
	}
}

func (p *Package) decodeGlobal(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "Authors":
				if err := p.decodeGlobalAuthors(d); err != nil {
					return err
				}
			case "Sources":
				if err := p.decodeGlobalSources(d); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if tok.Name.Local == "global" {
				return nil
			}
		}
	}
}

func (p *Package) decodeGlobalAuthors(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			if tok.Name.Local == "Author" {
				author := AuthorInfo{}
				for _, attr := range tok.Attr {
					if attr.Name.Local == "id" {
						author.ID = attr.Value
					}
				}
				if err := decodeAuthorFields(d, &author); err != nil {
					return err
				}
				p.Global.Authors = append(p.Global.Authors, author)
			}
		case xml.EndElement:
			if tok.Name.Local == "Authors" {
				return nil
			}
		}
	}
}

func decodeAuthorFields(d *xml.Decoder, a *AuthorInfo) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var s string
			if err := d.DecodeElement(&s, &tok); err != nil {
				return err
			}
			switch tok.Name.Local {
			case "Name":
				a.Name = s
			case "SecondName":
				a.SecondName = s
			case "Surname":
				a.Surname = s
			case "Country":
				a.Country = s
			case "City":
				a.City = s
			}
		case xml.EndElement:
			if tok.Name.Local == "Author" {
				return nil
			}
		}
	}
}

func (p *Package) decodeGlobalSources(d *xml.Decoder) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			if tok.Name.Local == "Source" {
				source := SourceInfo{}
				for _, attr := range tok.Attr {
					if attr.Name.Local == "id" {
						source.ID = attr.Value
					}
				}
				if err := decodeSourceFields(d, &source); err != nil {
					return err
				}
				p.Global.Sources = append(p.Global.Sources, source)
			}
		case xml.EndElement:
			if tok.Name.Local == "Sources" {
				return nil
			}
		}
	}
}

func decodeSourceFields(d *xml.Decoder, s *SourceInfo) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var str string
			if err := d.DecodeElement(&str, &tok); err != nil {
				return err
			}
			switch tok.Name.Local {
			case "Author":
				s.Author = str
			case "Title":
				s.Title = str
			case "Year":
				s.Year = str
			case "Publish":
				s.Publish = str
			case "City":
				s.City = str
			}
		case xml.EndElement:
			if tok.Name.Local == "Source" {
				return nil
			}
		}
	}
}
