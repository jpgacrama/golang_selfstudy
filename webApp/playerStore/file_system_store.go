package playerstore

import (
	"golang_selfstudy/webApp/player"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []player.Player {
	return nil
}
