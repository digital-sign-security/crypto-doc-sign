package domains

type Doc struct {
	ID              string
	HashDS          string
	DecryptedText   string
	SenderUserID    string
	RecipientUserID string
}
