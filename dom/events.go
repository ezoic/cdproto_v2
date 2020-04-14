package dom

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/ezoic/cdproto_v2/cdp"
)

// EventAttributeModified fired when Element's attribute is modified.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-attributeModified
type EventAttributeModified struct {
	NodeID cdp.NodeID `json:"nodeId"` // Id of the node that has changed.
	Name   string     `json:"name"`   // Attribute name.
	Value  string     `json:"value"`  // Attribute value.
}

// EventAttributeRemoved fired when Element's attribute is removed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-attributeRemoved
type EventAttributeRemoved struct {
	NodeID cdp.NodeID `json:"nodeId"` // Id of the node that has changed.
	Name   string     `json:"name"`   // A ttribute name.
}

// EventCharacterDataModified mirrors DOMCharacterDataModified event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-characterDataModified
type EventCharacterDataModified struct {
	NodeID        cdp.NodeID `json:"nodeId"`        // Id of the node that has changed.
	CharacterData string     `json:"characterData"` // New text value.
}

// EventChildNodeCountUpdated fired when Container's child node count has
// changed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-childNodeCountUpdated
type EventChildNodeCountUpdated struct {
	NodeID         cdp.NodeID `json:"nodeId"`         // Id of the node that has changed.
	ChildNodeCount int64      `json:"childNodeCount"` // New node count.
}

// EventChildNodeInserted mirrors DOMNodeInserted event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-childNodeInserted
type EventChildNodeInserted struct {
	ParentNodeID   cdp.NodeID `json:"parentNodeId"`   // Id of the node that has changed.
	PreviousNodeID cdp.NodeID `json:"previousNodeId"` // If of the previous siblint.
	Node           *cdp.Node  `json:"node"`           // Inserted node data.
}

// EventChildNodeRemoved mirrors DOMNodeRemoved event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-childNodeRemoved
type EventChildNodeRemoved struct {
	ParentNodeID cdp.NodeID `json:"parentNodeId"` // Parent id.
	NodeID       cdp.NodeID `json:"nodeId"`       // Id of the node that has been removed.
}

// EventDistributedNodesUpdated called when distribution is changed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-distributedNodesUpdated
type EventDistributedNodesUpdated struct {
	InsertionPointID cdp.NodeID         `json:"insertionPointId"` // Insertion point where distributed nodes were updated.
	DistributedNodes []*cdp.BackendNode `json:"distributedNodes"` // Distributed nodes for given insertion point.
}

// EventDocumentUpdated fired when Document has been totally updated. Node
// ids are no longer valid.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-documentUpdated
type EventDocumentUpdated struct{}

// EventInlineStyleInvalidated fired when Element's inline style is modified
// via a CSS property modification.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-inlineStyleInvalidated
type EventInlineStyleInvalidated struct {
	NodeIds []cdp.NodeID `json:"nodeIds"` // Ids of the nodes for which the inline styles have been invalidated.
}

// EventPseudoElementAdded called when a pseudo element is added to an
// element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-pseudoElementAdded
type EventPseudoElementAdded struct {
	ParentID      cdp.NodeID `json:"parentId"`      // Pseudo element's parent element id.
	PseudoElement *cdp.Node  `json:"pseudoElement"` // The added pseudo element.
}

// EventPseudoElementRemoved called when a pseudo element is removed from an
// element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-pseudoElementRemoved
type EventPseudoElementRemoved struct {
	ParentID        cdp.NodeID `json:"parentId"`        // Pseudo element's parent element id.
	PseudoElementID cdp.NodeID `json:"pseudoElementId"` // The removed pseudo element id.
}

// EventSetChildNodes fired when backend wants to provide client with the
// missing DOM structure. This happens upon most of the calls requesting node
// ids.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-setChildNodes
type EventSetChildNodes struct {
	ParentID cdp.NodeID  `json:"parentId"` // Parent node id to populate with children.
	Nodes    []*cdp.Node `json:"nodes"`    // Child nodes array.
}

// EventShadowRootPopped called when shadow root is popped from the element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-shadowRootPopped
type EventShadowRootPopped struct {
	HostID cdp.NodeID `json:"hostId"` // Host element id.
	RootID cdp.NodeID `json:"rootId"` // Shadow root id.
}

// EventShadowRootPushed called when shadow root is pushed into the element.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#event-shadowRootPushed
type EventShadowRootPushed struct {
	HostID cdp.NodeID `json:"hostId"` // Host element id.
	Root   *cdp.Node  `json:"root"`   // Shadow root.
}
