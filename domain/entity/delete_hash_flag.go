package entity

type DeleteHashFlag bool

const (
	DeleteHashFlagFalse DeleteHashFlag = false
	DeleteHashFlagTrue  DeleteHashFlag = true
)

func (f *DeleteHashFlag) ToBool() bool {
	return bool(*f)
}
