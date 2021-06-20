package schema

type Users struct {
	Base
	Name string
	Username string `gorm:"not null"`
	Password string
}

func (Users) TableName() string {
	return "users"
}

func (Users) Pk() string {
	return "id"
}

func (f Users) Ref() string {
	return f.TableName() + "(" + f.Pk() + ")"
}

func (f Users) AddForeignKeys() {
}

func (f Users) InsertDefaults() {
}