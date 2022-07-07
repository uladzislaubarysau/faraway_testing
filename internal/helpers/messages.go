package helpers

import (
	"encoding/binary"
	"io"
	"net"
)

const sizeOfLength = 8

// To prevent closing the connection between the client and the server
// after sending messages, we will add meta information about the
// length of the transmitted message. This information will be added
// as the first 8 bytes to the forwarded message itself

func WriteMessage(conn net.Conn, message []byte) error {
	buf := make([]byte, sizeOfLength+len(message))
	binary.BigEndian.PutUint64(buf, uint64(len(message)))

	copy(buf[sizeOfLength:], message)
	_, err := conn.Write(buf)

	return err
}

func ReadMessage(conn net.Conn) ([]byte, error) {
	var buf [sizeOfLength]byte
	if _, err := io.ReadFull(conn, buf[:]); err != nil {
		return nil, err
	}

	n := binary.BigEndian.Uint64(buf[:])
	data := make([]byte, n)
	_, err := io.ReadFull(conn, data)

	return data, err
}
