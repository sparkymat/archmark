package localize

import "fmt"

type API interface {
	Lookup(lang Language, identifier StringIdentifier, args ...interface{}) string
}

func New() API {
	return &service{
		translations: map[Language]map[StringIdentifier]string{
			English:   englishStrings,
			Malayalam: malayalamStrings,
			German:    germanStrings,
		},
	}
}

type service struct {
	translations map[Language]map[StringIdentifier]string
}

func (s *service) Lookup(lang Language, identifier StringIdentifier, args ...interface{}) string {
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
