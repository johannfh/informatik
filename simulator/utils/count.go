package utils

import (
	"strings"

	"github.com/johannfh/go-utils/assert"
)

// example: CountMap["Item/Apple"]++
type CountMap map[string]int

func (e CountMap) Increment(key string) {
	assert.Assert(key != "", "empty counter key")
	e[key]++
}

type CountNode struct {
	Label string `json:"label"`
	Count int    `json:"count"`

	SubNodes []*CountNode `json:"subNodes,omitempty"`
}

func (ecn *CountNode) Append(namespace []string, count int) {
	ecn.Count += count
	if len(namespace) == 0 {
		return
	}

	label := namespace[0]
	subNs := namespace[1:]

	// search the subnodes for the label
	var subNode *CountNode
	for _, sn := range ecn.SubNodes {
		if sn.Label == label {
			subNode = sn
		}
	}

	// when no subnode with the label is found
	// create a new node with the label
	if subNode == nil {
		subNode = &CountNode{
			Label: label,
		}
	}
	ecn.SubNodes = append(ecn.SubNodes, subNode)

	// repeat for every node in the namespace downwards
	subNode.Append(subNs, count)
}

// extract all the counts collected by cMap
// into subNodes of this node
func (cn *CountNode) ExtractFromCountMap(cMap CountMap) {
	for key, val := range cMap {
		assert.NotEmpty(key, "empty key in countmap")
		ns := strings.Split(key, "/")
		assert.Assert(len(ns) > 0, "no namespace found")
		cn.Append(ns, val)
	}
}
