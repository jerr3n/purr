// SPDX-License-Identifier: MPL-2.0
package parsing

import "strings"

type NodeKind int

const (
	Document NodeKind = iota
	Paragraph
	Heading
	CodeBlock
	Text
	Bold
	Italic
	Strikethrough
	CodeSpan
	Image
)

type Node struct {
	Kind     NodeKind
	Children []*Node
	Level    int
	Lang     string
	Literal  string
}

func leading(ln string) (int, string) {
	i := 0
	cols := 0
	for i < len(ln) {
		switch ln[i] {
		case ' ':
			cols++
		case '\t':
			cols += 4 - cols%4
		default:
			return cols, ln[i:]
		}
		i++
	}
	return cols, ""
}

func ParseBlocks(src string) *Node {
	doc := &Node{Kind: Document}
	lns := strings.Split(src, "\n")
	var para *Node
	for i := 0; i < len(lns); i++ {
		ln := lns[i]
		trimmed := strings.TrimLeft(ln, " ")
		switch {
		case strings.TrimSpace(ln) == "":
			para = nil
		case strings.HasPrefix(trimmed, "```"):
			block := &Node{
				Kind: CodeBlock,
				Lang: trimmed[3:],
			}
			var body string = ""
			for i++; i < len(lns) && !strings.HasPrefix(strings.TrimLeft(lns[i], " "), "```"); i++ {
				body += ("\n" + lns[i])
			}
			block.Literal = body
			doc.Children = append(doc.Children, block)
		case strings.HasPrefix(trimmed, "#"):
			lvl := len(trimmed) - len(strings.TrimLeft(trimmed, "#"))
			if lvl <= 6 && (len(trimmed) == lvl || trimmed[lvl] == ' ') {
				para = nil
				doc.Children = append(doc.Children, &Node{
					Kind:    Heading,
					Level:   lvl,
					Literal: strings.TrimSpace(trimmed[lvl:]),
				})
				continue
			}
			fallthrough
		default:
			if para == nil {
				para = &Node{Kind: Paragraph, Literal: strings.TrimSpace(ln)}
				doc.Children = append(doc.Children, para)
			} else {
				para.Literal += "\n" + strings.TrimSpace(ln)
			}
		}
	}
	return doc
}
