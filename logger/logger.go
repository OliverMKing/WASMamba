package logger

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/olivermking/wasmamba/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var (
	def       *zap.Logger
	loggerKey = ctxKey{}
)

func getDef() *zap.Logger {

	if def != nil {
		return def
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	def = zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(zap.InfoLevel),
		))
	return def
}

// WithContext sets the logger as the logger for the context
func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the logger from the context
func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok && logger != nil {
		return logger
	}

	return getDef()
}

// WithRequest sets common structured fields for a request to the logger for the context
func WithRequest(ctx context.Context, r *http.Request) context.Context {
	logger := FromContext(ctx)
	return WithContext(ctx, logger.With(
		zap.String("url", r.URL.String()),
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),

		// typically Battlensake game id + turn functions can function operation id
		// but some routes don't have that and we also need an operation id
		// in case of failure to marshal
		zap.String("operationId", uuid.NewString()),
	))
}

// WithGame sets common structured fields for a game to the logger for the context
func WithGame(ctx context.Context, g model.GameReq) context.Context {
	logger := FromContext(ctx)
	return WithContext(ctx, logger.With(
		zap.String("gameId", g.Game.Id),
		zap.Int("turn", g.Turn),
	))
}
