package features

type (
	Flag    uint64
	Toggles uint64
)

func New(flags ...Flag) Toggles {
	var r Toggles
	for _, m := range flags {
		r |= Toggles(1 << m)
	}

	return r
}

func (u Toggles) Has(flags ...Flag) bool {
	for _, m := range flags {
		if u&Toggles(1<<m) == 0 {
			return false
		}
	}

	return true
}
