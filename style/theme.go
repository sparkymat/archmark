package style

type Theme struct {
	ActionLinkAlert   string
	ActionLinkPrimary string
	BackgroundColor   string
	Banner            BannerStyles
	Button            ButtonStyles
	Form              FormStyles
	HintText          string
	MainLink          string
	MainLinkDisabled  string
	NavbarBackground  string
	NavbarLink        string
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
	Input string
}
