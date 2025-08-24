package logger

import (
	"errors"
	"fmt"
	stdlog "log"
	"os"
	"syscall"
	"time"

	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/zaytcevcom/golang-ninja/internal/buildinfo"
)

//go:generate options-gen -out-filename=logger_options.gen.go -from-struct=Options -defaults-from=var
type Options struct {
	level          string `option:"mandatory" validate:"required,oneof=debug info warn error"`
	productionMode bool
	clock          zapcore.Clock
	sentryDSN      string `validate:"omitempty,url"`
	env            string `validate:"omitempty"`
}

var defaultOptions = Options{
	clock: zapcore.DefaultClock,
}

func MustInit(opts Options) {
	if err := Init(opts); err != nil {
		panic(err)
	}
}

func Init(opts Options) error {
	if err := opts.Validate(); err != nil {
		return fmt.Errorf("validate options: %v", err)
	}

	level, err := zapcore.ParseLevel(opts.level)
	if err != nil {
		return fmt.Errorf("parse log level: %w", err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "level",
		NameKey:        "component",
		CallerKey:      "caller",
		MessageKey:     "msg",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}

	var encoder zapcore.Encoder
	if opts.productionMode {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), level)
	cores := []zapcore.Core{core}

	if opts.sentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         opts.sentryDSN,
			Environment: opts.env,
			Release:     buildinfo.BuildInfo.Main.Version,
		}); err != nil {
			return fmt.Errorf("init sentry: %w", err)
		}

		cfg := zapsentry.Configuration{
			Level:             zapcore.WarnLevel,
			EnableBreadcrumbs: true,
			BreadcrumbLevel:   zapcore.InfoLevel,
		}
		sentryCore, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(sentry.CurrentHub().Client()))
		if err != nil {
			return fmt.Errorf("create zapsentry core: %w", err)
		}
		cores = append(cores, sentryCore)
	}

	l := zap.New(zapcore.NewTee(cores...), zap.WithClock(opts.clock))
	zap.ReplaceGlobals(l)

	return nil
}

func Sync() {
	if err := zap.L().Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		stdlog.Printf("cannot sync logger: %v", err)
	}

	sentry.Flush(2 * time.Second)
}
