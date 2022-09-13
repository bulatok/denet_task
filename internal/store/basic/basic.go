package basic

import (
	"github.com/bulatok/denet_task/pkg/storage"
)

type Basic struct {
	files *files
}

func NewBasic() (*Basic, error) {
	return &Basic{
		files: newFiles(),
	}, nil
}

func (b *Basic) Close() error {
	return nil
}

func (b *Basic) Files() storage.Files {
	return b.files
}
