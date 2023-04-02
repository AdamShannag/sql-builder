package builder

func operationAsString(operation int) string {
	switch operation {
	case GREATER_THAN:
		return ">"
	case GREATER_THAN_OR_EQUAL:
		return ">="
	case SMALLER_THAN:
		return "<"
	case SMALLER_THAN_OR_EQUAL:
		return "<="
	case EQUAL_TO:
		return "="
	case NOT_EQUAL_TO:
		return "<>"
	default:
		return "="
	}
}
