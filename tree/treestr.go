package tree

import (
	"fmt"
	"strings"
)

func ToStr[VertexId comparable, EdgeId comparable](t *Vertex[VertexId, EdgeId]) string {
	visited := map[VertexId]bool{}
	var sb strings.Builder
	visit(t, visited, &sb, 0)
	return sb.String()
}

func visit[VertexId comparable, EdgeId comparable](
	t *Vertex[VertexId, EdgeId], visited map[VertexId]bool, sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat(" ", depth))
	fmt.Fprintf(sb, "%v (%d)", t.Id, len(t.Edges))
	fmt.Fprintln(sb)
	if !visited[t.Id] {
		visited[t.Id] = true
		for _, e := range t.Edges {
			visit(e.To, visited, sb, depth+1)
		}
	}
}
