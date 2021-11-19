package style

func (s *Service) Navbar() string {
	return map[Theme]string{
		LightTheme: "bg-gray-800",
		DarkTheme:  "bg-black",
	}[s.theme]
}
