package twitter

import (
	"fmt"
)

// Client exposes the Twitter API web service layer controllers for both v1.1 and v2
type Client interface {
	V2() ControllerV2
}

type broker struct {
	// TODO: add support for v1.1
	v2 ControllerV2
}

// ClientConfig the Twitter API client configuration of both v1.1 and v2 controllers
type ClientConfig struct {
	ControllerConfigV2 ControllerConfigV2 `json:"controller_v2" yaml:"controller_v2"`
}

// NewClient constructs a Twitter API client compatible with both v1 and v2
func NewClient(cfg ClientConfig) (Client, error) {
	bkr := &broker{}
	// TODO: Add support for API v1
	if cfg.ControllerConfigV2 != (ControllerConfigV2{}) {
		v2, err := NewControllerV2(cfg.ControllerConfigV2)
		if err != nil {
			return nil, fmt.Errorf("failed to construct Twitter v2 API client")
		}

		bkr.v2 = v2
	}

	return bkr, nil
}

// V1 returns the Twitter API v2 controller if it has been enabled
//
// If not the controller operator will be empty to avoid throwing a nil-pointer exception
func (bkr *broker) V2() ControllerV2 {
	if bkr.v2 == nil {
		return &operatorV2{}
	}

	return bkr.v2
}
