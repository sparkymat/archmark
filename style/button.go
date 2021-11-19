package style

type ButtonType string

const (
	ButtonPrimary   ButtonType = "primary"
	ButtonSecondary ButtonType = "secondary"
	ButtonAlert     ButtonType = "alert"
)

func (s *Service) Button(buttonType ButtonType) string {
	switch buttonType {
	case ButtonPrimary:
		return "text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-8 py-2"
	case ButtonSecondary:
		return "text-l text-white bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2"
	case ButtonAlert:
		return "text-l text-white bg-red-600 hover:bg-red-800 rounded shadow-md px-8 py-2"
	default:
		return "text-l text-white bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2"
	}
}
