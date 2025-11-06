package healthcheck

type Checker interface {
	Check() error
}
