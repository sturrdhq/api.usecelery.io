package config

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/fatih/color"
)

type PrettyJSONHandler struct {
	w    io.Writer // Store the writer ourselves
	opts *slog.HandlerOptions
	mu   sync.Mutex // For thread safety
}

func NewPrettyJSONHandler(w io.Writer, opts *slog.HandlerOptions) *PrettyJSONHandler {
	return &PrettyJSONHandler{
		w:    w,
		opts: opts,
	}
}

func (h *PrettyJSONHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *PrettyJSONHandler) Handle(_ context.Context, r slog.Record) error {
	logMap := make(map[string]any)
	logMap["time"] = r.Time.Format("2006-01-02T15:04:05Z07:00")
	logMap["level"] = r.Level.String()
	logMap["msg"] = r.Message

	r.Attrs(func(a slog.Attr) bool {
		logMap[a.Key] = a.Value.Any()
		return true
	})

	var buf bytes.Buffer
	prettyJSON, err := json.MarshalIndent(logMap, "", "  ")
	if err != nil {
		return err
	}
	buf.Write(prettyJSON)
	buf.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()

	h.w.Write([]byte("\n"))
	switch r.Level {
	case slog.LevelDebug:
		color.New(color.BgBlue).Fprintln(h.w, "DEBUG!")
	case slog.LevelError:
		color.New(color.BgRed).Fprintln(h.w, "ERROR!")
	case slog.LevelWarn:
		color.New(color.BgYellow).Fprintln(h.w, "WARN!")
	case slog.LevelInfo:
		color.New(color.BgGreen).Fprintln(h.w, "INFO!")
	}

	_, err = h.w.Write(buf.Bytes())
	h.w.Write([]byte("\n"))

	return err
}

func (h *PrettyJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewPrettyJSONHandler(h.w, h.opts)
}

func (h *PrettyJSONHandler) WithGroup(name string) slog.Handler {
	return h
}

func SetupLogger() {
	handler := NewPrettyJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
