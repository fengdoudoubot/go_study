package model

type student struct {
	name string
	score float64
}

//工厂模式创建实例
func NewStudent(n string, s float64) *student{
	return &student{
		name : n, 
		score : s,
	}
}

func (s *student) GetScore() float64{
	return s.score
}