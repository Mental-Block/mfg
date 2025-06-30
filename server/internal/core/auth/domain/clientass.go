package domain

type ClientAssertion string

const (
	// SessionClientAssertion is used to authenticate using session cookie
	SessionClientAssertion ClientAssertion = "session"
	// AccessTokenClientAssertion is used to authenticate using access token generated
	// by the system for the user
	AccessTokenClientAssertion ClientAssertion = "access_token"
	// OpaqueTokenClientAssertion is used to authenticate using opaque token generated
	// for API clients
	OpaqueTokenClientAssertion ClientAssertion = "opaque"
	// JWTGrantClientAssertion is used to authenticate using JWT token generated
	// using public/private key pair that provides access token for the client
	JWTGrantClientAssertion ClientAssertion = "jwt_grant"
	
	// ClientCredentialsClientAssertion is used to authenticate using client_id and client_secret
	// that provides access token for the client
	ClientCredentialsClientAssertion ClientAssertion = "client_credentials"

	// PassthroughHeaderClientAssertion is used to authenticate using headers passed by the client
	// this is non secure way of authenticating client in test environments
	PassthroughHeaderClientAssertion ClientAssertion = "passthrough_header"
)

func (a ClientAssertion) String() string {
	return string(a)
}

var APIAssertions = []ClientAssertion{
	SessionClientAssertion,
	AccessTokenClientAssertion,
	OpaqueTokenClientAssertion,
	JWTGrantClientAssertion,
	
	// ClientCredentialsClientAssertion should be removed in future to avoid DDOS attacks on CPU
	// and should only be allowed to be used get access token for the client
	// ClientCredentialsClientAssertion,
	PassthroughHeaderClientAssertion,
}
