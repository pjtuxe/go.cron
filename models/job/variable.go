package job

type VariableModel struct {
	Key   string `validate:"empty=false > empty=false [empty=false] > ne=0"`
	Value string `validate:"empty=false > empty=false [empty=false] > ne=0"`
}
