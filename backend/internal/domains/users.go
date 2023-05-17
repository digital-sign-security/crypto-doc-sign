package domains

type User struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithKey struct {
	ID        string `json:"-"`
	Username  string `json:"username"`
	PublicKey string `json:"public_key"`
}

type JWTToken struct {
	ID      string `json:"-"`
	Token   string `json:"token"`
	IsAlive bool   `json:"is_alive"`
	UserID  string `json:"user_id"`
}
