package trace

// TraceName represents the tracing name.
const TraceName = "go-zero"

// A Config is an opentelemetry config.
type Config struct {
	Name        string            `json:",optional"`
	Endpoint    string            `json:",optional"`
	Sampler     float64           `json:",default=1.0"`
	Batcher     string            `json:",default=jaeger,options=jaeger|zipkin|otlpgrpc|otlphttp"`
	OtlpHeaders map[string]string `json:",optional"` // uptrace-dsn: 'http://project2_secret_token@localhost:14317/2'
}
