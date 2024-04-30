package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, data []byte) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   id,
		Name: filename,
		Data: data,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File: {Name:%s, Id:%v}", f.Name, f.ID)
}
