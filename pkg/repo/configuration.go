package repo

type ConfigRecord interface {
	PermanentId() string
	VersionId() string
}

func History[C ConfigRecord](r C) ([]C, error) {
	// TODO implement me
	panic("not implemented")
}
