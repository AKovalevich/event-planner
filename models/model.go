package models

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in models
//    type User struct {
//      BaseModel
//    }
type BaseModel struct {
	ID        uint `json:"id" gorm:"primary_key" valid:"-;"`
	CreatedAt int64 `json:"created_at" valid:"-"`
	UpdatedAt int64 `json:"updated_at" valid:"-"`
	Deleted bool `json:"deteted" valid:"-"`
}
