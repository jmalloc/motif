package tlv

type (
	// Signed1 is a signed 1 octet signed integer.
	Signed1 int8

	// Signed2 is a signed 2 octet signed integer.
	Signed2 int16

	// Signed4 is a signed 4 octet signed integer.
	Signed4 int32

	// Signed8 is a signed 8 octet signed integer.
	Signed8 int64
)

// AcceptElementVisitor calls the appropriate method on vis.
func (v Signed1) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor calls the appropriate method on vis.
func (v Signed2) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor calls the appropriate method on vis.
func (v Signed4) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor calls the appropriate method on vis.
func (v Signed8) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptValueVisitor calls the appropriate method on vis.
func (v Signed1) AcceptValueVisitor(vis ValueVisitor) { vis.VisitSigned1(v) }

// AcceptValueVisitor calls the appropriate method on vis.
func (v Signed2) AcceptValueVisitor(vis ValueVisitor) { vis.VisitSigned2(v) }

// AcceptValueVisitor calls the appropriate method on vis.
func (v Signed4) AcceptValueVisitor(vis ValueVisitor) { vis.VisitSigned4(v) }

// AcceptValueVisitor calls the appropriate method on vis.
func (v Signed8) AcceptValueVisitor(vis ValueVisitor) { vis.VisitSigned8(v) }
