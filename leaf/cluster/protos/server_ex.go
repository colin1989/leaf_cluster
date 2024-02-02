package protos

func (x *Server) CheckType(typ ServerType) bool {
	return x.Typ == typ
}
