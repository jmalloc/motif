package tlv

func (m marshaler) VisitAnonymousTag() {
}

func (m marshaler) VisitContextSpecificTag(t ContextSpecificTag) {
	m.WriteControl(contextSpecificTagForm)
	m.WriteByte(byte(t))
}
