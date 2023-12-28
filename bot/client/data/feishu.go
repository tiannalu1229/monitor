package data

type Card struct {
	Header   Header        `json:"header"`
	Elements []CardElement `json:"elements"`
}

// Header header结构
type Header struct {
	Template string `json:"template"`
	Title    Title  `json:"title"`
}

// Title header中title 填写content
type Title struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

// CardElement 属于card
type CardElement struct {
	Tag             string   `json:"tag"`
	FlexMode        string   `json:"flex_mode"`
	BackgroundStyle string   `json:"background_style"`
	Columns         []Column `json:"columns"`
}

// Column Weight
type Column struct {
	Tag           string              `json:"tag"`
	Width         string              `json:"width"`
	Weight        int                 `json:"weight"`
	VerticalAlign string              `json:"vertical_align"`
	Elements      []ColumnTextElement `json:"elements"`
}

type ColumnTextElement struct {
	Tag  string `json:"tag"`
	Text Text   `json:"text,omitempty"`
}

// Text 填写Content
type Text struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}
