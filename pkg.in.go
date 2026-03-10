type Params struct {
	// Dir holds the name of the directory holding the scripts.
	// All files in the directory with a .txtar or .txt suffix will be
	// considered as test scripts. By default the current directory is used.
	// Dir is interpreted relative to the current test directory.
	Dir string

	// Files holds a set of script filenames. If Dir is empty and this
	// is non-nil, these files will be used instead of reading
	// a directory.
	Files []string

	// Setup is called, if not nil, to complete any setup required
	// for a test. The WorkDir and Vars fields will have already
	// been initialized and all the files extracted into WorkDir,
	// and Cd will be the same as WorkDir.
	// The Setup function may modify Vars and Cd as it wishes.
	Setup func(*Env) error

	// Condition is called, if not nil, to determine whether a particular
	// condition is true. It's called only for conditions not in the
	// standard set, and may be nil.
	Condition func(cond string) (bool, error)

	// Cmds holds a map of commands available to the script.
	// It will only be consulted for commands not part of the standard set.
	Cmds map[string]func(ts *TestScript, neg bool, args []string)

	// TestWork specifies that working directories should be
	// left intact for later inspection.
	TestWork bool

	// WorkdirRoot specifies the directory within which scripts' work
	// directories will be created. Setting WorkdirRoot implies TestWork=true.
	// If empty, the work directories will be created inside
	// $GOTMPDIR/go-test-script*, where $GOTMPDIR defaults to os.TempDir().
	WorkdirRoot string

	// Deprecated: this option is no longer used.
	IgnoreMissedCoverage bool

	// UpdateScripts specifies that if a `cmp` command fails and its second
	// argument refers to a file inside the testscript file, the command will
	// succeed and the testscript file will be updated to reflect the actual
	// content (which could be stdout, stderr or a real file).
	//
	// The content will be quoted with txtar.Quote if needed;
	// a manual change will be needed if it is not unquoted in the
	// script.
	UpdateScripts bool

	// RequireExplicitExec requires that commands passed to [Main] must be used
	// in test scripts via `exec cmd` and not simply `cmd`. This can help keep
	// consistency across test scripts as well as keep separate process
	// executions explicit.
	RequireExplicitExec bool

	// RequireUniqueNames requires that names in the txtar archive are unique.
	// By default, later entries silently overwrite earlier ones.
	RequireUniqueNames bool

	// ContinueOnError causes a testscript to try to continue in
	// the face of errors. Once an error has occurred, the script
	// will continue as if in verbose mode.
	ContinueOnError bool

	// Deadline, if not zero, specifies the time at which the test run will have
	// exceeded the timeout. It is equivalent to testing.T's Deadline method,
	// and Run will set it to the method's return value if this field is zero.
	Deadline time.Time
}
