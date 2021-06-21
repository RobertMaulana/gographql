package schema

type Links struct {
	Base
	Title   string `gorm:"not null"`
	Address string `gorm:"not null"`
	UserId  int    `gorm:"not null"`
}

func (Links) TableName() string {
	return "links"
}

func (Links) Pk() string {
	return "id"
}

func (f Links) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Links) AddForeignKeys() {
}

func (f Links) InsertDefaults() {
}
