// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people/ent/person"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// DateUpdated holds the value of the "date_updated" field.
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// ObjectClass holds the value of the "object_class" field.
	ObjectClass string `json:"object_class,omitempty"`
	// UgentUsername holds the value of the "ugent_username" field.
	UgentUsername string `json:"ugent_username,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// UgentID holds the value of the "ugent_id" field.
	UgentID []string `json:"ugent_id,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate string `json:"birth_date,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender person.Gender `json:"gender,omitempty"`
	// Nationality holds the value of the "nationality" field.
	Nationality string `json:"nationality,omitempty"`
	// UgentBarcode holds the value of the "ugent_barcode" field.
	UgentBarcode []string `json:"ugent_barcode,omitempty"`
	// UgentJobCategory holds the value of the "ugent_job_category" field.
	UgentJobCategory []string `json:"ugent_job_category,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// UgentTel holds the value of the "ugent_tel" field.
	UgentTel string `json:"ugent_tel,omitempty"`
	// UgentCampus holds the value of the "ugent_campus" field.
	UgentCampus []string `json:"ugent_campus,omitempty"`
	// UgentDepartmentID holds the value of the "ugent_department_id" field.
	UgentDepartmentID []string `json:"ugent_department_id,omitempty"`
	// UgentFacultyID holds the value of the "ugent_faculty_id" field.
	UgentFacultyID []string `json:"ugent_faculty_id,omitempty"`
	// UgentJobTitle holds the value of the "ugent_job_title" field.
	UgentJobTitle []string `json:"ugent_job_title,omitempty"`
	// UgentStreetAddress holds the value of the "ugent_street_address" field.
	UgentStreetAddress string `json:"ugent_street_address,omitempty"`
	// UgentPostalCode holds the value of the "ugent_postal_code" field.
	UgentPostalCode string `json:"ugent_postal_code,omitempty"`
	// UgentLocality holds the value of the "ugent_locality" field.
	UgentLocality string `json:"ugent_locality,omitempty"`
	// UgentLastEnrolled holds the value of the "ugent_last_enrolled" field.
	UgentLastEnrolled string `json:"ugent_last_enrolled,omitempty"`
	// HomeStreetAddress holds the value of the "home_street_address" field.
	HomeStreetAddress string `json:"home_street_address,omitempty"`
	// HomePostalCode holds the value of the "home_postal_code" field.
	HomePostalCode string `json:"home_postal_code,omitempty"`
	// HomeLocality holds the value of the "home_locality" field.
	HomeLocality string `json:"home_locality,omitempty"`
	// HomeCountry holds the value of the "home_country" field.
	HomeCountry string `json:"home_country,omitempty"`
	// HomeTel holds the value of the "home_tel" field.
	HomeTel string `json:"home_tel,omitempty"`
	// DormStreetAddress holds the value of the "dorm_street_address" field.
	DormStreetAddress string `json:"dorm_street_address,omitempty"`
	// DormPostalCode holds the value of the "dorm_postal_code" field.
	DormPostalCode string `json:"dorm_postal_code,omitempty"`
	// DormLocality holds the value of the "dorm_locality" field.
	DormLocality string `json:"dorm_locality,omitempty"`
	// DormCountry holds the value of the "dorm_country" field.
	DormCountry string `json:"dorm_country,omitempty"`
	// ResearchDiscipline holds the value of the "research_discipline" field.
	ResearchDiscipline []string `json:"research_discipline,omitempty"`
	// ResearchDisciplineCode holds the value of the "research_discipline_code" field.
	ResearchDisciplineCode []string `json:"research_discipline_code,omitempty"`
	// UgentExpirationDate holds the value of the "ugent_expiration_date" field.
	UgentExpirationDate string `json:"ugent_expiration_date,omitempty"`
	// UzgentJobTitle holds the value of the "uzgent_job_title" field.
	UzgentJobTitle []string `json:"uzgent_job_title,omitempty"`
	// UzgentDepartmentName holds the value of the "uzgent_department_name" field.
	UzgentDepartmentName []string `json:"uzgent_department_name,omitempty"`
	// UzgentID holds the value of the "uzgent_id" field.
	UzgentID []string `json:"uzgent_id,omitempty"`
	// UgentExtCategory holds the value of the "ugent_ext_category" field.
	UgentExtCategory []string `json:"ugent_ext_category,omitempty"`
	// UgentAppointmentDate holds the value of the "ugent_appointment_date" field.
	UgentAppointmentDate string `json:"ugent_appointment_date,omitempty"`
	// UgentDepartmentName holds the value of the "ugent_department_name" field.
	UgentDepartmentName []string `json:"ugent_department_name,omitempty"`
	// OrcidBio holds the value of the "orcid_bio" field.
	OrcidBio string `json:"orcid_bio,omitempty"`
	// OrcidID holds the value of the "orcid_id" field.
	OrcidID string `json:"orcid_id,omitempty"`
	// OrcidSettings holds the value of the "orcid_settings" field.
	OrcidSettings map[string]interface{} `json:"orcid_settings,omitempty"`
	// OrcidToken holds the value of the "orcid_token" field.
	OrcidToken string `json:"orcid_token,omitempty"`
	// OrcidVerify holds the value of the "orcid_verify" field.
	OrcidVerify string `json:"orcid_verify,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// Roles holds the value of the "roles" field.
	Roles []string `json:"roles,omitempty"`
	// PublicationCount holds the value of the "publication_count" field.
	PublicationCount int `json:"publication_count,omitempty"`
	// UgentMemorialisID holds the value of the "ugent_memorialis_id" field.
	UgentMemorialisID string `json:"ugent_memorialis_id,omitempty"`
	// PreferredFirstName holds the value of the "preferred_first_name" field.
	PreferredFirstName string `json:"preferred_first_name,omitempty"`
	// PreferredLastName holds the value of the "preferred_last_name" field.
	PreferredLastName string `json:"preferred_last_name,omitempty"`
	// Replaces holds the value of the "replaces" field.
	Replaces []map[string]string `json:"replaces,omitempty"`
	// ReplacedBy holds the value of the "replaced_by" field.
	ReplacedBy []map[string]string `json:"replaced_by,omitempty"`
	// DateLastLogin holds the value of the "date_last_login" field.
	DateLastLogin time.Time `json:"date_last_login,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldUgentID, person.FieldUgentBarcode, person.FieldUgentJobCategory, person.FieldUgentCampus, person.FieldUgentDepartmentID, person.FieldUgentFacultyID, person.FieldUgentJobTitle, person.FieldResearchDiscipline, person.FieldResearchDisciplineCode, person.FieldUzgentJobTitle, person.FieldUzgentDepartmentName, person.FieldUzgentID, person.FieldUgentExtCategory, person.FieldUgentDepartmentName, person.FieldOrcidSettings, person.FieldSettings, person.FieldRoles, person.FieldReplaces, person.FieldReplacedBy:
			values[i] = new([]byte)
		case person.FieldActive, person.FieldDeleted:
			values[i] = new(sql.NullBool)
		case person.FieldPublicationCount:
			values[i] = new(sql.NullInt64)
		case person.FieldID, person.FieldObjectClass, person.FieldUgentUsername, person.FieldFirstName, person.FieldMiddleName, person.FieldLastName, person.FieldBirthDate, person.FieldEmail, person.FieldGender, person.FieldNationality, person.FieldTitle, person.FieldUgentTel, person.FieldUgentStreetAddress, person.FieldUgentPostalCode, person.FieldUgentLocality, person.FieldUgentLastEnrolled, person.FieldHomeStreetAddress, person.FieldHomePostalCode, person.FieldHomeLocality, person.FieldHomeCountry, person.FieldHomeTel, person.FieldDormStreetAddress, person.FieldDormPostalCode, person.FieldDormLocality, person.FieldDormCountry, person.FieldUgentExpirationDate, person.FieldUgentAppointmentDate, person.FieldOrcidBio, person.FieldOrcidID, person.FieldOrcidToken, person.FieldOrcidVerify, person.FieldUgentMemorialisID, person.FieldPreferredFirstName, person.FieldPreferredLastName:
			values[i] = new(sql.NullString)
		case person.FieldDateCreated, person.FieldDateUpdated, person.FieldDateLastLogin:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Person", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pe.ID = value.String
			}
		case person.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				pe.DateCreated = value.Time
			}
		case person.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_updated", values[i])
			} else if value.Valid {
				pe.DateUpdated = value.Time
			}
		case person.FieldObjectClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field object_class", values[i])
			} else if value.Valid {
				pe.ObjectClass = value.String
			}
		case person.FieldUgentUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_username", values[i])
			} else if value.Valid {
				pe.UgentUsername = value.String
			}
		case person.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				pe.FirstName = value.String
			}
		case person.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				pe.MiddleName = value.String
			}
		case person.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pe.LastName = value.String
			}
		case person.FieldUgentID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentID); err != nil {
					return fmt.Errorf("unmarshal field ugent_id: %w", err)
				}
			}
		case person.FieldBirthDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				pe.BirthDate = value.String
			}
		case person.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pe.Email = value.String
			}
		case person.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pe.Gender = person.Gender(value.String)
			}
		case person.FieldNationality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nationality", values[i])
			} else if value.Valid {
				pe.Nationality = value.String
			}
		case person.FieldUgentBarcode:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_barcode", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentBarcode); err != nil {
					return fmt.Errorf("unmarshal field ugent_barcode: %w", err)
				}
			}
		case person.FieldUgentJobCategory:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_job_category", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentJobCategory); err != nil {
					return fmt.Errorf("unmarshal field ugent_job_category: %w", err)
				}
			}
		case person.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pe.Title = value.String
			}
		case person.FieldUgentTel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_tel", values[i])
			} else if value.Valid {
				pe.UgentTel = value.String
			}
		case person.FieldUgentCampus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_campus", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentCampus); err != nil {
					return fmt.Errorf("unmarshal field ugent_campus: %w", err)
				}
			}
		case person.FieldUgentDepartmentID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_department_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentDepartmentID); err != nil {
					return fmt.Errorf("unmarshal field ugent_department_id: %w", err)
				}
			}
		case person.FieldUgentFacultyID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_faculty_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentFacultyID); err != nil {
					return fmt.Errorf("unmarshal field ugent_faculty_id: %w", err)
				}
			}
		case person.FieldUgentJobTitle:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_job_title", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentJobTitle); err != nil {
					return fmt.Errorf("unmarshal field ugent_job_title: %w", err)
				}
			}
		case person.FieldUgentStreetAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_street_address", values[i])
			} else if value.Valid {
				pe.UgentStreetAddress = value.String
			}
		case person.FieldUgentPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_postal_code", values[i])
			} else if value.Valid {
				pe.UgentPostalCode = value.String
			}
		case person.FieldUgentLocality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_locality", values[i])
			} else if value.Valid {
				pe.UgentLocality = value.String
			}
		case person.FieldUgentLastEnrolled:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_last_enrolled", values[i])
			} else if value.Valid {
				pe.UgentLastEnrolled = value.String
			}
		case person.FieldHomeStreetAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_street_address", values[i])
			} else if value.Valid {
				pe.HomeStreetAddress = value.String
			}
		case person.FieldHomePostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_postal_code", values[i])
			} else if value.Valid {
				pe.HomePostalCode = value.String
			}
		case person.FieldHomeLocality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_locality", values[i])
			} else if value.Valid {
				pe.HomeLocality = value.String
			}
		case person.FieldHomeCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_country", values[i])
			} else if value.Valid {
				pe.HomeCountry = value.String
			}
		case person.FieldHomeTel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_tel", values[i])
			} else if value.Valid {
				pe.HomeTel = value.String
			}
		case person.FieldDormStreetAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dorm_street_address", values[i])
			} else if value.Valid {
				pe.DormStreetAddress = value.String
			}
		case person.FieldDormPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dorm_postal_code", values[i])
			} else if value.Valid {
				pe.DormPostalCode = value.String
			}
		case person.FieldDormLocality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dorm_locality", values[i])
			} else if value.Valid {
				pe.DormLocality = value.String
			}
		case person.FieldDormCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dorm_country", values[i])
			} else if value.Valid {
				pe.DormCountry = value.String
			}
		case person.FieldResearchDiscipline:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field research_discipline", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ResearchDiscipline); err != nil {
					return fmt.Errorf("unmarshal field research_discipline: %w", err)
				}
			}
		case person.FieldResearchDisciplineCode:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field research_discipline_code", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ResearchDisciplineCode); err != nil {
					return fmt.Errorf("unmarshal field research_discipline_code: %w", err)
				}
			}
		case person.FieldUgentExpirationDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_expiration_date", values[i])
			} else if value.Valid {
				pe.UgentExpirationDate = value.String
			}
		case person.FieldUzgentJobTitle:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uzgent_job_title", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UzgentJobTitle); err != nil {
					return fmt.Errorf("unmarshal field uzgent_job_title: %w", err)
				}
			}
		case person.FieldUzgentDepartmentName:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uzgent_department_name", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UzgentDepartmentName); err != nil {
					return fmt.Errorf("unmarshal field uzgent_department_name: %w", err)
				}
			}
		case person.FieldUzgentID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uzgent_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UzgentID); err != nil {
					return fmt.Errorf("unmarshal field uzgent_id: %w", err)
				}
			}
		case person.FieldUgentExtCategory:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_ext_category", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentExtCategory); err != nil {
					return fmt.Errorf("unmarshal field ugent_ext_category: %w", err)
				}
			}
		case person.FieldUgentAppointmentDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_appointment_date", values[i])
			} else if value.Valid {
				pe.UgentAppointmentDate = value.String
			}
		case person.FieldUgentDepartmentName:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_department_name", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.UgentDepartmentName); err != nil {
					return fmt.Errorf("unmarshal field ugent_department_name: %w", err)
				}
			}
		case person.FieldOrcidBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_bio", values[i])
			} else if value.Valid {
				pe.OrcidBio = value.String
			}
		case person.FieldOrcidID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_id", values[i])
			} else if value.Valid {
				pe.OrcidID = value.String
			}
		case person.FieldOrcidSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.OrcidSettings); err != nil {
					return fmt.Errorf("unmarshal field orcid_settings: %w", err)
				}
			}
		case person.FieldOrcidToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_token", values[i])
			} else if value.Valid {
				pe.OrcidToken = value.String
			}
		case person.FieldOrcidVerify:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_verify", values[i])
			} else if value.Valid {
				pe.OrcidVerify = value.String
			}
		case person.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				pe.Active = value.Bool
			}
		case person.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				pe.Deleted = value.Bool
			}
		case person.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case person.FieldRoles:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field roles", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Roles); err != nil {
					return fmt.Errorf("unmarshal field roles: %w", err)
				}
			}
		case person.FieldPublicationCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field publication_count", values[i])
			} else if value.Valid {
				pe.PublicationCount = int(value.Int64)
			}
		case person.FieldUgentMemorialisID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ugent_memorialis_id", values[i])
			} else if value.Valid {
				pe.UgentMemorialisID = value.String
			}
		case person.FieldPreferredFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_first_name", values[i])
			} else if value.Valid {
				pe.PreferredFirstName = value.String
			}
		case person.FieldPreferredLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_last_name", values[i])
			} else if value.Valid {
				pe.PreferredLastName = value.String
			}
		case person.FieldReplaces:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field replaces", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Replaces); err != nil {
					return fmt.Errorf("unmarshal field replaces: %w", err)
				}
			}
		case person.FieldReplacedBy:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field replaced_by", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ReplacedBy); err != nil {
					return fmt.Errorf("unmarshal field replaced_by: %w", err)
				}
			}
		case person.FieldDateLastLogin:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_last_login", values[i])
			} else if value.Valid {
				pe.DateLastLogin = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("date_created=")
	builder.WriteString(pe.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_updated=")
	builder.WriteString(pe.DateUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("object_class=")
	builder.WriteString(pe.ObjectClass)
	builder.WriteString(", ")
	builder.WriteString("ugent_username=")
	builder.WriteString(pe.UgentUsername)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(pe.FirstName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(pe.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(pe.LastName)
	builder.WriteString(", ")
	builder.WriteString("ugent_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentID))
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(pe.BirthDate)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", pe.Gender))
	builder.WriteString(", ")
	builder.WriteString("nationality=")
	builder.WriteString(pe.Nationality)
	builder.WriteString(", ")
	builder.WriteString("ugent_barcode=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentBarcode))
	builder.WriteString(", ")
	builder.WriteString("ugent_job_category=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentJobCategory))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pe.Title)
	builder.WriteString(", ")
	builder.WriteString("ugent_tel=")
	builder.WriteString(pe.UgentTel)
	builder.WriteString(", ")
	builder.WriteString("ugent_campus=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentCampus))
	builder.WriteString(", ")
	builder.WriteString("ugent_department_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentDepartmentID))
	builder.WriteString(", ")
	builder.WriteString("ugent_faculty_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentFacultyID))
	builder.WriteString(", ")
	builder.WriteString("ugent_job_title=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentJobTitle))
	builder.WriteString(", ")
	builder.WriteString("ugent_street_address=")
	builder.WriteString(pe.UgentStreetAddress)
	builder.WriteString(", ")
	builder.WriteString("ugent_postal_code=")
	builder.WriteString(pe.UgentPostalCode)
	builder.WriteString(", ")
	builder.WriteString("ugent_locality=")
	builder.WriteString(pe.UgentLocality)
	builder.WriteString(", ")
	builder.WriteString("ugent_last_enrolled=")
	builder.WriteString(pe.UgentLastEnrolled)
	builder.WriteString(", ")
	builder.WriteString("home_street_address=")
	builder.WriteString(pe.HomeStreetAddress)
	builder.WriteString(", ")
	builder.WriteString("home_postal_code=")
	builder.WriteString(pe.HomePostalCode)
	builder.WriteString(", ")
	builder.WriteString("home_locality=")
	builder.WriteString(pe.HomeLocality)
	builder.WriteString(", ")
	builder.WriteString("home_country=")
	builder.WriteString(pe.HomeCountry)
	builder.WriteString(", ")
	builder.WriteString("home_tel=")
	builder.WriteString(pe.HomeTel)
	builder.WriteString(", ")
	builder.WriteString("dorm_street_address=")
	builder.WriteString(pe.DormStreetAddress)
	builder.WriteString(", ")
	builder.WriteString("dorm_postal_code=")
	builder.WriteString(pe.DormPostalCode)
	builder.WriteString(", ")
	builder.WriteString("dorm_locality=")
	builder.WriteString(pe.DormLocality)
	builder.WriteString(", ")
	builder.WriteString("dorm_country=")
	builder.WriteString(pe.DormCountry)
	builder.WriteString(", ")
	builder.WriteString("research_discipline=")
	builder.WriteString(fmt.Sprintf("%v", pe.ResearchDiscipline))
	builder.WriteString(", ")
	builder.WriteString("research_discipline_code=")
	builder.WriteString(fmt.Sprintf("%v", pe.ResearchDisciplineCode))
	builder.WriteString(", ")
	builder.WriteString("ugent_expiration_date=")
	builder.WriteString(pe.UgentExpirationDate)
	builder.WriteString(", ")
	builder.WriteString("uzgent_job_title=")
	builder.WriteString(fmt.Sprintf("%v", pe.UzgentJobTitle))
	builder.WriteString(", ")
	builder.WriteString("uzgent_department_name=")
	builder.WriteString(fmt.Sprintf("%v", pe.UzgentDepartmentName))
	builder.WriteString(", ")
	builder.WriteString("uzgent_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UzgentID))
	builder.WriteString(", ")
	builder.WriteString("ugent_ext_category=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentExtCategory))
	builder.WriteString(", ")
	builder.WriteString("ugent_appointment_date=")
	builder.WriteString(pe.UgentAppointmentDate)
	builder.WriteString(", ")
	builder.WriteString("ugent_department_name=")
	builder.WriteString(fmt.Sprintf("%v", pe.UgentDepartmentName))
	builder.WriteString(", ")
	builder.WriteString("orcid_bio=")
	builder.WriteString(pe.OrcidBio)
	builder.WriteString(", ")
	builder.WriteString("orcid_id=")
	builder.WriteString(pe.OrcidID)
	builder.WriteString(", ")
	builder.WriteString("orcid_settings=")
	builder.WriteString(fmt.Sprintf("%v", pe.OrcidSettings))
	builder.WriteString(", ")
	builder.WriteString("orcid_token=")
	builder.WriteString(pe.OrcidToken)
	builder.WriteString(", ")
	builder.WriteString("orcid_verify=")
	builder.WriteString(pe.OrcidVerify)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", pe.Active))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", pe.Deleted))
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", pe.Settings))
	builder.WriteString(", ")
	builder.WriteString("roles=")
	builder.WriteString(fmt.Sprintf("%v", pe.Roles))
	builder.WriteString(", ")
	builder.WriteString("publication_count=")
	builder.WriteString(fmt.Sprintf("%v", pe.PublicationCount))
	builder.WriteString(", ")
	builder.WriteString("ugent_memorialis_id=")
	builder.WriteString(pe.UgentMemorialisID)
	builder.WriteString(", ")
	builder.WriteString("preferred_first_name=")
	builder.WriteString(pe.PreferredFirstName)
	builder.WriteString(", ")
	builder.WriteString("preferred_last_name=")
	builder.WriteString(pe.PreferredLastName)
	builder.WriteString(", ")
	builder.WriteString("replaces=")
	builder.WriteString(fmt.Sprintf("%v", pe.Replaces))
	builder.WriteString(", ")
	builder.WriteString("replaced_by=")
	builder.WriteString(fmt.Sprintf("%v", pe.ReplacedBy))
	builder.WriteString(", ")
	builder.WriteString("date_last_login=")
	builder.WriteString(pe.DateLastLogin.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person