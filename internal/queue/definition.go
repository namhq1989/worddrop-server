package queue

var TypeNames = struct {
	FetchNews         string
	GenerateTextAudio string
}{
	FetchNews:         "content.fetchNews",
	GenerateTextAudio: "tts.generateTextAudio",
}
