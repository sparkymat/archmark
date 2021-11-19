package style

func (s *Service) BackgroundColor() string {
	return map[Theme]string{
		LightTheme: "bg-white",
		DarkTheme:  "bg-gray-800",
	}[s.theme]
}
