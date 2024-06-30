package structure

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      int                `json:"userId"`
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phoneNumber"`
	NID         string             `json:"nid"`
	Birthdate   string             `json:"birthdate"`
	FatherName  string             `json:"fatherName"`
	MotherName  string             `json:"motherName"`
    UserType    UserType           `json:"userType"`
	// validation func () error       // Custom validation function for user input data
}

var (
    phoneRegex = regexp.MustCompile(`^0[0-9]{10}$`) // Matches 11-digit phone numbers starting with 0
    nidRegex   = regexp.MustCompile(`^[0-9]{13}$`)  // Matches 13-digit NID numbers
)

// userValidation validates the User struct fields.
func (u *User) userValidation() error {
    if u.UserID <= 0 {
        return errors.New("UserID must be greater than zero")
    }
    if u.Name == "" {
        return errors.New("Name cannot be empty")
    }
    if u.PhoneNumber != "" && !phoneRegex.MatchString(u.PhoneNumber) {
        return errors.New("Invalid phone number format")
    }
    if u.NID != "" && !nidRegex.MatchString(u.NID) {
        return errors.New("Invalid NID format")
    }
    // You can add more validation rules as needed for other fields
    return nil
}

// func user_validation(){

// }

type Client struct {
	
	User             User               `json:"user"`
	Location           string             `json:"location"`
}

type ServiceProvider struct {
	User               User               `json:"user"`
	Skill              string             `json:"skill"`
	Location           string             `json:"location"`
	Education          Education          `json:"education"`
	VerifiedByPorichoy bool               `json:"verifiedByporichoy"`
	SPBalance          Balance            `json:"Balance"`
}

type Education struct {
	Level     string `json:"level"`
	Institute string `json:"institute"`
}

type Balance struct {
	Amount float64 `json:"amount"`
}

type Review struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Review        string             `json:"review" bson:"review"`
	Timelines     float64            `json:"timelines" bson:"timelines"`
	Quality       float64            `json:"quality" bson:"quality"`
	Communication float64            `json:"communication" bson:"communication"`
	Behavior      float64            `json:"behavior" bson:"behavior"`
}

type Bid struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description"`
	Time        string             `json:"t_time"`
	BidAmount   float64            `json:"bidAmount"`
	PostedTime  time.Time          `json:"postedTime,omitempty" bson:"postedTime,omitempty"`
}

type Payment struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SPID       string             `json:"sp_id"`
	Balance    float64            `json:"balance"`
	CashinDate time.Time          `json:"cashinDate,omitempty" bson:"cashinDate,omitempty"`
	FirstJob   time.Time          `json:"firstJob,omitempty" bson:"firstJob,omitempty"`
}

type GpsCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Point struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title                    string             `json:"title"`
	Address                  string             `json:"address"`
	Gps                      []GpsCoordinate    `json:"gps"`
	ContactPerson            string             `json:"contactPerson"`
	ContactPersonPhoneNumber string             `json:"contactPersonPhoneNumber"`
}

type UserType string

const (
	UserTypeClient          UserType = "client"
	UserTypeServiceProvider UserType = "serviceProvider"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusAccepted Status = "accepted"
	StatusRejected Status = "rejected"
)


type QuestionAnswer struct {
	Questions              []string        `json:"questions"`
	Answers                []string        `json:"answers"`
	ServiceProviderDetails ServiceProvider `json:"serviceProviderDetails"`
}

type JobStatus string

const (
	JobStatusJobPosted   JobStatus = "job_posted"
	JobStatusBidAccepted JobStatus = "bid_accepted"
	JobStatusJobStarted  JobStatus = "job_started"
	JobStatusBan         JobStatus = "ban"
)

type StatusChange string

const (
	StatusChangePostingTime    StatusChange = "posting_time"
	StatusChangeAcceptanceTime StatusChange = "acceptance_time"
)


type Job struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title            string             `json:"title"`
	Posted           time.Time          `json:"posted"`
	Budget           string             `json:"budget"`
	Description      string             `json:"description"`
	Clients          Client             `json:"clients"`
	ServiceProviders ServiceProvider    `json:"serviceProviders"`
	Status           Status             `json:"status"`
	JobStatus        JobStatus          `json:"jobStatus"`
	StatusChange    []StatusChange       `json:"statusChange"`
	Point            []Point            `json:"point,omitempty"`
	QuestionAnswer   []QuestionAnswer   `json:"questionAnswer,omitempty"`
	Bid              []Bid              `json:"bid,omitempty"`
	Review           Review           `json:"review,omitempty"`
}

var structureType = []string{"User", "Client", "ServiceProvider", "Review", "Bid", "Payment", "Point", "Job", "QuestionAnswer"}
