package default_implementation

import "testing"

func TestNewGameDirector(t *testing.T) {
	director1 := NewGameDirector(&AGame{})
	director2 := NewGameDirector(&BAame{})

	director1.Play()
	director2.Play()
}
