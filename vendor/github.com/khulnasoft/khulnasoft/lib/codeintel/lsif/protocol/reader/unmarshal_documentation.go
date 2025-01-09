package reader

import "github.com/khulnasoft/khulnasoft/lib/codeintel/lsif/protocol"

// This file contains code for the Khulnasoft documentation LSIF extension.

func init() {
	// Vertex unmarshalers
	vertexUnmarshalers[string(protocol.VertexKhulnasoftDocumentationResult)] = unmarshalDocumentationResult
	vertexUnmarshalers[string(protocol.VertexKhulnasoftDocumentationString)] = unmarshalDocumentationString

	// Edge unmarshalers
	edgeUnmarshalers[string(protocol.EdgeKhulnasoftDocumentationString)] = unmarshalDocumentationStringEdge
}

func unmarshalDocumentationResult(line []byte) (any, error) {
	var payload struct {
		Result protocol.Documentation `json:"result"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}
	return payload.Result, nil
}

func unmarshalDocumentationString(line []byte) (any, error) {
	var payload struct {
		Result protocol.MarkupContent `json:"result"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return nil, err
	}
	return payload.Result, nil
}

type DocumentationStringEdge struct {
	OutV int
	InV  int
	Kind protocol.DocumentationStringKind
}

func unmarshalDocumentationStringEdge(line []byte) (any, error) {
	var payload struct {
		OutV int    `json:"outV"`
		InV  int    `json:"inV"`
		Kind string `json:"kind"`
	}
	if err := unmarshaller.Unmarshal(line, &payload); err != nil {
		return DocumentationStringEdge{}, err
	}
	return DocumentationStringEdge{
		OutV: payload.OutV,
		InV:  payload.InV,
		Kind: protocol.DocumentationStringKind(payload.Kind),
	}, nil
}
