package access_token

// AccessToken is the response to the public client
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenDetails contains the token info
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails is the dto for db data
type AccessDetails struct {
	AccessUuid string `json:"access_uuid"`
	UserID     string `json:"user_id"`
}
