package features

type (
	Flag    uint64
	Toggles uint64
)

// New creates a new Toggles instance with the specified flags set.
func New(flags ...Flag) Toggles {
	var toggles Toggles
	for _, flag := range flags {
		toggles |= Toggles(1 << flag)
	}

	return toggles
}

// Has checks if all the specified flags are set in the Toggles instance.
// It returns true if all flags are present, otherwise false.
func (t Toggles) Has(flags ...Flag) bool {
	for _, f := range flags {
		if t&Toggles(1<<f) == 0 {
			return false
		}
	}

	return true
}
