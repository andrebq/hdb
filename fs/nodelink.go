package fs

func (nl *NodeLink) Valid() bool {
	return nl != nil && nl.Name != ""
}
