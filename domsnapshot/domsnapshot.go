// Package domsnapshot provides the Chrome DevTools Protocol
// commands, types, and events for the DOMSnapshot domain.
//
// This domain facilitates obtaining document snapshots with DOM, layout, and
// style information.
//
// Generated by the cdproto-gen command.
package domsnapshot

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/ezoic/cdproto_v2/cdp"
)

// DisableParams disables DOM snapshot agent for the given page.
type DisableParams struct{}

// Disable disables DOM snapshot agent for the given page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes DOMSnapshot.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables DOM snapshot agent for the given page.
type EnableParams struct{}

// Enable enables DOM snapshot agent for the given page.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes DOMSnapshot.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// CaptureSnapshotParams returns a document snapshot, including the full DOM
// tree of the root node (including iframes, template contents, and imported
// documents) in a flattened array, as well as layout and white-listed computed
// style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
type CaptureSnapshotParams struct {
	ComputedStyles    []string `json:"computedStyles"`              // Whitelist of computed styles to return.
	IncludePaintOrder bool     `json:"includePaintOrder,omitempty"` // Whether to include layout object paint orders into the snapshot.
	IncludeDOMRects   bool     `json:"includeDOMRects,omitempty"`   // Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
}

// CaptureSnapshot returns a document snapshot, including the full DOM tree
// of the root node (including iframes, template contents, and imported
// documents) in a flattened array, as well as layout and white-listed computed
// style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#method-captureSnapshot
//
// parameters:
//   computedStyles - Whitelist of computed styles to return.
func CaptureSnapshot(computedStyles []string) *CaptureSnapshotParams {
	return &CaptureSnapshotParams{
		ComputedStyles: computedStyles,
	}
}

// WithIncludePaintOrder whether to include layout object paint orders into
// the snapshot.
func (p CaptureSnapshotParams) WithIncludePaintOrder(includePaintOrder bool) *CaptureSnapshotParams {
	p.IncludePaintOrder = includePaintOrder
	return &p
}

// WithIncludeDOMRects whether to include DOM rectangles (offsetRects,
// clientRects, scrollRects) into the snapshot.
func (p CaptureSnapshotParams) WithIncludeDOMRects(includeDOMRects bool) *CaptureSnapshotParams {
	p.IncludeDOMRects = includeDOMRects
	return &p
}

// CaptureSnapshotReturns return values.
type CaptureSnapshotReturns struct {
	Documents []*DocumentSnapshot `json:"documents,omitempty"` // The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Strings   []string            `json:"strings,omitempty"`   // Shared string table that all string properties refer to with indexes.
}

// Do executes DOMSnapshot.captureSnapshot against the provided context.
//
// returns:
//   documents - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
//   strings - Shared string table that all string properties refer to with indexes.
func (p *CaptureSnapshotParams) Do(ctx context.Context) (documents []*DocumentSnapshot, strings []string, err error) {
	// execute
	var res CaptureSnapshotReturns
	err = cdp.Execute(ctx, CommandCaptureSnapshot, p, &res)
	if err != nil {
		return nil, nil, err
	}

	return res.Documents, res.Strings, nil
}

// Command names.
const (
	CommandDisable         = "DOMSnapshot.disable"
	CommandEnable          = "DOMSnapshot.enable"
	CommandCaptureSnapshot = "DOMSnapshot.captureSnapshot"
)
