package style

type Theme struct {
	ActionLinkAlert   string
	ActionLinkPrimary string
	BackgroundColor   string
	Button            ButtonStyles
	HintText          string
	MainLink          string
	MainLinkDisabled  string
	NavbarBackground  string
}

type ButtonStyles struct {
	Primary   string
	Secondary string
	Alert     string
}
