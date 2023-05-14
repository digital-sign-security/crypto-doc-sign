package domains

type Doc struct {
	ID              string `json:"-"`
	HashDS          string `json:"hash"`
	Theme           string `json:"Theme"`
	DecryptedText   string `json:"decrypted_text"`
	SenderUserID    string `json:"sender_id"`
	RecipientUserID string `json:"recipient_id"`
}
