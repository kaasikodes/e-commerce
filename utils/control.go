package utils

// Ternary is a utility function that acts like the ternary operator
func Ternary(condition bool, trueValue, falseValue interface{}) interface{} {
    if condition {
        return trueValue
    }
    return falseValue
}
