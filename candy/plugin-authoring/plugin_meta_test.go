package authoring

import (
	"context"
	"testing"

	pb "github.com/opencharly/spec/proto"
)

// TestNewMeta_DeclaresBoxParent proves Describe advertises every authoring command word
// with CommandParent=="box" — the capability IDENTITY (`command:<word>:box`) that charly
// keys the provider at. This FAILS on the pre-change code, which declared no command_parent
// and relied on a Go-interface sniff: the declaration is the thing under test, and it is
// asserted over the wire (Describe), not inferred.
func TestNewMeta_DeclaresBoxParent(t *testing.T) {
	caps, err := NewMeta().Describe(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Describe: %v", err)
	}
	got := map[string]bool{}
	for _, c := range caps.GetProvided() {
		if c.GetClass() != "command" {
			t.Errorf("capability %s:%s is not class command", c.GetClass(), c.GetWord())
		}
		if c.GetCommandParent() != authoringCommandParent {
			t.Errorf("command:%s must declare command_parent %q (its identity), got %q", c.GetWord(), authoringCommandParent, c.GetCommandParent())
		}
		got[c.GetWord()] = true
	}
	for _, want := range authoringCommandWords {
		if !got[want] {
			t.Errorf("Describe missing command:%s (got %v)", want, got)
		}
	}
	if len(caps.GetProvided()) != len(authoringCommandWords) {
		t.Errorf("want %d command capabilities, got %d", len(authoringCommandWords), len(caps.GetProvided()))
	}
}
