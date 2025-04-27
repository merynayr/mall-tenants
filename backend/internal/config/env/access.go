package env

import (
	"errors"
	"os"
	"strings"

	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/model"
)

const (
	directorEndpointsEnvName  = "DIRECTOR_ENDPOINTS"
	moderatorEndpointsEnvName = "MODERATOR_ENDPOINTS"
	clientEndpointsEnvName    = "CLIENT_ENDPOINTS"
)

type accessConfig struct {
	userAccesses map[model.UserRole]map[string]struct{}
}

// NewAccessConfig returns new access config
func NewAccessConfig() (config.AccessConfig, error) {
	userAccesses := make(map[model.UserRole]map[string]struct{})

	userAccesses[model.RoleModerator] = make(map[string]struct{})
	for entry := range strings.SplitSeq(os.Getenv(moderatorEndpointsEnvName), ",") {
		entry = strings.TrimSpace(entry)
		if entry != "" {
			userAccesses[model.RoleModerator][entry] = struct{}{}
		}
	}

	userAccesses[model.RoleClient] = make(map[string]struct{})
	for entry := range strings.SplitSeq(os.Getenv(clientEndpointsEnvName), ",") {
		entry = strings.TrimSpace(entry)
		if entry != "" {
			userAccesses[model.RoleClient][entry] = struct{}{}
		}
	}

	userAccesses[model.RoleDirector] = make(map[string]struct{})
	for entry := range strings.SplitSeq(os.Getenv(directorEndpointsEnvName), ",") {
		entry = strings.TrimSpace(entry)
		if entry != "" {
			userAccesses[model.RoleModerator][entry] = struct{}{}
		}
	}

	return &accessConfig{
		userAccesses: userAccesses,
	}, nil
}

// UserAccessesMap returns map of endpoints allowed to users
func (cfg *accessConfig) UserAccessesMap() (map[model.UserRole]map[string]struct{}, error) {
	if cfg.userAccesses == nil {
		return nil, errors.New("user access map is nil")
	}
	return cfg.userAccesses, nil
}
