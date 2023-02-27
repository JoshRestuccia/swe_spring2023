package user

type Stock struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	UserRefer string  `json:"user_email"`
	User      User    `gorm:"foreignKey:UserRefer"`
}
