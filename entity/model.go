package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BloodDetails struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BloodGroup  string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units       int                `bson:"units,omitempty" json:"units,omitempty"`
	Location    string             `bson:"location,omitempty" json:"location,omitempty"`
	DepositDate time.Time          `bson:"deposit_date,omitempty" json:"deposit_date,omitempty"`
	CreatedDate time.Time          `bson:"created_date,omitempty" json:"created_date,omitempty"`
}

type DonorDetails struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Age         int64              `bson:"age,omitempty" json:"age,omitempty"`
	DOB         time.Time          `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup  string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units       string             `bson:"units,omitempty" json:"units,omitempty"`
	DepositDate time.Time          `bson:"deposit_date,omitempty" json:"deposit_date,omitempty"`
	Location    string             `bson:"location,omitempty" json:"location,omitempty"`
	AdharCard   string             `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	Active      bool               `bson:"active,omitempty" json:"active,omitempty"`
	MailId      string             `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
}

type UserDetails struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Age         int64              `bson:"age,omitempty" json:"age,omitempty"`
	DOB         time.Time          `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup  string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	AdharCard   string             `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	CreatedDate time.Time          `bson:"created_date,omitempty" json:"created_date,omitempty"`
	Location    string             `bson:"location,omitempty" json:"location,omitempty"`
	Active      bool               `bson:"active,omitempty" json:"active,omitempty"`
	MailId      string             `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
}

type PatientDetails struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Age         int64              `bson:"age,omitempty" json:"age,omitempty"`
	DOB         time.Time          `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup  string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	AdharCard   string             `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	CreatedDate time.Time          `bson:"created_date,omitempty" json:"created_date,omitempty"`
	Location    string             `bson:"location,omitempty" json:"location,omitempty"`
	Active      bool               `bson:"active,omitempty" json:"active,omitempty"`
	ApplyUnits  string             `bson:"apply_units,omitempty" json:"apply_units,omitempty"`
	ApplyDate   time.Time          `bson:"apply_date,omitempty" json:"apply_date,omitempty"`
	GivenDate   time.Time          `bson:"given_date,omitempty" json:"given_date,omitempty"`
	BloodDate   time.Time          `bson:"blood_date,omitempty" json:"blood_date,omitempty"`
	MailId      string             `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
}

type LoginDetails struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MailId   string             `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Active   bool               `bson:"active,omitempty" json:"active,omitempty"`
}

type BloodDetailsRequest struct {
	BloodGroup  string `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units       string `bson:"units,omitempty" json:"units,omitempty"`
	Location    string `bson:"location,omitempty" json:"location,omitempty"`
	DepositDate string `bson:"deposit_date,omitempty" json:"deposit_date,omitempty"`
}

type DonorDetailsRequest struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Age         int64  `bson:"age,omitempty" json:"age,omitempty"`
	DOB         string `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup  string `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units       string `bson:"units,omitempty" json:"units,omitempty"`
	DepositDate string `bson:"deposit_date,omitempty" json:"deposit_date,omitempty"`
	Location    string `bson:"location,omitempty" json:"location,omitempty"`
	AdharCard   string `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	MailId      string `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password    string `bson:"password,omitempty" json:"password,omitempty"`
}

type UserDetailsRequest struct {
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	Age        int64  `bson:"age,omitempty" json:"age,omitempty"`
	DOB        string `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup string `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Location   string `bson:"location,omitempty" json:"location,omitempty"`
	AdharCard  string `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	MailId     string `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password   string `bson:"password,omitempty" json:"password,omitempty"`
}

type PatientDetailsRequest struct {
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	Age        int64  `bson:"age,omitempty" json:"age,omitempty"`
	DOB        string `bson:"dob,omitempty" json:"dob,omitempty"`
	BloodGroup string `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	AdharCard  string `bson:"adhar_card,omitempty" json:"adhar_card,omitempty"`
	Location   string `bson:"location,omitempty" json:"location,omitempty"`
	ApplyUnits string `bson:"apply_units,omitempty" json:"apply_units,omitempty"`
	BloodDate  string `bson:"blood_date,omitempty" json:"blood_date,omitempty"`
	MailId     string `bson:"mail_id,omitempty" json:"mail_id,omitempty"`
	Password   string `bson:"password,omitempty" json:"password,omitempty"`
}
