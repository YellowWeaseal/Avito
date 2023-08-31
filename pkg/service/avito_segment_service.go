package service

import (
	"awesomeProject1"
	"awesomeProject1/pkg/repository"
)

type AvitoSegmentService struct {
	repo repository.AvitoSegment
}

func NewAvitoSegmentService(repo repository.AvitoSegment) *AvitoSegmentService {
	return &AvitoSegmentService{repo: repo}
}

func (s *AvitoSegmentService) Create(segment awesomeProject1.SegmentsUpdate) (int, error) {
	return s.repo.Create(segment)
}

func (s *AvitoSegmentService) GetAll() ([]awesomeProject1.Segment, error) {
	return s.repo.GetAll()
}

func (s *AvitoSegmentService) GetById(segmentId int) (awesomeProject1.Segment, error) {
	return s.repo.GetById(segmentId)
}

func (s *AvitoSegmentService) Delete(segmentId int) error {
	return s.repo.Delete(segmentId)
}

func (s *AvitoSegmentService) Update(segmentId int, input awesomeProject1.SegmentsUpdate) (int, error) {
	/*if err := input.Validate(); err != nil {
		return err
	}
	*/
	return s.repo.Update(segmentId, input)
}
