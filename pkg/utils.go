package pkg

import "strings"

func CreateSegment(fields ...string) string {
	segmentTemplate := []string{"%d.", "%d.", "%d.", "%d"}
	var s strings.Builder
	for _, field := range fields {
		s.WriteString(field)
		s.WriteString(".")
	}
	for i := len(fields); i < 4; i++ {
		s.WriteString(segmentTemplate[i])
	}
	segment := s.String()
	if len(fields) == 4 {
		return segment[:len(segment)-1]
	}
	return segment
}

func CreateSegmentTest(segments []string) {
	if len(segments) == 0 {
		return
	}
}
