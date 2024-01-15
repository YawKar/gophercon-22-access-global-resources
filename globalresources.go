package gophercon22accessglobalresources

type (
	MockDatabase struct{}
	MockRequest  struct{}
	MockUser     struct{}
	MockHttp     struct{}
	MockSecrets  struct{}
	MockLogger   struct{}
	MockTimeout  struct{}
)

func (*MockDatabase) Read(userKey string, etc ...any) (MockUser, error) {
	return MockUser{}, nil
}

func (*MockRequest) UserKey() (string, error) {
	return "mockery mock", nil
}

func (*MockUser) CanDoThing(thing string) bool {
	return thing == "thing"
}

func (*MockUser) ID() int {
	return 3
}

func (*MockHttp) Post(url string, content any) error {
	return nil
}

var (
	database   *MockDatabase
	request    *MockRequest
	httpClient *MockHttp
)
