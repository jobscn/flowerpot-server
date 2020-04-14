package common

import "context"

// context.Context用的是指针this，因此context.Context一定是指针传值，这里效率不会低
func ContextWithGin(ctx context.Context, ginContext interface{}) context.Context {
	return context.WithValue(ctx, "gin-context", ginContext)
}