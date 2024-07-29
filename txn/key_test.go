package txn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyIsEqualTo(t *testing.T) {
	key := NewStringKey("consensus")
	assert.True(t, key.IsRawKeyEqualTo(NewStringKey("consensus")))
}

func TestKeyIsNotEqualTo(t *testing.T) {
	key := NewStringKey("consensus")
	assert.False(t, key.IsRawKeyEqualTo(NewStringKey("raft")))
}

func TestKeySize(t *testing.T) {
	key := NewStringKey("consensus")
	assert.Equal(t, 17, key.EncodedSizeInBytes())
}

func TestKeyIsLessThan(t *testing.T) {
	key := NewStringKey("consensus")
	assert.True(t, key.IsLessThanOrEqualTo(NewStringKey("raft")))
}

func TestKeyIsLessThanOrEqualTo(t *testing.T) {
	key := NewStringKey("consensus")
	assert.True(t, key.IsLessThanOrEqualTo(NewStringKey("consensus")))
}

func TestKeyIsNotLessThanOrEqualTo(t *testing.T) {
	key := NewStringKey("consensus")
	assert.False(t, key.IsLessThanOrEqualTo(NewStringKey("accurate")))
}

func TestKeyComparisonLessThan(t *testing.T) {
	key := NewStringKey("consensus")
	assert.Equal(t, -1, key.Compare(NewStringKey("distributed")))
}

func TestKeyComparisonEqualTo(t *testing.T) {
	key := NewStringKey("consensus")
	assert.Equal(t, 0, key.Compare(NewStringKey("consensus")))
}

func TestKeyComparisonGreaterThan(t *testing.T) {
	key := NewStringKey("consensus")
	assert.Equal(t, 1, key.Compare(NewStringKey("accurate")))
}
