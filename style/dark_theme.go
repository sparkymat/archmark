package style

func DarkTheme() Theme {
	return Theme{
		ActionLinkAlert:   "text-sm text-gray-200 hover:text-red-200 border-b-2 border-dashed",
		ActionLinkPrimary: "text-sm text-gray-200 hover:text-blue-200 border-b-2 border-dashed",
		BackgroundColor:   "bg-gray-800",
		Button: ButtonStyles{
			Primary:   "text-l text-black bg-gray-200 hover:bg-gray-400 rounded shadow-md px-8 py-2",
			Secondary: "text-l text-black bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			Alert:     "text-l text-black bg-red-400 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
		HintText:         "text-sm text-gray-200",
		MainLink:         "text-white",
		MainLinkDisabled: "text-gray-500",
		NavbarBackground: "bg-black",
	}
}
