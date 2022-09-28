package performance

// Code generated by cdproto-gen. DO NOT EDIT.

// EventMetrics current values of the metrics.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#event-metrics
type EventMetrics struct {
	Metrics []*Metric `json:"metrics"` // Current values of the metrics.
	Title   string    `json:"title"`   // Timestamp title.
}
