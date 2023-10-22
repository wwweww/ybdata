package module

type Sex string

type StData struct {
	S_id       string `gorm:"column:studentId;type:char(11);not null" json:"studentId,omitempty"`
	S_name     string `gorm:"column:name;type:varchar(5); not null" json:"name,omitempty"`
	S_sex      Sex    `gorm:"column:sex;type:char(2); not null" json:"sex,omitempty"`
	S_conllege string `gorm:"column:college;type:varchar(20); not null" json:"conllege,omitempty"`
	S_major    string `gorm:"column:major;type:varchar(20); not null" json:"major,omitempty"`
	S_class    string `gorm:"column:classid;type:varchar(20); not null" json:"classid,omitempty"`
	S_grade    int    `gorm:"column:grade;type:int; not null" json:"grade,omitempty"`
	S_ssystem  int    `gorm:"column:years;type:int; not null" json:"system,omitempty"`
}
