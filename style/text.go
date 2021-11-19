package style

func (s *Service) MainLink() string {
	return "font-light text-black"
}

func (s *Service) MainLinkDisabled() string {
	return "font-light text-gray-300"
}

func (s *Service) ActionLinkPrimary() string {
	return "text-sm text-gray-400 hover:text-blue-400 border-b-2 border-dashed"
}

func (s *Service) ActionLinkAlert() string {
	return "text-sm text-gray-400 hover:text-red-400 border-b-2 border-dashed"
}

func (s *Service) HintText() string {
	return "text-sm text-gray-400"
}
