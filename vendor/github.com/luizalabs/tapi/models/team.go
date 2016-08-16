package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Team team

swagger:model Team
*/
type Team struct {

	/* apps
	 */
	Apps []*App `json:"apps,omitempty"`

	/* email
	 */
	Email strfmt.Email `json:"email,omitempty"`

	/* i am member
	 */
	IAmMember bool `json:"iAmMember,omitempty"`

	/* id
	 */
	ID int64 `json:"id,omitempty"`

	/* members
	 */
	Members []*User `json:"members,omitempty"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* url
	 */
	URL string `json:"url,omitempty"`
}

// Validate validates this team
func (m *Team) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Team) validateApps(formats strfmt.Registry) error {

	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	for i := 0; i < len(m.Apps); i++ {

		if swag.IsZero(m.Apps[i]) { // not required
			continue
		}

		if m.Apps[i] != nil {

			if err := m.Apps[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Team) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {

		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {

			if err := m.Members[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Team) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
