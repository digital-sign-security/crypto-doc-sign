package domains

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

type JWTToken struct {
	ID      string
	Token   string
	IsAlive bool
	UserID  string
}
