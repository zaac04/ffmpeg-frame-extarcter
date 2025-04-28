package ffmpeg

type FFProbeOutput struct {
	Format Format `json:"format"`
}

type Format struct {
	Filename       string            `json:"filename"`
	NbStreams      int               `json:"nb_streams"`
	NbPrograms     int               `json:"nb_programs"`
	NbStreamGroups int               `json:"nb_stream_groups"`
	FormatName     string            `json:"format_name"`
	FormatLongName string            `json:"format_long_name"`
	StartTime      string            `json:"start_time"`
	Duration       string            `json:"duration"`
	Size           string            `json:"size"`
	BitRate        string            `json:"bit_rate"`
	ProbeScore     int               `json:"probe_score"`
	Tags           map[string]string `json:"tags"`
}
