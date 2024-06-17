package main

import (
	"errordetails"
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	//"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	//zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//err := errordetails.NewErrorDetails(errors.New("error test")).Str("nuevo", "tema")
	//log.Error().Err(err).Msg("")
	//log.Error().Object("error", err).Msg("An error occurred")
	//log.Error().Stack().Err(err).Msg("")
	/*log.Log().
	Str("foo", "bar").
	EmbedObject(err).
	Msg("hello world")*/
	/*err2 := &errordetails.CustomError{
		File:    "C:/Users/CarrenoG/go/src/errordetails/cmd/main.go",
		Line:    11,
		Message: "error test",
	}

	// Usar .Object en lugar de .Err para asegurarse de que el error personalizado sea tratado como un objeto
	//log.Error().Object("error", err).Msg("An error occurred")
	log.Error().Err(err2).Msg("An error occurred")*/
	/*u := errordetails.NewUser()

	log := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Object("user", u).
		Logger()

	log.Log().Msg("hello world")*/

	errorcust := errordetails.NewErrorDetails(errors.New("error test")).Str("nuevo", "tema")
	/*log := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Object("error", errorcust).
		Logger()

	log.Log().Msg("hello world")*/

	// Log detailed error information
	log.Error().Object("error_details", errorcust).Msg("An error occurred")

	/*errorcust2 := errordetails.NewErrorDetails(errors.New("error test")).Str("nuevo", "tema")
	fmt.Println(errorcust2)
	//log.Error().Object("error_details", errorcust2).Msg("An error occurred")
	log := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Object("user", errorcust2).
		Logger()

	log.Log().Msg("hello world")*/
}
