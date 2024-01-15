package gophercon22accessglobalresources

// Disadvantages of this approach:
//
//   - In order to add some dependency you need to manually, parameter-by-parameter compare if it
//
// already has this dependency with some other name
//   - It is not simple and callsites look crazy bloated: function(arg1, arg2, arg3, arg4, arg5.........)
//   - It is not easy to mock/inject dependencies (well, it seems easy but if we add up other disadvantages it won't)
//
// Advantages:
//
//   - Dependencies are explicit
func UseExplicitDependencies(
	thing string,
	request *MockRequest,
	database *MockDatabase,
	httpClient *MockHttp,
	secrets *MockSecrets,
	logger *MockLogger,
	timeout *MockTimeout,
) error {
	userKey, err := request.UserKey()
	if err != nil {
		return err
	}

	// notice how we bloat our code when try to pass secrets, logger and timeout into some dependency function
	user, err := database.Read(userKey, secrets, logger, timeout)
	if err != nil {
		return err
	}

	if user.CanDoThing(thing) {
		err = httpClient.Post("nihongo.jp", user.ID())
	}
	return err
}
