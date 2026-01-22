package siq

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Question represents a game question.
type Question struct {
	Price      int                       `json:"price"`
	TypeName   string                    `json:"typeName,omitempty"`
	Script     *Script                   `json:"script,omitempty"`
	Parameters map[string]*StepParameter `json:"parameters,omitempty"`
	Right      []string                  `json:"right"`
	Wrong      []string                  `json:"wrong,omitempty"`
	Info       Info                      `json:"info,omitempty"`
}

// NewQuestion creates a new question with default values.
func NewQuestion(price int) *Question {
	return &Question{
		Price:      price,
		TypeName:   QuestionTypeDefault,
		Parameters: make(map[string]*StepParameter),
		Right:      []string{""},
	}
}

// Clone creates a deep copy of Question.
func (q *Question) Clone() *Question {
	clone := &Question{
		Price:    q.Price,
		TypeName: q.TypeName,
		Script:   q.Script.Clone(),
		Info:     q.Info.Clone(),
	}
	if q.Parameters != nil {
		clone.Parameters = make(map[string]*StepParameter)
		for k, v := range q.Parameters {
			clone.Parameters[k] = v.Clone()
		}
	}
	if q.Right != nil {
		clone.Right = make([]string, len(q.Right))
		copy(clone.Right, q.Right)
	}
	if q.Wrong != nil {
		clone.Wrong = make([]string, len(q.Wrong))
		copy(clone.Wrong, q.Wrong)
	}
	return clone
}

// GetText returns the question text from content.
func (q *Question) GetText() string {
	if q.Parameters == nil {
		return ""
	}
	if param, ok := q.Parameters["question"]; ok {
		return getTextFromContent(param)
	}
	return ""
}

func getTextFromContent(param *StepParameter) string {
	if param == nil {
		return ""
	}
	if param.ContentValue != nil {
		for _, item := range param.ContentValue {
			if item.Type == ContentTypeText || item.Type == "" {
				return item.Value
			}
		}
	}
	return param.SimpleValue
}

