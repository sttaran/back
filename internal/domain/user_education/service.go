package user_education

import (
	"back/internal/domain/education_institution"
	"context"
	"errors"
)

type UserEducationService struct {
	userEducationStorage  *Storage
	eduInstitutionStorage *education_institution.Storage
}

func NewUserEducationService(userEducationStorage *Storage, eduInstitutionStorage *education_institution.Storage) *UserEducationService {
	return &UserEducationService{
		userEducationStorage:  userEducationStorage,
		eduInstitutionStorage: eduInstitutionStorage,
	}
}

func (u *UserEducationService) Create(ctx context.Context, dto CreateUserEducationDTO) (int64, error) {
	if dto.EduInstitutionId == 0 {
		if dto.EduInstitutionName == "" {
			return 0, errors.New("edu_institution institution not provided")
		}
		var eduInstitutionDto education_institution.CreateEducationInstitutionDTO = education_institution.CreateEducationInstitutionDTO{
			FullName:    dto.EduInstitutionName,
			ShortName:   "",
			Description: "",
		}

		id, err := u.eduInstitutionStorage.Create(ctx, eduInstitutionDto)
		if err != nil {
			return 0, err
		}

		dto.EduInstitutionId = id
	}

	if dto.InProgress == 1 {
		dto.EndDate = dto.StartDate
	}

	id, err := u.userEducationStorage.Create(ctx, dto)
	if err != nil {
		return 0, err
	}

	return id, nil
}
