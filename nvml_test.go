package scheduler

import (
	nvml "github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNVMLLink(t *testing.T) {
	assert.Equal(t, nvml.ERROR_UNKNOWN, nvml.Return(999))
}
