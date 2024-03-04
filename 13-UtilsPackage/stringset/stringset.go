package stringset

type Set map[string]struct{}

func New() Set {
	return map[string]struct{}{}
}

func (s Set) Sort() []string {
	return nil
}
