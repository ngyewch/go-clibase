package slog_utils

import (
	"encoding/json"
	"log/slog"
)

type JsonLogValuer struct {
	Value any
}

func (v JsonLogValuer) LogValue() slog.Value {
	jsonBytes, err := json.Marshal(v.Value)
	if err != nil {
		return slog.StringValue(err.Error())
	}
	return slog.StringValue(string(jsonBytes))
}

type JsonIndentedLogValuer struct {
	Value any
}

func (v JsonIndentedLogValuer) LogValue() slog.Value {
	jsonBytes, err := json.MarshalIndent(v.Value, "", "  ")
	if err != nil {
		return slog.StringValue(err.Error())
	}
	return slog.StringValue(string(jsonBytes))
}
