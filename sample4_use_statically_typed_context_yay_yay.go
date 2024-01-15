package gophercon22accessglobalresources

import "context"

type RequestContext interface {
	Request() *MockRequest
}
type DatabaseContext interface {
	Database() *MockDatabase
}
type HttpClientContext interface {
	HttpClient() *MockHttp
}
type SecretsContext interface {
	Secrets() *MockSecrets
}
type LoggerContext interface {
	Logger() *MockLogger
}

// Disadvantages of this approach:
//
//   - Huh?
//
// Advantages of this approach:
//
//   - Nice code, we can pass all dependencies at once
//   - Easy to mock/inject dependencies
//   - Easy to add new dependencies
//   - Easy to see upon exactly what the function does depend
//   - When we pass interface object to dependency functions we actually pass only the needed subset of interfaces
//     that the dependency function actually needs
//   - We can now have multiple requests (before we could only use 1 because it was global)
//   - If some dependency function introduce new dependency - our callsite code will break at compile-time! Yay!
func UseInterfacedContext(
	ctx interface {
		context.Context
		RequestContext
		DatabaseContext
		HttpClientContext
		SecretsContext
		LoggerContext
	},
	thing string,
) error {
	userKey, err := ctx.Request().UserKey()
	if err != nil {
		return err
	}

	user, err := ctx.Database().Read(userKey, ctx)
	if err != nil {
		return err
	}

	if user.CanDoThing(thing) {
		err = ctx.HttpClient().Post("nihongo.jp", user.ID())
	}
	return err

}
