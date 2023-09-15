package slog_utils

import (
	"encoding/hex"
	"log/slog"
)

type HexStringLogValuer []byte

func (v HexStringLogValuer) LogValue() slog.Value {
	return slog.StringValue(hex.EncodeToString(v))
}