// MarshalXML implements xml.Marshaler for Question.
func (q *Question) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "question"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "price"}, Value: strconv.Itoa(q.Price)},
	}
	if q.TypeName != "" && q.TypeName != QuestionTypeDefault {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: q.TypeName})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	// Info
	if !q.Info.IsEmpty() {
		if err := e.Encode(&q.Info); err != nil {
			return err
		}
	}
	// Script
	if q.Script != nil && len(q.Script.Steps) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "script"}}); err != nil {
			return err
		}
		for _, step := range q.Script.Steps {
			if err := encodeStep(e, step); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "script"}}); err != nil {
			return err
		}
	}
	// Parameters - always write params block (required by canonical SIGame client)
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "params"}}); err != nil {
		return err
	}
	// Ensure there's at least a question param
	hasQuestionParam := false
	for name, param := range q.Parameters {
		if name == "question" {
			hasQuestionParam = true
		}
		if err := encodeParameter(e, name, param); err != nil {
			return err
		}
	}
	// Add empty question param if missing
	if !hasQuestionParam {
		emptyParam := &StepParameter{
			Type:         StepParamTypeContent,
			ContentValue: []ContentItem{{Type: ContentTypeText, Value: ""}},
		}
		if err := encodeParameter(e, "question", emptyParam); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "params"}}); err != nil {
		return err
	}
	// Right answers
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "right"}}); err != nil {
		return err
	}
	for _, answer := range q.Right {
		if err := e.EncodeElement(answer, xml.StartElement{Name: xml.Name{Local: "answer"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "right"}}); err != nil {
		return err
	}
	// Wrong answers
	if len(q.Wrong) > 0 {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "wrong"}}); err != nil {
			return err
		}
		for _, answer := range q.Wrong {
			if err := e.EncodeElement(answer, xml.StartElement{Name: xml.Name{Local: "answer"}}); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "wrong"}}); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func encodeStep(e *xml.Encoder, step Step) error {
	start := xml.StartElement{Name: xml.Name{Local: "step"}}
	if step.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: step.Type})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for name, param := range step.Parameters {
		if err := encodeParameter(e, name, param); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func encodeParameter(e *xml.Encoder, name string, param *StepParameter) error {
	// Infer type from content if not explicitly set (must be done BEFORE writing start element)
	effectiveType := param.Type
	if effectiveType == "" || effectiveType == StepParamTypeSimple {
		if len(param.ContentValue) > 0 {
			effectiveType = StepParamTypeContent
		} else if len(param.GroupValue) > 0 {
			effectiveType = StepParamTypeGroup
		}
	}
	start := xml.StartElement{Name: xml.Name{Local: "param"}}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"}, Value: name})
	if effectiveType != "" && effectiveType != StepParamTypeSimple {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: effectiveType})
	}
	if param.IsRef {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isRef"}, Value: "True"})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	switch effectiveType {
	case StepParamTypeContent:
		for _, item := range param.ContentValue {
			if err := encodeContentItem(e, item); err != nil {
				return err
			}
		}
	case StepParamTypeGroup:
		for k, v := range param.GroupValue {
			if err := encodeParameter(e, k, v); err != nil {
				return err
			}
		}
	case StepParamTypeNumberSet:
		if param.NumberSetValue != nil {
			nsStart := xml.StartElement{
				Name: xml.Name{Local: "numberSet"},
				Attr: []xml.Attr{
					{Name: xml.Name{Local: "minimum"}, Value: strconv.Itoa(param.NumberSetValue.Minimum)},
					{Name: xml.Name{Local: "maximum"}, Value: strconv.Itoa(param.NumberSetValue.Maximum)},
					{Name: xml.Name{Local: "step"}, Value: strconv.Itoa(param.NumberSetValue.Step)},
				},
			}
			if err := e.EncodeToken(nsStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.EndElement{Name: nsStart.Name}); err != nil {
				return err
			}
		}
	default:
		if err := e.EncodeToken(xml.CharData(param.SimpleValue)); err != nil {
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func encodeContentItem(e *xml.Encoder, item ContentItem) error {
	start := xml.StartElement{Name: xml.Name{Local: "item"}}
	if item.Type != "" && item.Type != ContentTypeText {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"}, Value: item.Type})
	}
	if item.IsRef {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isRef"}, Value: "True"})
	}
	if item.Placement != "" && item.Placement != item.GetDefaultPlacement() {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "placement"}, Value: item.Placement})
	} else if item.Type == ContentTypeAudio {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "placement"}, Value: item.GetDefaultPlacement()})
	}
	if item.Duration > 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "duration"}, Value: formatDuration(item.Duration)})
	}
	if !item.WaitForFinish && item.Duration > 0 {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "waitForFinish"}, Value: "False"})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(item.Value)); err != nil {
		return err
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// UnmarshalXML implements xml.Unmarshaler for Question.
func (q *Question) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	q.Parameters = make(map[string]*StepParameter)
	q.TypeName = QuestionTypeDefault
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "price":
			price, _ := strconv.Atoi(attr.Value)
			q.Price = price
		case "type":
			q.TypeName = attr.Value
		}
	}
	var inRight, inWrong bool
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "info":
				if err := d.DecodeElement(&q.Info, &t); err != nil {
					return err
				}
			case "script":
				script, err := decodeScript(d)
				if err != nil {
					return err
				}
				q.Script = script
			case "params":
				params, err := decodeParameters(d)
				if err != nil {
					return err
				}
				q.Parameters = params
			case "right":
				inRight = true
				inWrong = false
			case "wrong":
				inRight = false
				inWrong = true
			case "answer":
				var answer string
				if err := d.DecodeElement(&answer, &t); err != nil {
					return err
				}
				if inRight {
					q.Right = append(q.Right, answer)
				} else if inWrong {
					q.Wrong = append(q.Wrong, answer)
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if t.Name.Local == "question" {
				if len(q.Right) == 0 {
					q.Right = []string{""}
				}
				return nil
			}
		}
	}
}

func decodeScript(d *xml.Decoder) (*Script, error) {
	script := &Script{}
	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "step" {
				step, err := decodeStep(d, t)
				if err != nil {
					return nil, err
				}
				script.Steps = append(script.Steps, step)
			}
		case xml.EndElement:
			if t.Name.Local == "script" {
				return script, nil
			}
		}
	}
}

func decodeStep(d *xml.Decoder, start xml.StartElement) (Step, error) {
	step := Step{Parameters: make(map[string]*StepParameter)}
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			step.Type = attr.Value
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return step, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "param" {
				name, param, err := decodeParameter(d, t)
				if err != nil {
					return step, err
				}
				step.Parameters[name] = param
			}
		case xml.EndElement:
			if t.Name.Local == "step" {
				return step, nil
			}
		}
	}
}

func decodeParameters(d *xml.Decoder) (map[string]*StepParameter, error) {
	params := make(map[string]*StepParameter)
	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "param" {
				name, param, err := decodeParameter(d, t)
				if err != nil {
					return nil, err
				}
				params[name] = param
			}
		case xml.EndElement:
			if t.Name.Local == "params" {
				return params, nil
			}
		}
	}
}

