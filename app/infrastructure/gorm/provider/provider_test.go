package provider_test

import (
	"queryservice/infrastructure/gorm/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryProvider(t *testing.T) {
	provider, err := provider.NewRepositoryProvider()
	if err != nil {
		assert.Error(t, err)
		return
	}
	assert.NotNil(t, provider.CategoryRep)
	assert.NotNil(t, provider.ProductRep)
}
