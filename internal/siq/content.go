package siq

import (
	"encoding/xml"
	"time"
)

// ContentItem represents a piece of content in a question.
type ContentItem struct {
	Type          string        `json:"type" xml:"type,attr,omitempty"`
	Value         string        `json:"value" xml:",chardata"`
	IsRef         bool          `json:"isRef,omitempty" xml:"isRef,attr,omitempty"`
	Placement     string        `json:"placement,omitempty" xml:"placement,attr,omitempty"`
	Duration      time.Duration `json:"duration,omitempty"`
	WaitForFinish bool          `json:"waitForFinish,omitempty" xml:"waitForFinish,attr,omitempty"`
}

// xmlContentItem is used for XML serialization.
type xmlContentItem struct {
	XMLName       xml.Name `xml:"item"`
	Type          string   `xml:"type,attr,omitempty"`
	Value         string   `xml:",chardata"`
	IsRef         string   `xml:"isRef,attr,omitempty"`
	Placement     string   `xml:"placement,attr,omitempty"`
	Duration      string   `xml:"duration,attr,omitempty"`
	WaitForFinish string   `xml:"waitForFinish,attr,omitempty"`
}

// GetDefaultPlacement returns the default placement for the content type.
func (c *ContentItem) GetDefaultPlacement() string {
	if c.Type == ContentTypeAudio {
		return PlacementBackground
	}
	return PlacementScreen
}

// GetPlacement returns the effective placement.
func (c *ContentItem) GetPlacement() string {
	if c.Placement == "" {
		return c.GetDefaultPlacement()
	}
	return c.Placement
}

// Clone creates a deep copy of ContentItem.
func (c *ContentItem) Clone() ContentItem {
	return ContentItem{
		Type:          c.Type,
		Value:         c.Value,
		IsRef:         c.IsRef,
		Placement:     c.Placement,
		Duration:      c.Duration,
		WaitForFinish: c.WaitForFinish,
	}
}

// NumberSet represents a numeric range for prices.
type NumberSet struct {
	Minimum int `json:"minimum" xml:"minimum,attr"`
	Maximum int `json:"maximum" xml:"maximum,attr"`
	Step    int `json:"step" xml:"step,attr"`
}

// Clone creates a deep copy of NumberSet.
func (n *NumberSet) Clone() NumberSet {
	return NumberSet{
		Minimum: n.Minimum,
		Maximum: n.Maximum,
		Step:    n.Step,
	}
}

// StepParameter represents a parameter in a question step.
type StepParameter struct {
	Type           string           `json:"type,omitempty"`
	SimpleValue    string           `json:"simpleValue,omitempty"`
	ContentValue   []ContentItem    `json:"contentValue,omitempty"`
	GroupValue     map[string]*StepParameter `json:"groupValue,omitempty"`
	NumberSetValue *NumberSet       `json:"numberSetValue,omitempty"`
	IsRef          bool             `json:"isRef,omitempty"`
}

// Clone creates a deep copy of StepParameter.
func (s *StepParameter) Clone() *StepParameter {
	if s == nil {
		return nil
	}
	clone := &StepParameter{
		Type:        s.Type,
		SimpleValue: s.SimpleValue,
		IsRef:       s.IsRef,
	}
	if s.ContentValue != nil {
		clone.ContentValue = make([]ContentItem, len(s.ContentValue))
		for i, item := range s.ContentValue {
			clone.ContentValue[i] = item.Clone()
		}
	}
	if s.GroupValue != nil {
		clone.GroupValue = make(map[string]*StepParameter)
		for k, v := range s.GroupValue {
			clone.GroupValue[k] = v.Clone()
		}
	}
	if s.NumberSetValue != nil {
		ns := s.NumberSetValue.Clone()
		clone.NumberSetValue = &ns
	}
	return clone
}

// Step represents a script step in a question.
type Step struct {
	Type       string                    `json:"type"`
	Parameters map[string]*StepParameter `json:"parameters,omitempty"`
}

// Clone creates a deep copy of Step.
func (s *Step) Clone() Step {
	clone := Step{
		Type: s.Type,
	}
	if s.Parameters != nil {
		clone.Parameters = make(map[string]*StepParameter)
		for k, v := range s.Parameters {
			clone.Parameters[k] = v.Clone()
		}
	}
	return clone
}

// Script represents a question play script.
type Script struct {
	Steps []Step `json:"steps"`
}

// Clone creates a deep copy of Script.
func (s *Script) Clone() *Script {
	if s == nil {
		return nil
	}
	clone := &Script{
		Steps: make([]Step, len(s.Steps)),
	}
	for i, step := range s.Steps {
		clone.Steps[i] = step.Clone()
	}
	return clone
}
