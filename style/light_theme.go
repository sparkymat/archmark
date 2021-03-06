package style

func lightTheme() Theme {
	return Theme{
		ActionLinkAlert:   "text-sm text-gray-400 hover:text-red-400 border-b-2 border-dashed",
		ActionLinkPrimary: "text-sm text-gray-400 hover:text-blue-600 border-b-2 border-dashed",
		BackgroundColor:   "bg-white",
		Banner: BannerStyles{
			Container: "px-4 py-2 mt-4 bg-gray-200 rounded-sm border border-dashed flex flex-row justify-between items-center",
			Text:      "text-xl italic text-gray-600",
		},
		BodyText: "text-lg text-gray-700 font-light",
		Button: ButtonStyles{
			Primary:   "text-l text-white bg-gray-600 hover:bg-gray-800 rounded shadow-md px-8 py-2",
			Secondary: "text-l text-white bg-gray-300 hover:bg-gray-500 rounded shadow-md px-8 py-2",
			Alert:     "text-l text-white bg-red-600 hover:bg-red-800 rounded shadow-md px-8 py-2",
		},
		Form: FormStyles{
			Input:  "text-xl p-2 border bg-white text-black",
			Select: "text-xl p-2 border-b-2 bg-transparent border-dashed text-gray-700",
		},
		HintText:         "text-sm text-gray-400",
		MainLink:         "font-light text-black",
		MainLinkDisabled: "font-light text-gray-300",
		Modal: ModalStyles{
			BackgroundColor: "bg-white",
			HeaderText:      "text-lg font-bold text-black",
			BodyText:        "text-md text-gray-400",
		},
		NavbarBackground: "bg-gray-800",
		NavbarLink:       "text-white px-4 py-2 rounded-md text-sm font-medium",
		SectionHeader:    "text-2xl text-light py-2 border-b text-black",
		Table: TableStyles{
			RowBackground:    "bg-gray-200",
			RowAltBackground: "bg-gray-100",
			RowText:          "text-md italic break-all text-black",
		},
	}
}
