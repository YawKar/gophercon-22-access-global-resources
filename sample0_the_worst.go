package gophercon22accessglobalresources

// Disadvantages of this approach:
//
//   - This function doesn't document what it uses.
//
//   - We don't know that it uses the database by looking at the function's signature.
//
//   - We don't know what else it might use.
//
//   - Testing this code is tricky cz you have to mock out all of the resources and
//
// you don't even know what resources are
// (without digging deeply into the implementation not only of that function
// but of all functions called along the way).
func UseGlobalsStraightForwardly(thing string) error {
	// First session-scope global dependency
	userKey, err := request.UserKey()
	if err != nil {
		return err
	}

	// Second global dependency
	user, err := database.Read(userKey)
	if err != nil {
		return err
	}

	// Third - it can do http posts if it can (IO global side-effect)
	if user.CanDoThing(thing) {
		err = httpClient.Post("nihongo.jp", user.ID())
	}
	return err
}
