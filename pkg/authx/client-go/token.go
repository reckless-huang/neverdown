package authority

import (
	"fmt"

	"github.com/kzz45/neverdown/pkg/jwttoken"
)

func Token(token string) (*jwttoken.Claims, error) {
	if token == "" {
		return nil, fmt.Errorf("nil token")
	}
	return jwttoken.Parse(token)
}
