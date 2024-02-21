// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/17ac39c7f9266bc303baa029f90194aecb1c3b7c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IndexTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/17ac39c7f9266bc303baa029f90194aecb1c3b7c/specification/indices/_types/IndexTemplate.ts#L31-L70
type IndexTemplate struct {
	AllowAutoCreate *bool `json:"allow_auto_create,omitempty"`
	// ComposedOf An ordered list of component template names.
	// Component templates are merged in the order specified, meaning that the last
	// component template specified has the highest precedence.
	ComposedOf []string `json:"composed_of"`
	// DataStream If this object is included, the template is used to create data streams and
	// their backing indices.
	// Supports an empty object.
	// Data streams require a matching index template with a `data_stream` object.
	DataStream *IndexTemplateDataStreamConfiguration `json:"data_stream,omitempty"`
	// IndexPatterns Name of the index template.
	IndexPatterns []string `json:"index_patterns"`
	// Meta_ Optional user metadata about the index template. May have any contents.
	// This map is not automatically generated by Elasticsearch.
	Meta_ Metadata `json:"_meta,omitempty"`
	// Priority Priority to determine index template precedence when a new data stream or
	// index is created.
	// The index template with the highest priority is chosen.
	// If no priority is specified the template is treated as though it is of
	// priority 0 (lowest priority).
	// This number is not automatically generated by Elasticsearch.
	Priority *int64 `json:"priority,omitempty"`
	// Template Template to be applied.
	// It may optionally include an `aliases`, `mappings`, or `settings`
	// configuration.
	Template *IndexTemplateSummary `json:"template,omitempty"`
	// Version Version number used to manage index templates externally.
	// This number is not automatically generated by Elasticsearch.
	Version *int64 `json:"version,omitempty"`
}

func (s *IndexTemplate) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "allow_auto_create":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AllowAutoCreate = &value
			case bool:
				s.AllowAutoCreate = &v
			}

		case "composed_of":
			if err := dec.Decode(&s.ComposedOf); err != nil {
				return err
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return err
			}

		case "index_patterns":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.IndexPatterns = append(s.IndexPatterns, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.IndexPatterns); err != nil {
					return err
				}
			}

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "priority":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Priority = &value
			case float64:
				f := int64(v)
				s.Priority = &f
			}

		case "template":
			if err := dec.Decode(&s.Template); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndexTemplate returns a IndexTemplate.
func NewIndexTemplate() *IndexTemplate {
	r := &IndexTemplate{}

	return r
}
