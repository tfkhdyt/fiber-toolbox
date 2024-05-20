package exception

// NewSquirrelToSqlError creates a new internal server error specifically for
// situations where building an SQL query using the Squirrel library fails.
// It wraps the provided error with a message indicating the failure.
//
// Parameters:
//   - err: The original error encountered while building the SQL query.
//
// Returns:
//   - error: A new fiber.Error with status code 500 (Internal Server Error) and a message
//     indicating the failure to build the SQL query, including the original error message as the cause.
func NewSquirrelToSqlError(err error) error {
	return NewInternalServerError("failed to build sql query", err)
}
