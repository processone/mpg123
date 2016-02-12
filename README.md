# mpg123

Go wrapper for mpg123 command.

## Usage

For now if can be used to stream content from an URL.

Here is a simple basic example:

	p, err := mpg123.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	p.Play("https://archive.org/download/testmp3testfile/mpthreetest.mp3")
