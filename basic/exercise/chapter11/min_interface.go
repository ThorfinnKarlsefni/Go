package chapter11

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
	Swap(i, j int)
}

func Min(data Miner) interface{} {
	min := data.ElemIx(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemIx(i)
		} else {
			data.Swap(i, i-1)
		}
	}
	return min
}

type InArray []int

func (p InArray) Len() int                  { return len(p) }
func (p InArray) ElemIx(ix int) interface{} { return p[ix] }
func (p InArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p InArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int                  { return len(p) }
func (p StringArray) ElemIx(ix int) interface{} { return p[ix] }
func (p StringArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
