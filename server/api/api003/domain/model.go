package domain

// Jokes is for return value
// type Jokes struct {
// 	Jokes []Joke `json:"jokes"`
// 	// ID   string `json:"id"`
// 	// Joke string `json:"joke"`
// }

// RequestParam suppouse to contain new joke
type RequestParam struct {
	Joke     string `json:"joke"`
	ImageURL string `json:"image_url"`
}
