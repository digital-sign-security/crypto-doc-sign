package domains

type Keys struct {
	ID         string `json:"-"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	IsAlive    bool   `json:"is_alive"`
	UserID     string `json:"user_id"`
}
