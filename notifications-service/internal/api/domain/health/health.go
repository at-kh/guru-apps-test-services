package health

type (
	// Info describes meta information about application.
	Info struct {
		Name         string
		BuildCommit  string
		BuildDate    string
		BuildVersion string
	}
)
