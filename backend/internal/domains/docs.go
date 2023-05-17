package domains

type Doc struct {
	ID              string `json:"id"`
	HashDS          string `json:"hash"`
	Theme           string `json:"Theme"`
	DecryptedText   string `json:"decrypted_text"`
	SenderUserID    string `json:"sender_id"`
	RecipientUserID string `json:"recipient_id"`
}
