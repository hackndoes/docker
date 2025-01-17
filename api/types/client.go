package types

import (
	"bufio"
	"io"
	"net"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/pkg/ulimit"
)

// ContainerAttachOptions holds parameters to attach to a container.
type ContainerAttachOptions struct {
	ContainerID string
	Stream      bool
	Stdin       bool
	Stdout      bool
	Stderr      bool
}

// ContainerCommitOptions holds parameters to commit changes into a container.
type ContainerCommitOptions struct {
	ContainerID    string
	RepositoryName string
	Tag            string
	Comment        string
	Author         string
	Changes        []string
	Pause          bool
	Config         *container.Config
}

// ContainerExecInspect holds information returned by exec inspect.
type ContainerExecInspect struct {
	ExecID      string
	ContainerID string
	Running     bool
	ExitCode    int
}

// ContainerListOptions holds parameters to list containers with.
type ContainerListOptions struct {
	Quiet  bool
	Size   bool
	All    bool
	Latest bool
	Since  string
	Before string
	Limit  int
	Filter filters.Args
}

// ContainerLogsOptions holds parameters to filter logs with.
type ContainerLogsOptions struct {
	ContainerID string
	ShowStdout  bool
	ShowStderr  bool
	Since       string
	Timestamps  bool
	Follow      bool
	Tail        string
}

// ContainerRemoveOptions holds parameters to remove containers.
type ContainerRemoveOptions struct {
	ContainerID   string
	RemoveVolumes bool
	RemoveLinks   bool
	Force         bool
}

// CopyToContainerOptions holds information
// about files to copy into a container
type CopyToContainerOptions struct {
	ContainerID               string
	Path                      string
	Content                   io.Reader
	AllowOverwriteDirWithFile bool
}

// EventsOptions hold parameters to filter events with.
type EventsOptions struct {
	Since   string
	Until   string
	Filters filters.Args
}

// HijackedResponse holds connection information for a hijacked request.
type HijackedResponse struct {
	Conn   net.Conn
	Reader *bufio.Reader
}

// Close closes the hijacked connection and reader.
func (h *HijackedResponse) Close() {
	h.Conn.Close()
}

// CloseWriter is an interface that implement structs
// that close input streams to prevent from writing.
type CloseWriter interface {
	CloseWrite() error
}

// CloseWrite closes a readWriter for writing.
func (h *HijackedResponse) CloseWrite() error {
	if conn, ok := h.Conn.(CloseWriter); ok {
		return conn.CloseWrite()
	}
	return nil
}

// ImageBuildOptions holds the information
// necessary to build images.
type ImageBuildOptions struct {
	Tags           []string
	SuppressOutput bool
	RemoteContext  string
	NoCache        bool
	Remove         bool
	ForceRemove    bool
	PullParent     bool
	Isolation      string
	CPUSetCPUs     string
	CPUSetMems     string
	CPUShares      int64
	CPUQuota       int64
	CPUPeriod      int64
	Memory         int64
	MemorySwap     int64
	CgroupParent   string
	ShmSize        string
	Dockerfile     string
	Ulimits        []*ulimit.Ulimit
	BuildArgs      []string
	AuthConfigs    map[string]AuthConfig
	Context        io.Reader
}

// ImageBuildResponse holds information
// returned by a server after building
// an image.
type ImageBuildResponse struct {
	Body   io.ReadCloser
	OSType string
}

// ImageCreateOptions holds information to create images.
type ImageCreateOptions struct {
	// Parent is the image to create this image from
	Parent string
	// Tag is the name to tag this image
	Tag string
	// RegistryAuth is the base64 encoded credentials for this server
	RegistryAuth string
}

// ImageImportOptions holds information to import images from the client host.
type ImageImportOptions struct {
	// Source is the data to send to the server to create this image from
	Source io.Reader
	// Source is the name of the source to import this image from
	SourceName string
	// RepositoryName is the name of the repository to import this image
	RepositoryName string
	// Message is the message to tag the image with
	Message string
	// Tag is the name to tag this image
	Tag string
	// Changes are the raw changes to apply to the image
	Changes []string
}

// ImageListOptions holds parameters to filter the list of images with.
type ImageListOptions struct {
	MatchName string
	All       bool
	Filters   filters.Args
}

// ImagePullOptions holds information to pull images.
type ImagePullOptions struct {
	ImageID string
	Tag     string
	// RegistryAuth is the base64 encoded credentials for this server
	RegistryAuth string
}

//ImagePushOptions holds information to push images.
type ImagePushOptions ImagePullOptions

// ImageRemoveOptions holds parameters to remove images.
type ImageRemoveOptions struct {
	ImageID       string
	Force         bool
	PruneChildren bool
}

// ImageSearchOptions holds parameters to search images with.
type ImageSearchOptions struct {
	Term         string
	RegistryAuth string
}

// ImageTagOptions holds parameters to tag an image
type ImageTagOptions struct {
	ImageID        string
	RepositoryName string
	Tag            string
	Force          bool
}

// ResizeOptions holds parameters to resize a tty.
// It can be used to resize container ttys and
// exec process ttys too.
type ResizeOptions struct {
	ID     string
	Height int
	Width  int
}

// VersionResponse holds version information for the client and the server
type VersionResponse struct {
	Client *Version
	Server *Version
}

// ServerOK return true when the client could connect to the docker server
// and parse the information received. It returns false otherwise.
func (v VersionResponse) ServerOK() bool {
	return v.Server != nil
}
