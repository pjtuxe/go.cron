package job

type VariableModel struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}
