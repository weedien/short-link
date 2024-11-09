package errno

import (
	"errors"
	"testing"
)

func TestJoinErrors(t *testing.T) {
	slugError := SlugError{
		errorType: ErrorTypeServiceError,
		msg:       "短链接查询异常",
	}
	anotherError := errors.New("数据库连接异常")
	joinedError := errors.Join(slugError, anotherError)

	t.Logf("joinedError: %v", joinedError)

	var newSlugError SlugError
	if ok := errors.As(joinedError, &newSlugError); !ok {
		t.Errorf("errors.Join() failed")
	}
	t.Logf("newSlugError: %v", newSlugError)
}
