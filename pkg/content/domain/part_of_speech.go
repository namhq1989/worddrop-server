package domain

import (
	"strings"
)

type PartOfSpeech string

const (
	PartOfSpeechAdjective                PartOfSpeech = "adj"   // Tính từ
	PartOfSpeechAdposition               PartOfSpeech = "adp"   // Giới từ
	PartOfSpeechAdverb                   PartOfSpeech = "adv"   // Trạng từ
	PartOfSpeechAuxiliary                PartOfSpeech = "aux"   // Trợ động từ
	PartOfSpeechConjunction              PartOfSpeech = "conj"  // Liên từ
	PartOfSpeechDeterminer               PartOfSpeech = "det"   // Mạo từ
	PartOfSpeechInterjection             PartOfSpeech = "intj"  // Thán từ
	PartOfSpeechNoun                     PartOfSpeech = "noun"  // Danh từ
	PartOfSpeechNumeral                  PartOfSpeech = "num"   // Số từ
	PartOfSpeechParticle                 PartOfSpeech = "part"  // Tiểu từ
	PartOfSpeechPronoun                  PartOfSpeech = "pron"  // Đại từ
	PartOfSpeechProperNoun               PartOfSpeech = "propn" // Danh từ riêng
	PartOfSpeechPunctuation              PartOfSpeech = "punct" // Dấu câu
	PartOfSpeechSubordinatingConjunction PartOfSpeech = "sconj" // Liên từ phụ thuộc
	PartOfSpeechSymbol                   PartOfSpeech = "sym"   // Ký hiệu
	PartOfSpeechVerb                     PartOfSpeech = "verb"  // Động từ
	PartOfSpeechOther                    PartOfSpeech = "x"     // Khác
	PartOfSpeechUnknown                  PartOfSpeech = ""
)

func (s PartOfSpeech) String() string {
	switch s {
	case PartOfSpeechAdjective, PartOfSpeechAdposition, PartOfSpeechAdverb,
		PartOfSpeechAuxiliary, PartOfSpeechConjunction, PartOfSpeechDeterminer,
		PartOfSpeechInterjection, PartOfSpeechNoun, PartOfSpeechNumeral,
		PartOfSpeechParticle, PartOfSpeechPronoun, PartOfSpeechProperNoun,
		PartOfSpeechPunctuation, PartOfSpeechSubordinatingConjunction,
		PartOfSpeechSymbol, PartOfSpeechVerb, PartOfSpeechOther:
		return string(s)
	default:
		return ""
	}
}

func (s PartOfSpeech) IsValid() bool {
	return s != PartOfSpeechUnknown
}

func ToPartOfSpeech(value string) PartOfSpeech {
	switch strings.ToLower(value) {
	case PartOfSpeechAdjective.String():
		return PartOfSpeechAdjective
	case PartOfSpeechAdposition.String():
		return PartOfSpeechAdposition
	case PartOfSpeechAdverb.String():
		return PartOfSpeechAdverb
	case PartOfSpeechAuxiliary.String():
		return PartOfSpeechAuxiliary
	case PartOfSpeechConjunction.String():
		return PartOfSpeechConjunction
	case PartOfSpeechDeterminer.String():
		return PartOfSpeechDeterminer
	case PartOfSpeechInterjection.String():
		return PartOfSpeechInterjection
	case PartOfSpeechNoun.String():
		return PartOfSpeechNoun
	case PartOfSpeechNumeral.String():
		return PartOfSpeechNumeral
	case PartOfSpeechParticle.String():
		return PartOfSpeechParticle
	case PartOfSpeechPronoun.String():
		return PartOfSpeechPronoun
	case PartOfSpeechProperNoun.String():
		return PartOfSpeechProperNoun
	case PartOfSpeechPunctuation.String():
		return PartOfSpeechPunctuation
	case PartOfSpeechSubordinatingConjunction.String():
		return PartOfSpeechSubordinatingConjunction
	case PartOfSpeechSymbol.String():
		return PartOfSpeechSymbol
	case PartOfSpeechVerb.String():
		return PartOfSpeechVerb
	case PartOfSpeechOther.String():
		return PartOfSpeechOther
	default:
		return PartOfSpeechUnknown
	}
}

var posMapping = map[string]string{
	"adjective":    "adj",
	"adj":          "adj",
	"noun":         "noun",
	"n":            "noun",
	"verb":         "verb",
	"v":            "verb",
	"adverb":       "adv",
	"adv":          "adv",
	"pronoun":      "pron",
	"preposition":  "adp",
	"adp":          "adp",
	"conjunction":  "cconj",
	"cconj":        "cconj",
	"determiner":   "det",
	"det":          "det",
	"exclamation":  "intj",
	"interjection": "intj",
	"intj":         "intj",
	"numeral":      "num",
	"num":          "num",
	"particle":     "part",
	"part":         "part",
	"proper noun":  "propn",
	"propn":        "propn",
	"punctuation":  "punct",
	"punct":        "punct",
	"symbol":       "sym",
	"sym":          "sym",
	"x":            "x",
}

func MappingPos(tag string) string {
	if mappedTag, exists := posMapping[tag]; exists {
		return mappedTag
	} else {
		return tag
	}
}
