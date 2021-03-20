package credentials

var (
	// Define credentials for Book model.
	BookCredentials = map[string][]string{
		"full":        {"book:create", "book:update", "book:delete"}, // all credentials
		"only_create": {"book:create"},                               // only create
	}
)
