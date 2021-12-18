package request

type CarInsert struct {
	Userfk     uint `gorm:"column:userfk" json:"userfk"`
	Categoryfk uint `gorm:"column:categoryfk" json:"categoryfk"`
}
type CarDel struct {
	ID uint `json:"id"`
}

type CarInfo struct {
	ID uint `json:"id"`
}
