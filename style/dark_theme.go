package style

func DarkTheme() Theme {
	return Theme{
		ActionLinkAlert:   "text-sm text-gray-200 hover:text-red-500 border-b-2 border-dashed",
		ActionLinkPrimary: "text-sm text-gray-200 hover:text-blue-500 border-b-2 border-dashed",
		BackgroundColor:   "bg-black",
		Banner: BannerStyles{
			Container: "px-4 py-2 mt-4 bg-gray-800 rounded-sm border border-dashed flex flex-row justify-between items-center",
			Text:      "text-xl italic text-gray-200",
		},
		BodyText: "text-lg text-gray-300",
		Button: ButtonStyles{
			Primary:   "text-l text-black bg-gray-200 hover:bg-gray-400 rounded shadow-md px-8 py-2",
			Secondary: "text-l text-black bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			Alert:     "text-l text-black bg-red-400 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
		Form: FormStyles{
			Input:  "text-xl p-2 border bg-gray-900 text-white",
			Select: "text-xl p-2 border-b-2 bg-transparent border-dashed text-gray-300",
		},
		HintText:         "text-sm text-gray-200",
		MainLink:         "text-white",
		MainLinkDisabled: "text-gray-500",
		Modal: ModalStyles{
			BackgroundColor: "bg-gray-800",
			HeaderText:      "text-lg font-bold text-white",
			BodyText:        "text-md text-gray-400 text-gray-200",
		},
		NavbarBackground: "bg-gray-800",
		NavbarLink:       "text-gray-500 px-4 py-2 rounded-md text-sm font-medium",
		SectionHeader:    "text-2xl text-light py-2 border-b text-white",
		Table: TableStyles{
			RowBackground:    "bg-gray-800",
			RowAltBackground: "bg-gray-600",
			RowText:          "text-md italic break-all text-white",
		},
	}
}
