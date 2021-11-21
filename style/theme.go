package style

type Theme struct {
	ActionLinkAlert   string
	ActionLinkPrimary string
	BackgroundColor   string
	Banner            BannerStyles
	BodyText          string
	Button            ButtonStyles
	Form              FormStyles
	HintText          string
	MainLink          string
	MainLinkDisabled  string
	Modal             ModalStyles
	NavbarBackground  string
	NavbarLink        string
	SectionHeader     string
	Table             TableStyles
}

type BannerStyles struct {
	Container string
	Text      string
}

type ButtonStyles struct {
	Primary   string
	Secondary string
	Alert     string
}

type FormStyles struct {
	Input  string
	Select string
}

type ModalStyles struct {
	BackgroundColor string
	HeaderText      string
	BodyText        string
}

type TableStyles struct {
	RowBackground    string
	RowAltBackground string
	RowText          string
}
