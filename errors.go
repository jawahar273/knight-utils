package error

import (
	"errors"
	"fmt"

	"github.com/jawahar273/knight-utils/core"
	"go.mongodb.org/mongo-driver/mongo"
)

// Capture all mongo db error
// return code and error
func CaptureMongoError(err error) (int, error) {
	if errors.Is(err, ErrInValidEntityID) {
		return core.InvalidRequest, ErrInValidEntityID
	} else if errors.Is(err, mongo.ErrNoDocuments) {

		return core.NotFound, fmt.Errorf("no entity found")
	} else if mongo.IsDuplicateKeyError(err) {
		// FIXME: core.NotAcceptable
		return core.NotAcceptable, fmt.Errorf("duplicate entity")
	} else if mongo.IsNetworkError(err) {
		return core.ServerError, fmt.Errorf("something went wrong")
	} else if mongo.IsTimeout(err) {
		return core.ServerError, fmt.Errorf("something went wrong")
	}
	return 0, nil
}
