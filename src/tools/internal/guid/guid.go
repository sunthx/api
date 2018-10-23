package guid

import (
	"github.com/google/uuid"
)

//New 生成UUID
func New() string {
	newGUID := uuid.Must(uuid.NewRandom())
	return newGUID.String()
}
