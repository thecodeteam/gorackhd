package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Node20PartialNode Post a node into RackHD

swagger:model node.2.0_PartialNode
*/
type Node20PartialNode struct {

	/* auto discover
	 */
	AutoDiscover string `json:"autoDiscover,omitempty"`

	/* boot settings
	 */
	BootSettings interface{} `json:"bootSettings,omitempty"`

	/* identifiers
	 */
	Identifiers []string `json:"identifiers,omitempty"`

	/* Name of the node
	 */
	Name string `json:"name,omitempty"`

	/* The list of obm settings
	 */
	Obms []*NodesPostObmByID `json:"obms,omitempty"`

	/* relations
	 */
	Relations []*RelationsObj `json:"relations,omitempty"`

	/* tags
	 */
	Tags []string `json:"tags,omitempty"`

	/* Type of node
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this node 2 0 partial node
func (m *Node20PartialNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifiers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateObms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRelations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node20PartialNode) validateIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {

	}

	return nil
}

func (m *Node20PartialNode) validateObms(formats strfmt.Registry) error {

	if swag.IsZero(m.Obms) { // not required
		return nil
	}

	for i := 0; i < len(m.Obms); i++ {

		if swag.IsZero(m.Obms[i]) { // not required
			continue
		}

		if m.Obms[i] != nil {

			if err := m.Obms[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Node20PartialNode) validateRelations(formats strfmt.Registry) error {

	if swag.IsZero(m.Relations) { // not required
		return nil
	}

	for i := 0; i < len(m.Relations); i++ {

		if swag.IsZero(m.Relations[i]) { // not required
			continue
		}

		if m.Relations[i] != nil {

			if err := m.Relations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Node20PartialNode) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

	}

	return nil
}

var node20PartialNodeTypeTypePropEnum []interface{}

// prop value enum
func (m *Node20PartialNode) validateTypeEnum(path, location string, value string) error {
	if node20PartialNodeTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["compute","compute-container","switch","dae","pdu","mgmt","enclosure","rack"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			node20PartialNodeTypeTypePropEnum = append(node20PartialNodeTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, node20PartialNodeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Node20PartialNode) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
