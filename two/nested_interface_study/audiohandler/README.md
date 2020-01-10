# flow
```go
-- 定義初始化Server會有主要兩個方法，handler以及getter
type Server struct {
	handler Handler
	getter  GetWriter
}

    ---- Server.handler 為一個接口，具有兩個方法`HandleReader`以及`HandleWriter`
    type Handler interface {
    	HandleReader(ReadCloser)
    	HandleWriter(WriteCloser)
    }

        -------- HandleReader(ReadCloser)為一個輸入為接口`ReadCloser`的方法
        type ReaderCloser interface {
            Closer
            Alive
            Read(*Packet)
        }
            -------- Closer 為一個接口，具有方法`Info() Info`以及`Close(error)`
            type Closer interface {
                Info() Info
                Close(error)
            }
                -------- `Info() Info`為一個回傳資料結構為`Info`的方法
                type Info struct {
                    Key string
                    URL string
                    UID string
                    Inter bool
                }

            -------- Alive 為一個接口，具有方法`Alive() bool`
            type Alive interface {
                Alive() bool
            }

            -------- Read(*Packet) 為一個輸入為指標結構`Packet`的方法
            # note: `PacketHeader`為一個任意接口 type PacketHeadert interface{}
            type Packet struct {
                IsAudio     bool
                IsVideo     bool
                IsMetadata  bool
                TimeStamp   uint32
                StreamID    uint32
                Header      PacketHeader
                Data        []byte
            }

        -------- HandleWriter(WriteCloser)為一個輸入接口`WritCloser`的方法
        # note: `Closer`以及`Alive`同之前定義的接口
        type WriteCloser interface {
            Closer
            Alive
            CalcTime
            Write(*Packet) error
        }
            -------- CalcTime為一個具有`CalcBaseTimestamp()`方法的接口
            type CalcTime interface {
                CalcBaseTimestamp()
            }
            
            -------- Write(*Packet) error 為一個輸入為指標結構`Packet`的方法
            # note: `Packet`同之前定義的結構

    -------- GetWriter 為一個接口，且定義方法`GetWriter(Info) WriteCloser`
    # note: `Info`以及`WriteCloser`同之前定義的結構
    type GetWriter interface {
        GetWriter(Info) WriteCloser
    }


```


# refer:
- https://github.com/gwuhaolin/livego