package gophercon22accessglobalresources

import "context"

// Disadvantages of this approach:
//
//   - Ugly code, not simple
//   - Still there's no way to see what parts of the server the function uses other than dig into implementation details.
//     For example, the context may contain 40 resources but the function uses only 2 of them. :)
//   - When we pass ctx to other function that this function depends upon, we pass the whole ctx and doing that we give permissions to all resources
//
// Advantages of this approach:
//
//   - Easy to mock/inject dependencies
//   - Easy to add new dependencies
//   - We can now have multiple requests (before we could only use 1 because it was global)
func UseContextSuperStruct(
	thing string,
	ctx context.Context,
) error {
	userKey, err := ctx.Value("request").(*MockRequest).UserKey()
	if err != nil {
		return err
	}

	user, err := ctx.Value("database").(*MockDatabase).Read(userKey, ctx)
	if err != nil {
		return err
	}

	if user.CanDoThing(thing) {
		err = ctx.Value("httpClient").(*MockHttp).Post("nihongo.jp", user.ID())
	}
	return err

}
