package client

import (
	"errors"
	"fmt"
)

var ErrNoRootKeys = errors.New("tuf: no root keys found in local meta store")

type ErrMissingRemoteMetadata struct {
	Name string
}

func (e ErrMissingRemoteMetadata) Error() string {
	return fmt.Sprintf("tuf: missing remote metadata %s", e.Name)
}

type ErrDownloadFailed struct {
	File string
	Err  error
}

func (e ErrDownloadFailed) Error() string {
	return fmt.Sprintf("tuf: failed to download %s: %s", e.File, e.Err)
}

type ErrNotFound struct {
	File string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("tuf: file not found: %s", e.File)
}

func IsNotFound(err error) bool {
	_, ok := err.(ErrNotFound)
	return ok
}

type ErrWrongSize struct {
	File     string
	Actual   int64
	Expected int64
}

func (e ErrWrongSize) Error() string {
	return fmt.Sprintf("tuf: unexpected file size: %s (expected %d bytes, got %d bytes)", e.File, e.Expected, e.Actual)
}

type ErrLatestSnapshot struct {
	Version int
}

func (e ErrLatestSnapshot) Error() string {
	return fmt.Sprintf("tuf: the local snapshot version (%d) is the latest", e.Version)
}

func IsLatestSnapshot(err error) bool {
	_, ok := err.(ErrLatestSnapshot)
	return ok
}

type ErrUnknownTarget struct {
	Name string
}

func (e ErrUnknownTarget) Error() string {
	return fmt.Sprintf("tuf: unknown target file: %s", e.Name)
}

type ErrMetaTooLarge struct {
	Name string
	Size int64
}

func (e ErrMetaTooLarge) Error() string {
	return fmt.Sprintf("tuf: %s size %d bytes greater than maximum %d bytes", e.Name, e.Size, maxMetaSize)
}

type ErrExpiredMeta struct {
	Name string
}

func (e ErrExpiredMeta) Error() string {
	return fmt.Sprintf("tuf: %s has expired", e.Name)
}