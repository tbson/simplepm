package infra

type InputData struct {
	Locked       bool   `json:"locked" form:"locked" validate:"required"`
	LockedReason string `json:"locked_reason" form:"locked_reason" validate:"required"`
}
