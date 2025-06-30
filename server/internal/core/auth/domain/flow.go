package domain

import (
	"time"

	user "github.com/server/internal/core/user/domain"

	"github.com/server/pkg/metadata"
	"github.com/server/pkg/utils"
)

type FlowKey string

func (f FlowKey) String() string {
	return string(f)
}

func (f Flow) IsValid(currentTime time.Time) bool {
	return f.ExpiresAt.After(currentTime)
}

// Flow is a temporary state used to finish login/registration flows
type Flow struct {
	Id utils.UUID

	Reason Reason	

	// authentication flow type
	Strategy Strategy
	
	// Email is the email of the user
	Email Email

	// StartURL is where flow should start from for verification
	StartURL string

	// FinishURL is where flow should end to after successful verification
	FinishURL string

	// Nonce is a once time use random string
	Nonce string

	Metadata metadata.Metadata

	// CreatedAt will be used to clean-up dead auth flows
	CreatedAt time.Time

	// ExpiresAt is the time when the flow will expire
	ExpiresAt time.Time
}

type FinishFlowReq struct {
    Strategy 	string
	Reason 		string

    // used for OIDC & mail otp auth strategy
    Code        string
    State       string
    StateConfig map[string]any
}

type FinishFlowResp struct {
	User *user.User
	Flow *Flow
}

type StartFlowReq struct {
	Reason string

	Strategy string

	Email string

	Payload map[string]string

	// callback_url will be used by strategy as last step to finish authentication flow
	// in OIDC this host will receive "state" and "code" query params, in case of magic links
	// this will be the url where user is redirected after clicking on magic link.
	// For most cases it could be host of mfg but in case of proxies, this will be proxy public endpoint.
	// callback_url should be one of the allowed urls configured at instance level
	CallBackUrl string
	
	// ReturnToURL is where flow should end to after successful verification
	ReturnToURL string
} 

type StartFlowResp struct {
	Flow        *Flow
	State       string
	StateConfig map[string]any
}
