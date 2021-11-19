package style

func (s *Service) MainLink() string {
	return map[Theme]string{
		LightTheme: "font-light text-black",
		DarkTheme:  "text-white",
	}[s.theme]
}

func (s *Service) MainLinkDisabled() string {
	return map[Theme]string{
		LightTheme: "font-light text-gray-300",
		DarkTheme:  "text-gray-500",
	}[s.theme]
}

func (s *Service) ActionLinkPrimary() string {
	return map[Theme]string{
		LightTheme: "text-sm text-gray-400 hover:text-blue-400 border-b-2 border-dashed",
		DarkTheme:  "text-sm text-gray-200 hover:text-blue-200 border-b-2 border-dashed",
	}[s.theme]
}

func (s *Service) ActionLinkAlert() string {
	return map[Theme]string{
		LightTheme: "text-sm text-gray-400 hover:text-red-400 border-b-2 border-dashed",
		DarkTheme:  "text-sm text-gray-200 hover:text-red-200 border-b-2 border-dashed",
	}[s.theme]
}

func (s *Service) HintText() string {
	return map[Theme]string{
		LightTheme: "text-sm text-gray-400",
		DarkTheme:  "text-sm text-gray-200",
	}[s.theme]
}
