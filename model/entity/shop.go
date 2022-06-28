package entity

type Shop struct {
	Id          string
	SiupNumber  string `db:"no_siup"`
	Name        string
	Address     string
	PhoneNumber string `db:"phone"`
}
