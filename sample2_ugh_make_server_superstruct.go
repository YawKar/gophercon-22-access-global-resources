package gophercon22accessglobalresources

type Server struct {
	database   *MockDatabase
	request    *MockRequest
	httpClient *MockHttp
}

// Disadvantages of this approach:
//
//   - Still there's no way to see what parts of the server the function uses other than dig into implementation details.
//     For example, the server may contain 40 resources but the function uses only 2 of them. :)
//   - When we pass server to other function that this function depends upon, we pass the whole server and doing that we give permissions to all resources
//
// Advantages of this approach:
//
//   - Nice code, we can pass all dependencies at once
//   - Easy to mock/inject dependencies
//   - Easy to add new dependencies
//   - We can now have multiple requests (before we could only use 1 because it was global)
func UseServerSuperStruct(
	thing string,
	server *Server,
) error {
	userKey, err := server.request.UserKey()
	if err != nil {
		return err
	}

	user, err := server.database.Read(userKey, server)
	if err != nil {
		return err
	}

	if user.CanDoThing(thing) {
		err = server.httpClient.Post("nihongo.jp", user.ID())
	}
	return err

}
