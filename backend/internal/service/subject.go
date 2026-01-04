package service

import (
	"context"
	"errors"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
)

type SubjectService interface {
	GetSubjectById(ctx context.Context, id int64) (*domain.Subject, error)
}

type subjectService struct {
	subjectRepository repository.SubjectRepository
}

func NewSubjectService(subjectRepository repository.SubjectRepository) SubjectService {
	return &subjectService{
		subjectRepository: subjectRepository,
	}
}

func (s *subjectService) GetSubjectById(ctx context.Context, id int64) (*domain.Subject, error) {
	if id <= 0 {
		return nil, errors.New("invalid subject id")
	}
	resp, err := s.subjectRepository.GetSubjectById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Subject{
		Id:   resp.Id,
		Name: resp.Name,
	}, nil
}
