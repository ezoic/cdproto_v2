package overlay

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/ezoic/cdproto_v2/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// HighlightConfig configuration data for the highlighting of page elements.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#type-HighlightConfig
type HighlightConfig struct {
	ShowInfo           bool      `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowStyles         bool      `json:"showStyles,omitempty"`         // Whether the node styles in the tooltip (default: false).
	ShowRulers         bool      `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool      `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	ContentColor       *cdp.RGBA `json:"contentColor,omitempty"`       // The content box highlight fill color (default: transparent).
	PaddingColor       *cdp.RGBA `json:"paddingColor,omitempty"`       // The padding highlight fill color (default: transparent).
	BorderColor        *cdp.RGBA `json:"borderColor,omitempty"`        // The border highlight fill color (default: transparent).
	MarginColor        *cdp.RGBA `json:"marginColor,omitempty"`        // The margin highlight fill color (default: transparent).
	EventTargetColor   *cdp.RGBA `json:"eventTargetColor,omitempty"`   // The event target element highlight fill color (default: transparent).
	ShapeColor         *cdp.RGBA `json:"shapeColor,omitempty"`         // The shape outside fill color (default: transparent).
	ShapeMarginColor   *cdp.RGBA `json:"shapeMarginColor,omitempty"`   // The shape margin fill color (default: transparent).
	CSSGridColor       *cdp.RGBA `json:"cssGridColor,omitempty"`       // The grid layout color (default: transparent).
}

// InspectMode [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Overlay#type-InspectMode
type InspectMode string

// String returns the InspectMode as string value.
func (t InspectMode) String() string {
	return string(t)
}

// InspectMode values.
const (
	InspectModeSearchForNode         InspectMode = "searchForNode"
	InspectModeSearchForUAShadowDOM  InspectMode = "searchForUAShadowDOM"
	InspectModeCaptureAreaScreenshot InspectMode = "captureAreaScreenshot"
	InspectModeShowDistances         InspectMode = "showDistances"
	InspectModeNone                  InspectMode = "none"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InspectMode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InspectMode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InspectMode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InspectMode(in.String()) {
	case InspectModeSearchForNode:
		*t = InspectModeSearchForNode
	case InspectModeSearchForUAShadowDOM:
		*t = InspectModeSearchForUAShadowDOM
	case InspectModeCaptureAreaScreenshot:
		*t = InspectModeCaptureAreaScreenshot
	case InspectModeShowDistances:
		*t = InspectModeShowDistances
	case InspectModeNone:
		*t = InspectModeNone

	default:
		in.AddError(errors.New("unknown InspectMode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InspectMode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
