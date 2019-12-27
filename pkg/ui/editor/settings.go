package editor

// DefaultLocalSettings returns the default local settings
// Note that filetype is a local only option
func DefaultLocalSettings() map[string]interface{} {
	return map[string]interface{}{
		"autoindent":     true,
		"autosave":       false,
		"basename":       false,
		"colorcolumn":    float64(0),
		"cursorline":     false,
		"eofnewline":     false,
		"fastdirty":      true,
		"fileformat":     "unix",
		"filetype":       "Unknown",
		"hidehelp":       false,
		"ignorecase":     false,
		"indentchar":     " ",
		"keepautoindent": false,
		"matchbrace":     false,
		"matchbraceleft": false,
		"rmtrailingws":   false,
		"ruler":          false,
		"savecursor":     false,
		"saveundo":       false,
		"scrollbar":      false,
		"scrollmargin":   float64(3),
		"scrollspeed":    float64(2),
		"softwrap":       true,
		"smartpaste":     true,
		"splitbottom":    true,
		"splitright":     true,
		"statusline":     true,
		"syntax":         true,
		"tabmovement":    false,
		"tabsize":        float64(4),
		"tabstospaces":   false,
		"useprimary":     true,
	}
}
