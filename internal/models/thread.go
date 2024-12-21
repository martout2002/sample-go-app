package models

// Thread represents a single thread in the database
type Thread struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	StoreName     string  `json:"store_name"`
	StoreLocation string  `json:"store_location"`
	AuthorName    string  `json:"author_name"`
	Details       string  `json:"details"`
	Rating        float64 `json:"rating"`
	Comments      string  `json:"comments"`
}
