package domain

type Level string

const (
	LevelUnknown      Level = ""
	LevelBeginner     Level = "beginner"
	LevelIntermediate Level = "intermediate"
	LevelAdvanced     Level = "advanced"
)

func (s Level) IsValid() bool {
	return s != LevelUnknown
}

func (s Level) String() string {
	return string(s)
}

func ToLevel(value string) Level {
	switch value {
	case LevelBeginner.String():
		return LevelBeginner
	case LevelIntermediate.String():
		return LevelIntermediate
	case LevelAdvanced.String():
		return LevelAdvanced
	default:
		return LevelUnknown
	}
}
