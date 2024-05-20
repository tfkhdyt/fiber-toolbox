package sq

import "github.com/Masterminds/squirrel"

// Psql represents a statement builder configured to use PostgreSQL's dollar-sign syntax for placeholders in squirrel library.
var Psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
