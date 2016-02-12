package mpg123

import (
	"io"
	"net/http"
	"os/exec"
)

// Player is an abstraction for audio player using existing operating
// system mpg123 command.
type Player struct {
	State           int32
	currentURL      string
	mpg123          *exec.Cmd
	mpg123Handler   io.WriteCloser
	httpBodyHandler io.ReadCloser
}

const (
	stateStopped = iota
	statePlaying
)

// NewPlayer creates and initialize a new MPG123Player
func NewPlayer() (*Player, error) {
	err := checkMPG123Path()
	p := new(Player)
	return p, err
}

// Play uses mpg123 command to play audio with stream URL.
// TODO: Add ability to play a file
func (p *Player) Play(StreamURL string) (err error) {
	if p.State == statePlaying {
		p.Stop()
		return p.Play(StreamURL)
	}

	// mpg123 setup
	// TODO Reduce mpg123 buffer to be able to stream faster ?
	mpg123 := exec.Command("mpg123", "-q", "-")
	stdin, err := mpg123.StdinPipe()
	if err != nil {
		stdin.Close()
		return
	}

	err = mpg123.Start()
	if err != nil {
		stdin.Close()
		return
	}

	// Setup HTTP stream
	response, err := http.Get(StreamURL)
	if err != nil {
		return
	}

	// Setup mpg123 pipe
	p.currentURL = StreamURL
	p.mpg123 = mpg123
	p.mpg123Handler = stdin
	p.httpBodyHandler = response.Body
	p.State = statePlaying

	// Start streaming
	go streamData(stdin, response.Body)
	return
}

func streamData(stdin io.WriteCloser, body io.ReadCloser) {
	defer stdin.Close()
	defer body.Close()
	io.Copy(stdin, body)
	// We ignore errors for now
	// Errors typically happen when the stream source was disconnected or
	// force closed ou our side to interrupt player
}

// Stop terminates HTPP request and MPG123 command stream.
func (p *Player) Stop() {
	if p.State == stateStopped {
		return
	}

	p.State = stateStopped
	p.httpBodyHandler.Close()
	p.httpBodyHandler = nil

	p.mpg123Handler.Close()
	p.mpg123.Process.Kill() // We force kill to stop sound as fast as possible
	p.mpg123.Wait()         // Wait ensures Zombie process are cleaned up
	p.mpg123Handler = nil
}

func checkMPG123Path() error {
	_, err := exec.LookPath("mpg123")
	return err
}
