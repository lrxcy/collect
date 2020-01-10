package audiohandler

type Handler interface {
	HandleReader(ReadCloser)
	HandleWriter(WriteCloser)
}

type ReadCloser interface {
	Closer
	Alive
	Read(*Packet)
}

type WriteCloser interface {
	Closer
	Alive
	CalcTime
	Write(*Packet) error
}

type Closer interface {
	Info() Info
	Close(error)
}

type Info struct {
	Key   string
	URL   string
	UID   string
	Inter bool
}

type Packet struct {
	IsAudio    bool
	IsVideo    bool
	IsMetadata bool
	TimeStamp  uint32 // dts
	StreamID   uint32
	Header     PacketHeader
	Data       []byte
}

type PacketHeader interface {
}

type CalcTime interface {
	CalcBaseTimestamp()
}

type Alive interface {
	Alive() bool
}

type GetWriter interface {
	GetWriter(Info) WriteCloser
}
