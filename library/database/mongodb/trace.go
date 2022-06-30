package mongodb

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/net/trace"
)

func TraceHook(c context.Context, db *conn, statement string) (cb func(err error)) {
	if t, ok := trace.FromContext(c); ok {
		t = forkTrace(t, db, fmt.Sprintf("Collection:%s", statement))
		t.SetTag(trace.String(trace.TagDBStatement, statement))
		cb = func(e error) {
			if e != nil {
				t.Finish(&e)
			} else {
				t.Finish(nil)
			}
		}
	}
	return
}
