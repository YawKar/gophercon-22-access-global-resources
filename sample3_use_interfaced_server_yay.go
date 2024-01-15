package gophercon22accessglobalresources

import "context"

type RequestServer interface {
	Request() *MockRequest
}
type DatabaseServer interface {
	Database() *MockDatabase
}
type HttpClientServer interface {
	HttpClient() *MockHttp
}
type SecretsServer interface {
	Secrets() *MockSecrets
}
type LoggerServer interface {
	Logger() *MockLogger
}

// Disadvantages of this approach:
//
//   - Still pass ctx.Context separately
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
func UseInterfacedServer(
	ctx context.Context,
	server interface {
		RequestServer
		DatabaseServer
		HttpClientServer
		SecretsServer
		LoggerServer
	},
	thing string,
) error {
	userKey, err := server.Request().UserKey()
	if err != nil {
		return err
	}

	user, err := server.Database().Read(userKey, ctx, server)
	if err != nil {
		return err
	}

	if user.CanDoThing(thing) {
		err = server.HttpClient().Post("nihongo.jp", user.ID())
	}
	return err

}
