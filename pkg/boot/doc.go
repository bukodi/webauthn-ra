package boot

// Package boot executes an initial load.
// The load sequence is:
// 1. Load config form file, environment variables and command line arguments.
// 2. Init loggers.
// 2.1. If the config load was failed in step 1., logs the error to the stderr, end exits
// 2.2. If the config load was successful, use it for log initialization, and logs the first messages
// 3. Init repository
// 4. Start listeners.
// 5. Start process scheduler.
