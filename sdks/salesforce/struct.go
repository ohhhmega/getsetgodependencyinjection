package salesforce

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type" validate:"required"`
	Assertion    string `json:"assertion" validate:"required"`
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	InstanceURL string `json:"instance_url"`
	ID          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

type CreateTaskRequest struct {
	Status         string `json:"Status" validate:"required"`
	Subject        string `json:"Subject" validate:"required"`
	Priority       string `json:"Priority" validate:"required"`
	UserIDC        string `json:"UserId__c" validate:"required"`
	SocialNetworkC string `json:"social_network__c" validate:"required"`
}

type CreateTaskResponse struct {
	ID      string        `json:"id"`
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
}

type SalesforceSDKConfig struct {
	BaseURL    string
	Token      string
	APITimeout int
}
