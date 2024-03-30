package app

import (
	"context"
	"testing"
	"time"

	"github.com/didasy/winmetrics/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	l := logrus.New()
	cfg := config.New()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := Run(ctx, cfg, l)
	assert.NoError(t, err)
}
