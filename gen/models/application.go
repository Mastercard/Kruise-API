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

// Application application
// swagger:model application
type Application struct {

	// The environment to deploy to
	// Required: true
	// Min Length: 1
	// Enum: [Dev Stage Prod]
	Environment string `json:"environment"`

	// ingress
	Ingress *Ingress `json:"ingress,omitempty"`

	// The name of the application
	// Required: true
	// Min Length: 1
	Name string `json:"name"`

	// The namespace to deploy to
	// Required: true
	// Min Length: 1
	Namespace string `json:"namespace"`

	// The relative path to the manifests in the git repo
	// Min Length: 1
	Path string `json:"path,omitempty"`

	// The region to deploy to
	// Required: true
	// Min Length: 1
	// Enum: [STL KCI BEL]
	Region string `json:"region"`

	// The version or release name of the application
	// Required: true
	// Min Length: 1
	Release string `json:"release"`

	// The git repo URL
	// Required: true
	// Min Length: 1
	// Format: uri
	RepoURL strfmt.URI `json:"repoURL"`

	// services
	Services []*Service `json:"services"`

	// Defines the commit, tag, or branch in which to sync the application to.
	// Min Length: 1
	TargetRevision string `json:"targetRevision,omitempty"`

	// The name of the tenant
	// Required: true
	// Min Length: 1
	Tenant string `json:"tenant"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationTypeEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Dev","Stage","Prod"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationTypeEnvironmentPropEnum = append(applicationTypeEnvironmentPropEnum, v)
	}
}

const (

	// ApplicationEnvironmentDev captures enum value "Dev"
	ApplicationEnvironmentDev string = "Dev"

	// ApplicationEnvironmentStage captures enum value "Stage"
	ApplicationEnvironmentStage string = "Stage"

	// ApplicationEnvironmentProd captures enum value "Prod"
	ApplicationEnvironmentProd string = "Prod"
)

// prop value enum
func (m *Application) validateEnvironmentEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationTypeEnvironmentPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Application) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.RequiredString("environment", "body", string(m.Environment)); err != nil {
		return err
	}

	if err := validate.MinLength("environment", "body", string(m.Environment), 1); err != nil {
		return err
	}

	// value enum
	if err := m.validateEnvironmentEnum("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateIngress(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingress) { // not required
		return nil
	}

	if m.Ingress != nil {
		if err := m.Ingress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingress")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateNamespace(formats strfmt.Registry) error {

	if err := validate.RequiredString("namespace", "body", string(m.Namespace)); err != nil {
		return err
	}

	if err := validate.MinLength("namespace", "body", string(m.Namespace), 1); err != nil {
		return err
	}

	return nil
}

func (m *Application) validatePath(formats strfmt.Registry) error {

	if swag.IsZero(m.Path) { // not required
		return nil
	}

	if err := validate.MinLength("path", "body", string(m.Path), 1); err != nil {
		return err
	}

	return nil
}

var applicationTypeRegionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STL","KCI","BEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationTypeRegionPropEnum = append(applicationTypeRegionPropEnum, v)
	}
}

const (

	// ApplicationRegionSTL captures enum value "STL"
	ApplicationRegionSTL string = "STL"

	// ApplicationRegionKCI captures enum value "KCI"
	ApplicationRegionKCI string = "KCI"

	// ApplicationRegionBEL captures enum value "BEL"
	ApplicationRegionBEL string = "BEL"
)

// prop value enum
func (m *Application) validateRegionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applicationTypeRegionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Application) validateRegion(formats strfmt.Registry) error {

	if err := validate.RequiredString("region", "body", string(m.Region)); err != nil {
		return err
	}

	if err := validate.MinLength("region", "body", string(m.Region), 1); err != nil {
		return err
	}

	// value enum
	if err := m.validateRegionEnum("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateRelease(formats strfmt.Registry) error {

	if err := validate.RequiredString("release", "body", string(m.Release)); err != nil {
		return err
	}

	if err := validate.MinLength("release", "body", string(m.Release), 1); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateRepoURL(formats strfmt.Registry) error {

	if err := validate.Required("repoURL", "body", strfmt.URI(m.RepoURL)); err != nil {
		return err
	}

	if err := validate.MinLength("repoURL", "body", string(m.RepoURL), 1); err != nil {
		return err
	}

	if err := validate.FormatOf("repoURL", "body", "uri", m.RepoURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateTargetRevision(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetRevision) { // not required
		return nil
	}

	if err := validate.MinLength("targetRevision", "body", string(m.TargetRevision), 1); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateTenant(formats strfmt.Registry) error {

	if err := validate.RequiredString("tenant", "body", string(m.Tenant)); err != nil {
		return err
	}

	if err := validate.MinLength("tenant", "body", string(m.Tenant), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
