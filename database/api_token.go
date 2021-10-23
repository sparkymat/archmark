package database

import "github.com/sparkymat/archmark/model"

func (s *service) ListAPITokens() ([]model.APIToken, error) {
	panic("unimplemented")
	/*
		var apiTokens []model.APIToken

		if result := s.conn.Find(&apiTokens); result.Error != nil {
			return nil, result.Error
		}

		return apiTokens, nil
	*/
}

func (s *service) DeleteAPIToken(id uint) error {
	panic("unimplemented")
	/*
		err := s.conn.Delete(&model.APIToken{}, id)

		return err.Error
	*/
}

func (s *service) CreateAPIToken(token string) (*model.APIToken, error) {
	panic("unimplemented")
	/*
		apiToken := &model.APIToken{
			Token: token,
		}

		if result := s.conn.Create(&apiToken); result.Error != nil {
			return nil, result.Error
		}

		return apiToken, nil
	*/
}
