package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Event Event

swagger:model Event
*/
type Event struct {

	/* autosearch
	 */
	Autosearch bool `json:"Autosearch,omitempty"`

	/* created

	Read Only: true
	*/
	Created int64 `json:"Created,omitempty"`

	/* deleted

	Read Only: true
	*/
	Deleted *bool `json:"Deleted,omitempty"`

	/* end date
	 */
	EndDate strfmt.DateTime `json:"EndDate,omitempty"`

	/* Id

	Read Only: true
	*/
	ID int64 `json:"Id,omitempty"`

	/* latitude
	 */
	Latitude float32 `json:"Latitude,omitempty"`

	/* longitude
	 */
	Longitude float32 `json:"Longitude,omitempty"`

	/* name
	 */
	Name string `json:"Name,omitempty"`

	/* notes
	 */
	Notes string `json:"Notes,omitempty"`

	/* start date
	 */
	StartDate strfmt.DateTime `json:"StartDate,omitempty"`

	/* URL
	 */
	URL string `json:"URL,omitempty"`

	/* updated

	Read Only: true
	*/
	Updated int64 `json:"Updated,omitempty"`

	/* version

	Read Only: true
	*/
	Version int64 `json:"Version,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	return nil
}
