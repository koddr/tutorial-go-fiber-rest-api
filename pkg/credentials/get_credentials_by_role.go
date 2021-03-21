package credentials

import "fmt"

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case "admin":
		// Admin credentials (all access).
		credentials = []string{
			BookCreate, BookUpdate, BookDelete,
		}
	case "moderator":
		// Moderator credentials (only book creation and update).
		credentials = []string{
			BookCreate, BookUpdate,
		}
	case "user":
		// Simple user credentials (only book creation).
		credentials = []string{
			BookCreate,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
