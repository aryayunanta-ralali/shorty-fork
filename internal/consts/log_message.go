// Package consts
package consts

const (
	// LogMessageSuccess const
	LogMessageSuccess = `%v success`

	// LogMessageFailedToValidateRequestBody const
	LogMessageFailedToValidateRequestBody = `failed to validate request body, err: %v`

	// LogMessageDBFailedToStore const
	LogMessageDBFailedToStore = `failed to store data into the database, err: %v`

	// LogMessageDBFailedFetching const
	LogMessageDBFailedFetching = `failed fetching %v data from the database, err: %v`

	// LogMessageDBDataNotFound const
	LogMessageDBDataNotFound = `%v data is not found in the database`

	// LogMessageDBFailedToUpdate const
	LogMessageDBFailedToUpdate = `failed to update data into the database, err: %v`
)
