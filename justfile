set quiet

_default: _help

_help:
    just --list

check:
    weaver registry check -r model

resolve:
    weaver registry resolve -r model

fmt:
    go run github.com/google/yamlfmt/cmd/yamlfmt@latest model/attributes.yaml

# Generate all sdks
[group('generate')]
all: go ts py docs

# Generate go sdk
[group('generate')]
go:
    weaver registry generate -r model --templates "https://github.com/open-telemetry/opentelemetry-go.git[semconv/templates]" go .

# Generate ts sdk
[group('generate')]
ts:
    weaver registry generate -r model --templates "https://github.com/open-telemetry/opentelemetry-js.git[scripts/semconv/templates]" ts-stable src

[group('generate')]
py: && py-fmt
    weaver registry generate -r model --templates "https://github.com/open-telemetry/opentelemetry-python.git[scripts/semconv/templates]" '' ./semconv --param output="semconv" --param filter="any"

# Generate docs
[group('generate')]
docs:
    weaver registry generate -r model --templates "https://github.com/open-telemetry/semantic-conventions/archive/refs/tags/v1.34.0.zip[templates]" markdown docs

[group('py')]
py-fmt:
    uv run ruff format
