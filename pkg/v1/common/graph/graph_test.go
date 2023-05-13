package graph

import (
	"testing"

	g "github.com/m8u/gorgonia"
	"github.com/stretchr/testify/require"
)

func TestVisualize(t *testing.T) {
	graph := g.NewGraph()

	a := g.NewVector(graph, g.Float32, g.WithShape(4), g.WithName("test1"), g.WithInit(g.Zeroes()))
	b := g.NewVector(graph, g.Float32, g.WithShape(4), g.WithName("test2"), g.WithInit(g.Zeroes()))

	_, err := g.Mul(a, b)
	require.NoError(t, err)

	Visualize(graph)
}
