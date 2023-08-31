package awesomeProject1

type SegmentsUpdate struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
}
type Segment struct {
	SegmentId   int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
