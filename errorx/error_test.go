package errorx

import (
	"errors"
	"fmt"
	"github.com/v-mars/library/errorx/code"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrPermissionCode            = int32(1000000)
	errPermissionMessage         = "unauthorized access : {msg}"
	errPermissionAffectStability = false
)

func TestError(t *testing.T) {
	code.Register(
		ErrPermissionCode,
		errPermissionMessage,
		code.WithAffectStability(errPermissionAffectStability),
	)

	err := New(ErrPermissionCode, KV("msg", "test"))
	fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Println(err)

	var customErr StatusError
	b := errors.As(err, &customErr)
	assert.Equal(t, b, true)
	assert.Equal(t, customErr.Code(), ErrPermissionCode)
	assert.Equal(t, customErr.Msg(), strings.Replace(errPermissionMessage, "{msg}", "test", 1))
}
