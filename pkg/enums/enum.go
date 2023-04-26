package enums

type UserRole int

const (
	URNone UserRole = iota // 0
	URUser
	URAdmin
	URSuperUser
	URService
)
