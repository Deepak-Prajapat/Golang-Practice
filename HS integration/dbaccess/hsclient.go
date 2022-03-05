package dbaccess

func CheckForAPI(apiKey string) (*Client, error) {
	var client *Client
	err := s.db.Where("api_key=?", apiKey).First(&client).Error
	if err != nil {
		return nil, err
	}
	return client, nil
}
