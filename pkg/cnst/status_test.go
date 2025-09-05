package cnst_test

import (
	"gin-demo-framework/pkg/cnst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_String(t *testing.T) {
	assert.NotEqual(t, cnst.StatusInProgress.String(), "进行中1")
	assert.Equal(t, cnst.StatusCompleted.String(), "已完成")
	assert.Equal(t, cnst.StatusCancelled.String(), "已取消")
	assert.Equal(t, cnst.StatusPending.String(), "未开始")
}

func TestStatus(t *testing.T) {
	assert.Equal(t, cnst.StatusInProgress, cnst.Status(1))
	assert.Equal(t, cnst.StatusCompleted, cnst.Status(2))
	assert.Equal(t, cnst.StatusCancelled, cnst.Status(3))
	assert.Equal(t, cnst.StatusPending, cnst.Status(0))
	assert.NotEqual(t, cnst.StatusPending, 0)
}
