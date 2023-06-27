package routing

const (
	ErrorDirectoryNotExist        = "Specified client directory does not exists"
	ErrorIsNotADirectory          = "Client app directory, specified in config is not a directory"
	ErrorCouldNotServeFromRootDir = "Using root directory '/' for holding server app is strictly prohibited"
	ErrorClientDocumentNotExists  = "Client app document does not exists"
	ErrorClientDocumentWrongPath  = "Holding client app document not within client app directory is prohibited"
)

type routingGuardError struct {
	msg string
}

func (err routingGuardError) Error() string {
	return err.msg
}
