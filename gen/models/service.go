// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Service service
// swagger:model service
type Service struct {

	// The name of the service
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// ports
	// Required: true
	Ports []*ServicePort `json:"ports"`

	// The service type
	// Required: true
	// Enum: [ClusterIP ExternalName LoadBalancer]
	Type string `json:"type"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Service) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var serviceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ClusterIP","ExternalName","LoadBalancer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceTypeTypePropEnum = append(serviceTypeTypePropEnum, v)
	}
}

const (

	// ServiceTypeClusterIP captures enum value "ClusterIP"
	ServiceTypeClusterIP string = "ClusterIP"

	// ServiceTypeExternalName captures enum value "ExternalName"
	ServiceTypeExternalName string = "ExternalName"

	// ServiceTypeLoadBalancer captures enum value "LoadBalancer"
	ServiceTypeLoadBalancer string = "LoadBalancer"
)

// prop value enum
func (m *Service) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Service) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
