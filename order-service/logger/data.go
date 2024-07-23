package logger

import (
	"context"
	"time"
)

// Finalize load from context and delete data context
func (d *DataLogger) Finalize(ctx context.Context) {
	value, ok := extract(ctx)
	if !ok {
		return
	}

	if i, ok := value.LoadAndDelete(_StatusCode); ok && i != nil {
		d.StatusCode = i.(int)
	}

	if i, ok := value.LoadAndDelete(_Response); ok && i != nil {
		d.Response = i
	}

	if i, ok := value.LoadAndDelete(_Messages); ok && i != nil {
		d.Messages = i.([]string)
	}

	if i, ok := value.LoadAndDelete(_ThirdParties); ok && i != nil {
		d.ThirdParties = i.([]ThirdParty)
	}

	d.ExecTime = time.Since(d.TimeStart).Seconds()
}

// Store for storing data context to third parties
func (th ThirdParty) Store(ctx context.Context) {
	var data []ThirdParty

	val, ok := extract(ctx)
	if !ok {
		return
	}

	tmp, ok := val.LoadAndDelete(_ThirdParties)
	if ok {
		data = tmp.([]ThirdParty)
	}

	data = append(data, th)

	val.Set(_ThirdParties, data)
}
