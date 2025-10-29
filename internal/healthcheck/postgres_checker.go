package healthcheck

type DBChecker struct {
}

func NewDBChecker() *DBChecker {
	return &DBChecker{}
}

func (c *DBChecker) Check() error {
	return nil
}
