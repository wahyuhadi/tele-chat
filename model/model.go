package model

type GetUpdatesResponse struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date  int    `json:"date"`
			Text  string `json:"text"`
			Photo []struct {
				FileID       string `json:"file_id"`
				FileUniqueID string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
			}
			Document struct {
				FileName     string `json:"file_name"`
				MimeType     string `json:"mime_type"`
				FileID       string `json:"file_id"`
				FileUniqueID string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
			} `json:"document"`
		} `json:"message,omitempty"`
		EditedMessage struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date           int `json:"date"`
			EditDate       int `json:"edit_date"`
			ReplyToMessage struct {
				MessageID int `json:"message_id"`
				From      struct {
					ID        int    `json:"id"`
					IsBot     bool   `json:"is_bot"`
					FirstName string `json:"first_name"`
					Username  string `json:"username"`
				} `json:"from"`
				Chat struct {
					ID    int64  `json:"id"`
					Title string `json:"title"`
					Type  string `json:"type"`
				} `json:"chat"`
				Date     int    `json:"date"`
				Text     string `json:"text"`
				Entities []struct {
					Offset int    `json:"offset"`
					Length int    `json:"length"`
					Type   string `json:"type"`
				} `json:"entities"`
			} `json:"reply_to_message"`
			Text string `json:"text"`
		} `json:"edited_message,omitempty"`
	} `json:"result"`
}
