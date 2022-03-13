package sawwebservice

import (
	"fmt"
	"math"
)

type PointsVector struct {
	index *Rtree
	min   Point
	max   Point
	size  uint64
}

func NewPointsVector(dim int, objs ...Spatial) (*PointsVector, []error) {
	pv := &PointsVector{
		index: NewTree(dim, 1, 32, objs...),
		min:   make(Point, dim),
		max:   make(Point, dim),
		size:  0,
	}
	for i := 0; i < dim; i++ {
		pv.min[i] = math.MaxFloat64
		pv.max[i] = -math.MaxFloat64
	}
	var err []error

	for _, po := range objs {
		poi, ok := po.(*Point)

		if !ok {
			err = append(err, fmt.Errorf("1"))
		} else {
			for i := 0; i < pv.index.Dim; i++ {
				if pv.min[i] > (*poi)[i] {
					pv.min[i] = (*poi)[i]
				}
				if pv.max[i] < (*poi)[i] {
					pv.max[i] = (*poi)[i]
				}
			}
			err = append(err, fmt.Errorf("0"))
		}
	}
	return pv, err
}
func (pv *PointsVector) insertPoint(p Spatial) (bool, error) {
	poi, ok := p.(*Point)
	var err error

	if !ok {
		err = fmt.Errorf("1")
	} else {
		err = fmt.Errorf("0")
	}
	node := pv.index.chooseNode(pv.index.root, entry{poi.Bounds(), nil, poi}, 0)
	if node != nil {
		for _, entrie := range node.entries {
			poi1, ok1 := entrie.obj.(*Point)
			if ok1 {
				flag := false
				for i := 0; i < pv.index.Dim; i++ {
					if (*poi1)[i] != (*poi)[i] {
						flag = true
					}
				}
				if !flag {
					return true, err
				} else {
					for i := 0; i < pv.index.Dim; i++ {
						if pv.min[i] > (*poi)[i] {
							pv.min[i] = (*poi)[i]
						}
						if pv.max[i] < (*poi)[i] {
							pv.max[i] = (*poi)[i]
						}
					}
					pv.index.Insert(poi)
					return false, err
				}

			}
		}
		return false, err
	} else {
		for i := 0; i < pv.index.Dim; i++ {
			if pv.min[i] > (*poi)[i] {
				pv.min[i] = (*poi)[i]
			}
			if pv.max[i] < (*poi)[i] {
				pv.max[i] = (*poi)[i]
			}
		}
		pv.index.Insert(poi)
		return false, err
	}
}
