package creds

type UpdateBrokerCredRequest struct {
	Broker     string `uri:"broker" binding:"required"`
	TOTPSecret string `json:"totp_secret"`
	UserKey    string `json:"user_key"`
	PassKey    string `json:"pass_key"`
	AppCode    string `json:"app_code"`
	SecretKey  string `json:"secret_key"`
}

type UpdateBrokerCredResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
