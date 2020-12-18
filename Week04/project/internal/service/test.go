package service

import (
	"github.com/pkg/errors"
)

func (s *Service) GetName(id int) (name string, err error) {
	row, err := s.dao.GetRow(id)
	if err != nil {
		return "", errors.Wrapf(err, "service:GetName id=%d", id)
	}
	return row.Name, nil
}
