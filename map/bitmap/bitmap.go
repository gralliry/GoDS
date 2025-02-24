package bitmap

type Bitmap struct {
	size int
	data []byte
}

func New(cap int) *Bitmap {
	size := (cap + 7) / 8
	return &Bitmap{size: size, data: make([]byte, size)}
}

func (b *Bitmap) Set(index int) {
	b.data[index/8] |= 1 << (index % 8)
}

func (b *Bitmap) Unset(index int) {
	b.data[index/8] &= ^(1 << (index % 8))
}

func (b *Bitmap) Exist(index int) bool {
	return b.data[index/8]&(1<<(index%8)) != 0
}

func (b *Bitmap) Clear() {
	b.data = make([]byte, b.size)
}

func (b *Bitmap) Cap() int {
	return b.size * 8
}
