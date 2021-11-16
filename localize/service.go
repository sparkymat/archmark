package localize

import "fmt"

func New() *Service {
	return &Service{
		translations: map[Language]map[StringIdentifier]string{
			English:   englishStrings,
			Malayalam: malayalamStrings,
			German:    germanStrings,
		},
	}
}

type Service struct {
	translations map[Language]map[StringIdentifier]string
}

func (s *Service) Lookup(lang Language, identifier StringIdentifier, args ...interface{}) string {
	stringMap, ok := s.translations[lang]
	if !ok {
		return "?"
	}

	localizedString, ok := stringMap[identifier]
	if !ok {
		return "?"
	}

	if len(args) == 0 {
		return localizedString
	}

	return fmt.Sprintf(localizedString, args...)
}
