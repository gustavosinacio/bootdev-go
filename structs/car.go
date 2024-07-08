package structs

// We use structs in Go to represent structured data. It's often convenient to
// group different types of variables together. For example, if we want to
// represent a car we could do the following
type Car struct {
	Make    string `json:"make"`
	Model   string `json:"model"`
	Doors   int    `json:"doors"`
	Mileage int    `json:"mileage"`
}

// This creates a new struct type called car. All cars have a make, model, doors
// and mileage. Structs in Go are often used to represent data that you might use
// a dictionary or object for in other languages.
