package util

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

func (h *Helper) GetIntersection(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	var result []string
	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
		}
	}
	return result

}

func (h *Helper) GetDifference(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true
	}
	var result []string
	for _, item := range a {
		if _, ok := m[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}
