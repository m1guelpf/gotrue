package provider

import (
	"context"

	"github.com/supabase/gotrue/internal/conf"
	"golang.org/x/oauth2"
)

const (
	defaultWorldIDAPIBase = "id.worldcoin.org"
)

type WorldcoinProvider struct {
	*oauth2.Config
	APIPath string
}

type WorldcoinUser struct {
	ID      string  `json:"sub"`
	WorldID WorldID `json:"https://id.worldcoin.org/beta"`
}

type WorldID struct {
	LikelyHuman    string `json:"likely_human"`
	CredentialType string `json:"credential_type"`
}

// NewWorldcoinProvider creates a Worldcoin account provider.
func NewWorldcoinProvider(ext conf.OAuthProviderConfiguration) (OAuthProvider, error) {
	if err := ext.Validate(); err != nil {
		return nil, err
	}

	idPath := chooseHost(ext.URL, defaultWorldIDAPIBase)

	oauthScopes := []string{
		"openid",
	}

	return &WorldcoinProvider{
		Config: &oauth2.Config{
			ClientID:     ext.ClientID,
			ClientSecret: ext.Secret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  idPath + "/authorize",
				TokenURL: idPath + "/token",
			},
			Scopes:      oauthScopes,
			RedirectURL: ext.RedirectURI,
		},
		APIPath: idPath,
	}, nil
}

func (g WorldcoinProvider) GetOAuthToken(code string) (*oauth2.Token, error) {
	return g.Exchange(context.Background(), code)
}

func (g WorldcoinProvider) GetUserData(ctx context.Context, tok *oauth2.Token) (*UserProvidedData, error) {
	var u WorldcoinUser
	if err := makeRequest(ctx, tok, g.Config, g.APIPath+"/userinfo", &u); err != nil {
		return nil, err
	}

	return &UserProvidedData{
		Metadata: &Claims{
			Issuer:       g.APIPath,
			Subject:      u.ID,
			CustomClaims: map[string]interface{}{"likely_human": u.WorldID.LikelyHuman, "credential_type": u.WorldID.CredentialType},
			ProviderId:   u.ID,
		},
	}, nil
}
