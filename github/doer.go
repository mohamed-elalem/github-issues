package github

type Doer interface {
	Do() (Issues, error)
}
