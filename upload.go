package reducto

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// ErrEmptyUpload is returned by Upload when the reader yields no bytes.
var ErrEmptyUpload = errors.New("reducto: upload: file is empty")

// ErrUploadTooLarge is returned by Upload when the file exceeds WithMaxUploadSize.
var ErrUploadTooLarge = errors.New("reducto: upload: file exceeds the size limit")

// UploadOptions tune Client.Upload.
type UploadOptions struct {
	// Extension overrides the file extension the server uses to detect the document type, e.g. "pdf".
	Extension string
}

// Upload sends a document to Reducto and returns a reducto:// handle to use as DocumentInput.
//
// An io.ReadSeeker (such as *os.File) is streamed and rewound for each retry. Any other reader
// is buffered in memory first. The multipart body always has a known Content-Length.
//
// POST /upload
func (c *Client) Upload(ctx context.Context, r io.Reader, filename string, uo *UploadOptions, opts ...Option) (*UploadResponse, error) {
	c = c.with(opts)
	src, size, err := uploadSource(r)
	if err != nil {
		return nil, err
	}
	if size == 0 {
		return nil, ErrEmptyUpload
	}
	if c.maxUploadSize > 0 && size > c.maxUploadSize {
		return nil, fmt.Errorf("%w: %d bytes, limit %d", ErrUploadTooLarge, size, c.maxUploadSize)
	}
	head, tail, contentType, err := multipartFrame(filename)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	if uo != nil && uo.Extension != "" {
		q.Set("extension", strings.TrimPrefix(uo.Extension, "."))
	}
	req := request{
		method:      "POST",
		path:        "/upload",
		query:       q,
		contentType: contentType,
		stream: func() (io.Reader, int64, error) {
			file, err := src()
			if err != nil {
				return nil, 0, err
			}
			body := io.MultiReader(bytes.NewReader(head), io.LimitReader(file, size), bytes.NewReader(tail))
			return body, int64(len(head)) + size + int64(len(tail)), nil
		},
	}
	var out UploadResponse
	if err := c.send(ctx, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// uploadSource returns a function that yields the file positioned at its start, plus its size.
func uploadSource(r io.Reader) (func() (io.Reader, error), int64, error) {
	if s, ok := r.(io.ReadSeeker); ok {
		start, err := s.Seek(0, io.SeekCurrent)
		if err == nil {
			var end int64
			if end, err = s.Seek(0, io.SeekEnd); err == nil {
				return func() (io.Reader, error) {
					if _, err := s.Seek(start, io.SeekStart); err != nil {
						return nil, fmt.Errorf("reducto: upload: rewind: %w", err)
					}
					return s, nil
				}, end - start, nil
			}
		}
	}
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, 0, fmt.Errorf("reducto: upload: read: %w", err)
	}
	return func() (io.Reader, error) { return bytes.NewReader(buf), nil }, int64(len(buf)), nil
}

func multipartFrame(filename string) (head, tail []byte, contentType string, err error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if _, err := w.CreateFormFile("file", filename); err != nil {
		return nil, nil, "", err
	}
	head = append([]byte(nil), buf.Bytes()...)
	tail = []byte("\r\n--" + w.Boundary() + "--\r\n")
	return head, tail, w.FormDataContentType(), nil
}

// PresignUpload asks for a reducto:// handle without sending the bytes. PUT the file to
// PresignedURL, then use the handle as DocumentInput.
//
// POST /upload
func (c *Client) PresignUpload(ctx context.Context, uo *UploadOptions, opts ...Option) (*UploadResponse, error) {
	q := url.Values{}
	if uo != nil && uo.Extension != "" {
		q.Set("extension", strings.TrimPrefix(uo.Extension, "."))
	}
	var out UploadResponse
	if err := c.do(ctx, "POST", "/upload", q, nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// UploadFile opens path and uploads it. The extension is taken from the file name.
func (c *Client) UploadFile(ctx context.Context, path string, opts ...Option) (*UploadResponse, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return c.Upload(ctx, f, filepath.Base(path), &UploadOptions{Extension: filepath.Ext(path)}, opts...)
}

// Input returns the upload as a DocumentInput for parse/extract/split requests.
func (u *UploadResponse) Input() DocumentInput {
	return DocumentInputFromUploadResponse(*u)
}
