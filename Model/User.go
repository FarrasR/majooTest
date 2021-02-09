package model


type User struct {
  ID		uint  	`gorm:"primary_key"; gorm:"AUTO_INCREMENT"`
  Username	string 	`gorm:"type:varchar(20); unique_index"`
  Password	string 	`gorm:"type:varchar(60)"`
  Fullname	string 	`gorm:"type:varchar(40)"`
  Photo		string 	`gorm:"type:varchar(40)"`
}
