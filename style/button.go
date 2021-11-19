package style

type ButtonType string

const (
	ButtonPrimary   ButtonType = "primary"
	ButtonSecondary ButtonType = "secondary"
	ButtonAlert     ButtonType = "alert"
)

func (s *Service) Button(buttonType ButtonType) string {
	return map[Theme]map[ButtonType]string{
		LightTheme: map[ButtonType]string{
			ButtonPrimary:   "text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-8 py-2",
			ButtonSecondary: "text-l text-white bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			ButtonAlert:     "text-l text-white bg-red-600 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
		DarkTheme: map[ButtonType]string{
			ButtonPrimary:   "text-l text-black bg-gray-200 hover:bg-gray-400 rounded shadow-md px-8 py-2",
			ButtonSecondary: "text-l text-black bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			ButtonAlert:     "text-l text-black bg-red-400 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
	}[s.theme][buttonType]
}
