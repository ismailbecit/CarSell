package request

type CarInsert struct {
	Name       string `json:"name"`
	Userfk     uint   `gorm:"column:userfk" json:"userfk"`
	Categoryfk uint   `gorm:"column:categoryfk" json:"categoryfk"`
}
type CarDel struct {
	ID uint `json:"id"`
}
