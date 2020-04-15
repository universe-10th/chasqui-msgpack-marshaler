package msgpack

import (
	"github.com/universe-10th/chasqui/types"
	"github.com/vmihailenco/msgpack/v4"
	"io"
)

// The internal structure tu pass MsgPack objects.
type message struct {
	C   string
	A   types.Args
	KWA types.KWArgs
}

// Retrieves the command of this message, as
// per the interface implementation.
func (msg message) Command() string {
	return msg.C
}

// Retrieves the args of this message, as
// per the interface implementation.
func (msg message) Args() types.Args {
	return msg.A
}

// Retrieves the kwargs of this message, as
// per the interface implementation.
func (msg message) KWArgs() types.KWArgs {
	return msg.KWA
}

// Marshals MsgPack messages around a read-writer.
type MsgPackMessageMarshaler struct {
	encoder *msgpack.Encoder
	decoder *msgpack.Decoder
}

// Receives a MsgPack message from the underlying
// buffer (socket, most likely).
func (marshaler *MsgPackMessageMarshaler) Receive() (types.Message, error, bool) {
	msg := &message{}
	if err := marshaler.decoder.Decode(&msg); err != nil {
		return nil, err, err == io.EOF
	} else {
		return msg, nil, false
	}
}

// Sends a MsgPack message via the underlying buffer
// (socket, most likely).
func (marshaler *MsgPackMessageMarshaler) Send(command string, args types.Args, kwargs types.KWArgs) error {
	return marshaler.encoder.Encode(message{command, args, kwargs})
}

// Creates a new instance of MsgPack marshaler around
// a buffer (socket, most likely).
func (marshaler *MsgPackMessageMarshaler) Create(buffer io.ReadWriter) types.MessageMarshaler {
	return &MsgPackMessageMarshaler{
		encoder: msgpack.NewEncoder(buffer),
		decoder: msgpack.NewDecoder(buffer),
	}
}
