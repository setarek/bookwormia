package logger

import (
	"github.com/rs/zerolog/log"
)

var Logger = log.Logger.With().Caller().Logger()