func decodeParameter(d *xml.Decoder, start xml.StartElement) (string, *StepParameter, error) {
	var name string
	param := &StepParameter{Type: StepParamTypeSimple}
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			name = attr.Value
		case "type":
			param.Type = attr.Value
		case "isRef":
			param.IsRef = attr.Value == "True" || attr.Value == "true"
		}
	}
	switch param.Type {
	case StepParamTypeContent:
		items, err := decodeContentItems(d)
		if err != nil {
			return name, nil, err
		}
		param.ContentValue = items
	case StepParamTypeGroup:
		group, err := decodeGroupParameter(d)
		if err != nil {
			return name, nil, err
		}
		param.GroupValue = group
	case StepParamTypeNumberSet:
		ns, err := decodeNumberSet(d)
		if err != nil {
			return name, nil, err
		}
		param.NumberSetValue = ns
	default:
		// Try to parse as content items first (for params without explicit type="content")
		items, simpleValue, err := decodeParamContent(d)
		if err != nil {
			return name, nil, err
		}
		if len(items) > 0 {
			param.Type = StepParamTypeContent
			param.ContentValue = items
		} else {
			param.SimpleValue = simpleValue
		}
	}
	return name, param, nil
}

// decodeParamContent reads param content, returning either content items or simple value
func decodeParamContent(d *xml.Decoder) ([]ContentItem, string, error) {
	var items []ContentItem
	var simpleValue strings.Builder
	for {
		token, err := d.Token()
		if err != nil {
			return nil, "", err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "item" {
				item, err := decodeContentItem(d, t)
				if err != nil {
					return nil, "", err
				}
				items = append(items, item)
			}
		case xml.CharData:
			simpleValue.Write(t)
		case xml.EndElement:
			if t.Name.Local == "param" {
				return items, simpleValue.String(), nil
			}
		}
	}
}

func decodeContentItems(d *xml.Decoder) ([]ContentItem, error) {
	var items []ContentItem
	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "item" {
				item, err := decodeContentItem(d, t)
				if err != nil {
					return nil, err
				}
				items = append(items, item)
			}
		case xml.EndElement:
			if t.Name.Local == "param" {
				return items, nil
			}
		}
	}
}

func decodeContentItem(d *xml.Decoder, start xml.StartElement) (ContentItem, error) {
	item := ContentItem{Type: ContentTypeText, WaitForFinish: true}
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "type":
			item.Type = attr.Value
		case "isRef":
			item.IsRef = attr.Value == "True" || attr.Value == "true"
		case "placement":
			item.Placement = attr.Value
		case "duration":
			dur, _ := parseDuration(attr.Value)
			item.Duration = dur
		case "waitForFinish":
			item.WaitForFinish = attr.Value != "False" && attr.Value != "false"
		}
	}
	var value string
	for {
		token, err := d.Token()
		if err != nil {
			return item, err
		}
		switch t := token.(type) {
		case xml.CharData:
			value += string(t)
		case xml.EndElement:
			if t.Name.Local == "item" {
				item.Value = value
				return item, nil
			}
		}
	}
}

func parseDuration(s string) (time.Duration, error) {
	var h, m, sec int
	_, err := fmt.Sscanf(s, "%d:%d:%d", &h, &m, &sec)
	if err != nil {
		return 0, err
	}
	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(sec)*time.Second, nil
}

func decodeGroupParameter(d *xml.Decoder) (map[string]*StepParameter, error) {
	group := make(map[string]*StepParameter)
	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "param" {
				name, param, err := decodeParameter(d, t)
				if err != nil {
					return nil, err
				}
				group[name] = param
			}
		case xml.EndElement:
			if t.Name.Local == "param" {
				return group, nil
			}
		}
	}
}

func decodeNumberSet(d *xml.Decoder) (*NumberSet, error) {
	ns := &NumberSet{}
	for {
		token, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "numberSet" {
				for _, attr := range t.Attr {
					switch attr.Name.Local {
					case "minimum":
						ns.Minimum, _ = strconv.Atoi(attr.Value)
					case "maximum":
						ns.Maximum, _ = strconv.Atoi(attr.Value)
					case "step":
						ns.Step, _ = strconv.Atoi(attr.Value)
					}
				}
			}
		case xml.EndElement:
			if t.Name.Local == "param" {
				return ns, nil
			}
		}
	}
}

func decodeSimpleValue(d *xml.Decoder) (string, error) {
	var value string
	for {
		token, err := d.Token()
		if err != nil {
			return "", err
		}
		switch t := token.(type) {
		case xml.CharData:
			value += string(t)
		case xml.EndElement:
			if t.Name.Local == "param" {
				return value, nil
			}
		}
	}
}
