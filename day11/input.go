package main

var example = []*Monkey{
	{
		Items:       []int64{79, 98},
		InspectFn:   MulInspectingFn(19),
		TargetingFn: ModuloTargetingFn(23, 2, 3),
	},
	{
		Items:       []int64{54, 65, 75, 74},
		InspectFn:   AddInspectingFn(6),
		TargetingFn: ModuloTargetingFn(19, 2, 0),
	},
	{
		Items:       []int64{79, 60, 97},
		InspectFn:   SquareInspectingFn(),
		TargetingFn: ModuloTargetingFn(13, 1, 3),
	},
	{
		Items:       []int64{74},
		InspectFn:   AddInspectingFn(3),
		TargetingFn: ModuloTargetingFn(17, 0, 1),
	},
}

var input = []*Monkey{
	{
		Items:       []int64{54, 98, 50, 94, 69, 62, 53, 85},
		InspectFn:   MulInspectingFn(13),
		TargetingFn: ModuloTargetingFn(3, 2, 1),
	},
	{
		Items:       []int64{71, 55, 82},
		InspectFn:   AddInspectingFn(2),
		TargetingFn: ModuloTargetingFn(13, 7, 2),
	},
	{
		Items:       []int64{77, 73, 86, 72, 87},
		InspectFn:   AddInspectingFn(8),
		TargetingFn: ModuloTargetingFn(19, 4, 7),
	},
	{
		Items:       []int64{97, 91},
		InspectFn:   AddInspectingFn(1),
		TargetingFn: ModuloTargetingFn(17, 6, 5),
	},
	{
		Items:       []int64{78, 97, 51, 85, 66, 63, 62},
		InspectFn:   MulInspectingFn(17),
		TargetingFn: ModuloTargetingFn(5, 6, 3),
	},
	{
		Items:       []int64{88},
		InspectFn:   AddInspectingFn(3),
		TargetingFn: ModuloTargetingFn(7, 1, 0),
	},
	{
		Items:       []int64{87, 57, 63, 86, 87, 53},
		InspectFn:   SquareInspectingFn(),
		TargetingFn: ModuloTargetingFn(11, 5, 0),
	},
	{
		Items:       []int64{73, 59, 82, 65},
		InspectFn:   AddInspectingFn(6),
		TargetingFn: ModuloTargetingFn(2, 4, 3),
	},
}
