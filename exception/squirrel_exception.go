package exception

func NewSquirrelToSqlError(err error) error {
	return NewInternalServerError("failed to build sql query", err)
}
