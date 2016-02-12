# mpg123

Go wrapper for mpg123 command.

The library is simple as it just rely on having the mpg123
command-line tool available. It is less flexible than relying on a
binding to libmpg123, but it is pure go and easier to manage, deploy,
maintain.

## Usage

For now, it can be used to stream content from an URL.

Here is a simple basic example:

	p, err := mpg123.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	p.Play("https://archive.org/download/testmp3testfile/mpthreetest.mp3")
