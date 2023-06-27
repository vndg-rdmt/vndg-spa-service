package routing

import (
	"os"
	"path"
	"path/filepath"
	"silvex/app"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/*
A guard function that contains every assertion and input params validation due to security reasons,
before calling `serveClientApp` serving rouing modifier.
*/
func guardClientAppConfig() error {

	// Client directory
	DIRNAME, err := filepath.Abs(path.Clean(app.Config.Client.Directory))
	if err != nil {
		return err
	}

	// If directory is not exist
	if info, err := os.Stat(DIRNAME); os.IsNotExist(err) {
		return &routingGuardError{msg: ErrorDirectoryNotExist + ": " + DIRNAME}

		// If error occuried
	} else if err != nil {
		return err

		// If not a directory
	} else if !info.IsDir() {
		return &routingGuardError{msg: ErrorIsNotADirectory + ": " + DIRNAME}

		// If root directory is specified
	} else if "/" == DIRNAME {
		return &routingGuardError{msg: ErrorCouldNotServeFromRootDir}
	}

	// Document filename
	FILENAME, err := filepath.Abs(path.Join(DIRNAME, app.Config.Client.Document))
	if err != nil {
		return err

		// If not withing client directory
	} else if !strings.HasPrefix(FILENAME, DIRNAME) {
		return &routingGuardError{msg: ErrorClientDocumentWrongPath + ": " + FILENAME}

		// If not a file
	} else if _, err := os.Stat(FILENAME); os.IsNotExist(err) {
		return &routingGuardError{msg: ErrorClientDocumentNotExists + ": " + FILENAME}

		// If error occuried
	} else if err != nil {
		return err
	}

	app.Config.Client.Directory = DIRNAME
	app.Config.Client.Document = FILENAME
	return nil
}

/*
Function for serving client application from a specified directory.
*/
func serveClientApp(server *fiber.App) error {

	err := guardClientAppConfig()
	if err != nil {
		return err
	}

	server.Static("/", app.Config.Client.Directory)
	server.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile(app.Config.Client.Document)
	})
	app.Logs.Trace().
		Str("client.directory", app.Config.Client.Directory).
		Str("client.document", app.Config.Client.Document).
		Send()

	return nil
}
