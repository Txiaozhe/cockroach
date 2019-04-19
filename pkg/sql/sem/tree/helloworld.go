package tree

import "fmt"

type HelloworldMode int

type Helloworld struct {
	Mode HelloworldMode
}

var _ Statement = &Helloworld{}

const (
	HelloworldModeSay HelloworldMode = iota
	HelloworldModeSmile
)

// StatementType
func (node *Helloworld) StatementType() StatementType { return Ack }

// StatementTag
func (node *Helloworld) StatementTag() string { return "HELLOWORLD" }

// Format
func (node *Helloworld) Format(fc *FmtCtx) {
	fc.WriteString("HELLOWORLD ")
	switch node.Mode {
	case HelloworldModeSay:
		fc.WriteString("SAY")
	case HelloworldModeSmile:
		fc.WriteString("SMILE")
	default:
		panic(fmt.Errorf("Unknown HELLOWORLD mode %v!", node.Mode))
	}
}

func (node *Helloworld) String() string {
	return AsString(node)
}
