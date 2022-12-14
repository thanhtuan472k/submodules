package constant

import "golang.org/x/text/language"

// Constant
var (
	EnvDevelop    = "develop"
	EnvStaging    = "staging"
	EnvProduction = "production"

	LangVi   = language.Vietnamese.String()
	LangEn   = language.English.String()
	LangList = []interface{}{
		LangVi, LangEn,
	}

	OSNameIOS     = "iOS"
	OSNameAndroid = "android"
	OSNameList    = []interface{}{
		OSNameIOS, OSNameAndroid,
	}

	PhoneCountryCodeVietNam = "+84"
	PhoneCountryCodeList    = []interface{}{
		PhoneCountryCodeVietNam,
	}

	Limit10  = 10
	Limit20  = 20
	Limit50  = 50
	Limit500 = 500

	DateLayoutFull     = "2006-01-02T15:04:05.000Z"
	DateLayoutYYYYMMDD = "2006-01-02"

	StatusActive   = "active"
	StatusInactive = "inactive"
)

// Screen ...
var Screen = struct {
	Home    string
	Expense string
}{
	Home:    "home",
	Expense: "expense",
}

const (
	AuthGoogleRegex = "[0-9]{6}"
)
