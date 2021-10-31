package localize

type API interface {
	Lookup(lang Language, identifier StringIdentifier) string
}

func New() API {
	return &service{
		translations: map[Language]map[StringIdentifier]string{
			English: EnglishStrings,
		},
	}
}

type service struct {
	translations map[Language]map[StringIdentifier]string
}

func (s *service) Lookup(lang Language, identifier StringIdentifier) string {
	stringMap, ok := s.translations[lang]
	if !ok {
		return "?"
	}

	localizedString, ok := stringMap[identifier]
	if !ok {
		return "?"
	}

	return localizedString
}
