package internal

import "github.com/sparkymat/archmark/internal/service/features"

type FeaturesService interface {
	Enabled(feature features.FeatureFlag) bool
}
