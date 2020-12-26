package domain

// Jokes is for return value
// type Jokes struct {
// 	Jokes []Joke `json:"jokes"`
// 	// ID   string `json:"id"`
// 	// Joke string `json:"joke"`
// }

// RequestParam suppouse to contain new joke
type RequestParam struct {
	ID   string `json:"id"`
	Joke string `json:"joke"`
}
