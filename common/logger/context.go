package logger

import (
	"context"
	"github.com/rs/zerolog"
)

const SessionKey = 1

type Session struct {
	CID    string
	Logger zerolog.Logger
}

func SessionCid(ctx context.Context) string {
	session, ok := ctx.Value(SessionKey).(*Session)

	// Handle if session middleware is not used
	if !ok {
		return ""
	}

	return session.CID
}

func SessionLogger(ctx context.Context) zerolog.Logger {
	session, ok := ctx.Value(SessionKey).(*Session)

	// Handle if session middleware is not used
	if !ok {
		return zerolog.Logger{}
	}

	return session.Logger
}

func NewSessionCtx(cid string, log zerolog.Logger) context.Context {
	session := Session{
		cid,
		log,
	}
	return context.WithValue(context.Background(), SessionKey, &session)
}
