package sawwebservice

import "rtreego"

type PointsVector struct {
	index rtreego.Rtree
	min   Point
	max   Point
	size  uint64
}
