package style

func LightTheme() Theme {
	return Theme{
		ActionLinkAlert:   "text-sm text-gray-400 hover:text-red-400 border-b-2 border-dashed",
		ActionLinkPrimary: "text-sm text-gray-400 hover:text-blue-400 border-b-2 border-dashed",
		BackgroundColor:   "bg-white",
		Button: ButtonStyles{
			Primary:   "text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-8 py-2",
			Secondary: "text-l text-white bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			Alert:     "text-l text-white bg-red-600 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
		HintText:         "text-sm text-gray-400",
		MainLink:         "font-light text-black",
		MainLinkDisabled: "font-light text-gray-300",
		NavbarBackground: "bg-gray-800",
	}
}
