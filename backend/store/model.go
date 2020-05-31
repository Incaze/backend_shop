package store

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type Exception struct {
	Message string `json:"message"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Price       uint64  `json:"price"`
	Rating      float32 `json:"rating"`
	Category    string  `json:"category"`
}

type Products []Product
