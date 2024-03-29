// Its better to use panic, cos panics are recoverable also they will execute defers unlike fatals
// ensure you recover your code from panic, so that app does not crash

// os.Exit() is not recoverable, and does not execute defers (find out more about this later)
