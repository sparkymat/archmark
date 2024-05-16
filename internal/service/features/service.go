package features

type FeatureFlag string

const FeatureFlagRegistrationDisabled FeatureFlag = "registration_disabled"

func New(cfg Config) *Service {
	featureFlagMap := map[FeatureFlag]bool{}

	if cfg.DisableRegistration() {
		featureFlagMap[FeatureFlagRegistrationDisabled] = true
	}

	return &Service{
		featureFlagMap: featureFlagMap,
	}
}

type Service struct {
	featureFlagMap map[FeatureFlag]bool
}

func (s *Service) Enabled(feature FeatureFlag) bool {
	enabled, present := s.featureFlagMap[feature]

	return present && enabled
}
