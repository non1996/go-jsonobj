package jsonobj

import jsoniter "github.com/json-iterator/go"

var jsonAPI = jsoniter.Config{
	EscapeHTML:             true,
	UseNumber:              true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()
