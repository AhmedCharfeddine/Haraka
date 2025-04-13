package mapping

var TransliterationMap = map[string]string{
	"aa": "ا", "b": "ب", "t": "ت", "th": "ث",
	"j": "ج", "7": "ح", "kh": "خ", "d": "د",
	"w": "و", "y": "ي", "i": "ي", "tt": "ط",
	"dh": "ذ", "r": "ر", "z": "ز", "s": "س",
	"ch": "ش", "sh": "ش", "ss": "ص", "3": "ع",
	"gh": "غ", "8": "غ", "f": "ف", "9": "ق",
	"k": "ك", "l": "ل", "m": "م", "n": "ن",
	"h": "ه", "a": "ا", "5": "خ", "ou": "و",
}

const ALIF_MAQSURA = "ى"
const TA_MARBUTA = "ة"

func isValidKey(key string) bool {
	_, exists := TransliterationMap[key]
	return exists
}
