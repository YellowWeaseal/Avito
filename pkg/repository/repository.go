package repository

import (
	"awesomeProject1"
	"github.com/jmoiron/sqlx"
)

type AvitoUser interface {
	Create(user awesomeProject1.UserUpdate) (int, error)
	GetAll() ([]awesomeProject1.User, error)
	GetById(userId int) (awesomeProject1.User, error)
	Delete(userId int) error
	Update(userId int, input awesomeProject1.UserUpdate) (int, error)
}

type AvitoSegment interface {
	Create(segment awesomeProject1.SegmentsUpdate) (int, error)
	GetAll() ([]awesomeProject1.Segment, error)
	GetById(segmentId int) (awesomeProject1.Segment, error)
	Delete(segmentId int) error
	Update(segmentId int, input awesomeProject1.SegmentsUpdate) (int, error)
}

type Repository struct {
	AvitoUser
	AvitoSegment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AvitoUser:    NewAvitoUserPostgres(db),
		AvitoSegment: NewAvitoSegmentPostgres(db),
	}
}
