package models

type Book struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"desc"`
	Image_url    string  `json:"image"`
	Release_year int     `json:"year"`
	Price        string  `json:"price"`
	Total_page   int     `json:"page"`
	Thickness    string  `json:"thickness"`
	Created_at   []uint8 `json:"created_at"` //error jika menggunakan tipe date time (storing driver.Value type []uint8 into type *time.Time)
	Update_at    []uint8 `json:"update_at"`  //error jika menggunakan tipe date time (storing driver.Value type []uint8 into type *time.Time)
	Category_id  int     `json:"category_id"`
}
