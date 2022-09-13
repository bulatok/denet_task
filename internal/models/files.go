package models

import (
	"github.com/bulatok/denet_task/pkg/logger"
	"github.com/bulatok/denet_task/pkg/types"
)

// FileOut for listing all files stored by app
type FileOut struct {
	Name                string `json:"name"`
	EncodedSizeBytes    int64  `json:"encoded_size_in_bytes"`
	EncodedSizeMB       int64  `json:"encoded_size_MB"`
	OriginalSizeMB      int64  `json:"original_size_MB"`
	OriginalSizeInBytes int64  `json:"original_size_in_bytes"`
}

type File struct {
	Name           string
	ContentType    string
	data           []byte
	offset         int64
	underLineMB    int64
	underLineBytes int64 // this field is unnecessary, only for All() method
}

func (f *File) Bytes() []byte {
	return f.data
}

// GetUnderLineSizeMB returns the size of original file in MB
func (f *File) GetUnderLineSizeMB() int64 {
	return f.underLineMB
}

// GetCurrentSizeMB returns the size of encoded file in MB
func (f *File) GetCurrentSizeMB() int64 {
	return types.ToMB(int64(len(f.data)))
}

// GetUnderLineSizeBytes returns the size of encoded file in Bytes
func (f *File) GetUnderLineSizeBytes() int64 {
	return f.underLineBytes
}

// SetData data setter
func (f *File) SetData(b []byte) {
	(*f).data = b
}

// SetContentType content setter
func (f *File) SetContentType(ct string) {
	(*f).ContentType = ct
}

// Encode encodes file by adding 0 bytes
func (f *File) Encode() error {
	defer func() {
		if err := recover(); err != nil {
			logger.Warn("caught panic, while Encoding file")
		}
	}()
	if f.data == nil {
		return ErrNoData
	}

	// setting underline values
	size := int64(len(f.data))
	(*f).offset = types.GetOffsetMB(size)
	(*f).underLineMB = types.ToMB(size)
	(*f).underLineBytes = size

	// adding zero bytes
	times := types.MB - (*f).offset
	bytesZeroes := make([]byte, times)
	(*f).data = append((*f).data, bytesZeroes...)

	return nil
}

// Decode decodes file by removing redundant 0 bytes
func (f *File) Decode() error {
	if f.data == nil {
		return ErrNoData
	}

	till := (*f).underLineMB*types.MB + (*f).offset
	(*f).data = (*f).data[:till]
	return nil
}
