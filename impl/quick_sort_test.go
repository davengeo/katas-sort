package impl

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShouldDoQuickSort(t *testing.T) {
	assert.Equal(t, "Deniss", "Deniss")
	result, _ := QuickSort([]string{"Zorba", "David", "Olli", "Evegenijs", "David", "Antons", "Ivo", "Deniss"});

	assert.Equal(t, result[0], "Antons")
}