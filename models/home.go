package models

type Banner struct {
	ID        uint   `json:"id"`
	Image     string `json:"image"`
	Link      string `json:"link"`
	Is_Active bool   `json:"is_active"`
}
type BannerResponse struct {
	ID    uint   `json:"id"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

type News struct {
	ID             int    `json:"id"`
	Image          string `json:"image"`
	TM_description string `json:"tm_description"`
	TM_title       string `json:"tm_title"`
	EN_title       string `json:"en_title"`
	RU_title       string `json:"ru_title"`
	EN_description string `json:"en_description"`
	RU_description string `json:"ru_description"`
	View           int    `json:"view"`
	Date           string `json:"date"`
}
type Media struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Cover    string `json:"cover"`
	Video    string `json:"video"`
	TM_title string `json:"tm_title"`
	EN_title string `json:"en_title"`
	RU_title string `json:"ru_title"`
	Date     string `json:"date"`
	View     int    `json:"view"`
}
type Employer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Major   string `json:"major"`
	Image   string `json:"image"`
}
type Laws struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Laws  string `json:"laws"`
}
type Views struct {
	ID     int `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID int `json:"user_id"`
	NewsID int `json:"news_id"`
}
