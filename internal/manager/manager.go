package manager

import "github.com/bulatok/denet_task/internal/usecase"

type Manager struct {
	Files FilesManager
}

func NewManager(boot *usecase.BootStrap) *Manager {
	return &Manager{
		Files: NewFilesManager(boot),
	}
}
