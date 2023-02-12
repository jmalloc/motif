package tlv

type (
	// String1 is a string with a 1 octet length.
	String1 string

	// String2 is a string with a 2 octet length.
	String2 string

	// String4 is a string with a 4 octet length.
	String4 string

	// String8 is a string with an 8 octet length.
	String8 string
)

// AcceptElementVisitor invokes the appropriate method on vis.
func (v String1) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor invokes the appropriate method on vis.
func (v String2) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor invokes the appropriate method on vis.
func (v String4) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptElementVisitor invokes the appropriate method on vis.
func (v String8) AcceptElementVisitor(vis ElementVisitor) { vis.VisitAnonymousElement(v) }

// AcceptValueVisitor invokes the appropriate method on vis.
func (v String1) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString1(v) }

// AcceptValueVisitor invokes the appropriate method on vis.
func (v String2) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString2(v) }

// AcceptValueVisitor invokes the appropriate method on vis.
func (v String4) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString4(v) }

// AcceptValueVisitor invokes the appropriate method on vis.
func (v String8) AcceptValueVisitor(vis ValueVisitor) { vis.VisitString8(v) }
