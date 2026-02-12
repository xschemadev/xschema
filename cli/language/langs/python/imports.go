package python

import (
	"regexp"
	"slices"
	"sort"
	"strings"
)

// MergeImports dedupes and formats Python imports following PEP 8 conventions.
// Import order: standard library, third-party, local (relative) imports.
func MergeImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	// Track "from X import Y" style imports
	fromImports := make(map[string][]string)

	// Track "import X" and "import X as Y" style imports
	moduleImports := make(map[string]string) // module -> alias (empty string if no alias)

	// Regex patterns
	fromImportRe := regexp.MustCompile(`^from\s+(\S+)\s+import\s+(.+)$`)
	importAsRe := regexp.MustCompile(`^import\s+(\S+)(?:\s+as\s+(\w+))?$`)

	for _, imp := range imports {
		imp = strings.TrimSpace(imp)
		if imp == "" {
			continue
		}

		// Handle "from X import Y, Z" or "from X import Y as Z"
		if matches := fromImportRe.FindStringSubmatch(imp); matches != nil {
			module := matches[1]
			namesStr := matches[2]

			// Parse individual names (handles "name as alias" syntax)
			names := parseImportNames(namesStr)
			fromImports[module] = append(fromImports[module], names...)
			continue
		}

		// Handle "import X" or "import X as Y"
		if matches := importAsRe.FindStringSubmatch(imp); matches != nil {
			module := matches[1]
			alias := ""
			if len(matches) > 2 {
				alias = matches[2]
			}
			// Only store if not already present, or if this one has an alias
			if existing, ok := moduleImports[module]; !ok || (alias != "" && existing == "") {
				moduleImports[module] = alias
			}
			continue
		}
	}

	// Dedupe and sort names within each module
	for module, names := range fromImports {
		slices.Sort(names)
		fromImports[module] = slices.Compact(names)
	}

	// Categorize imports
	stdlibFromImports := make(map[string][]string)
	thirdPartyFromImports := make(map[string][]string)
	localFromImports := make(map[string][]string)

	stdlibModuleImports := make(map[string]string)
	thirdPartyModuleImports := make(map[string]string)
	localModuleImports := make(map[string]string)

	for module, names := range fromImports {
		if strings.HasPrefix(module, ".") {
			localFromImports[module] = names
		} else if isStdlibModule(module) {
			stdlibFromImports[module] = names
		} else {
			thirdPartyFromImports[module] = names
		}
	}

	for module, alias := range moduleImports {
		if strings.HasPrefix(module, ".") {
			localModuleImports[module] = alias
		} else if isStdlibModule(module) {
			stdlibModuleImports[module] = alias
		} else {
			thirdPartyModuleImports[module] = alias
		}
	}

	// Build result following PEP 8 order
	var sections [][]string

	// 1. Standard library imports
	stdlibLines := buildImportSection(stdlibFromImports, stdlibModuleImports)
	if len(stdlibLines) > 0 {
		sections = append(sections, stdlibLines)
	}

	// 2. Third-party imports
	thirdPartyLines := buildImportSection(thirdPartyFromImports, thirdPartyModuleImports)
	if len(thirdPartyLines) > 0 {
		sections = append(sections, thirdPartyLines)
	}

	// 3. Local imports
	localLines := buildImportSection(localFromImports, localModuleImports)
	if len(localLines) > 0 {
		sections = append(sections, localLines)
	}

	// Join sections with blank lines
	var result []string
	for i, section := range sections {
		if i > 0 {
			result = append(result, "")
		}
		result = append(result, section...)
	}

	return strings.Join(result, "\n")
}

func parseImportNames(namesStr string) []string {
	// Handle: "name1, name2, name3 as alias"
	// Also handles multi-line imports in parentheses (though we flatten them)
	namesStr = strings.Trim(namesStr, "()")

	var names []string
	parts := strings.Split(namesStr, ",")
	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name != "" {
			names = append(names, name)
		}
	}
	return names
}

func buildImportSection(fromImports map[string][]string, moduleImports map[string]string) []string {
	var lines []string

	// First add "import X" style imports
	var modules []string
	for m := range moduleImports {
		modules = append(modules, m)
	}
	sort.Strings(modules)

	for _, module := range modules {
		alias := moduleImports[module]
		if alias != "" {
			lines = append(lines, "import "+module+" as "+alias)
		} else {
			lines = append(lines, "import "+module)
		}
	}

	// Then add "from X import Y" style imports
	var fromModules []string
	for m := range fromImports {
		fromModules = append(fromModules, m)
	}
	sort.Strings(fromModules)

	for _, module := range fromModules {
		names := fromImports[module]
		if len(names) == 0 {
			continue
		}
		lines = append(lines, "from "+module+" import "+strings.Join(names, ", "))
	}

	return lines
}

// isStdlibModule checks if a module is part of Python's standard library.
// This is a subset of common stdlib modules - not exhaustive but covers common cases.
func isStdlibModule(module string) bool {
	// Get the top-level module name
	topLevel := module
	if idx := strings.Index(module, "."); idx != -1 {
		topLevel = module[:idx]
	}

	return stdlibModules[topLevel]
}

var stdlibModules = map[string]bool{
	// Built-in modules
	"__future__":   true,
	"abc":          true,
	"aifc":         true,
	"argparse":     true,
	"array":        true,
	"ast":          true,
	"asyncio":      true,
	"atexit":       true,
	"base64":       true,
	"bdb":          true,
	"binascii":     true,
	"binhex":       true,
	"bisect":       true,
	"builtins":     true,
	"bz2":          true,
	"calendar":     true,
	"cgi":          true,
	"cgitb":        true,
	"chunk":        true,
	"cmath":        true,
	"cmd":          true,
	"code":         true,
	"codecs":       true,
	"codeop":       true,
	"collections":  true,
	"colorsys":     true,
	"compileall":   true,
	"concurrent":   true,
	"configparser": true,
	"contextlib":   true,
	"contextvars":  true,
	"copy":         true,
	"copyreg":      true,
	"cProfile":     true,
	"crypt":        true,
	"csv":          true,
	"ctypes":       true,
	"curses":       true,
	"dataclasses":  true,
	"datetime":     true,
	"dbm":          true,
	"decimal":      true,
	"difflib":      true,
	"dis":          true,
	"distutils":    true,
	"doctest":      true,
	"email":        true,
	"encodings":    true,
	"enum":         true,
	"errno":        true,
	"faulthandler": true,
	"fcntl":        true,
	"filecmp":      true,
	"fileinput":    true,
	"fnmatch":      true,
	"fractions":    true,
	"ftplib":       true,
	"functools":    true,
	"gc":           true,
	"getopt":       true,
	"getpass":      true,
	"gettext":      true,
	"glob":         true,
	"graphlib":     true,
	"grp":          true,
	"gzip":         true,
	"hashlib":      true,
	"heapq":        true,
	"hmac":         true,
	"html":         true,
	"http":         true,
	"idlelib":      true,
	"imaplib":      true,
	"imghdr":       true,
	"imp":          true,
	"importlib":    true,
	"inspect":      true,
	"io":           true,
	"ipaddress":    true,
	"itertools":    true,
	"json":         true,
	"keyword":      true,
	"lib2to3":      true,
	"linecache":    true,
	"locale":       true,
	"logging":      true,
	"lzma":         true,
	"mailbox":      true,
	"mailcap":      true,
	"marshal":      true,
	"math":         true,
	"mimetypes":    true,
	"mmap":         true,
	"modulefinder": true,
	"multiprocessing": true,
	"netrc":        true,
	"nis":          true,
	"nntplib":      true,
	"numbers":      true,
	"operator":     true,
	"optparse":     true,
	"os":           true,
	"ossaudiodev":  true,
	"pathlib":      true,
	"pdb":          true,
	"pickle":       true,
	"pickletools":  true,
	"pipes":        true,
	"pkgutil":      true,
	"platform":     true,
	"plistlib":     true,
	"poplib":       true,
	"posix":        true,
	"posixpath":    true,
	"pprint":       true,
	"profile":      true,
	"pstats":       true,
	"pty":          true,
	"pwd":          true,
	"py_compile":   true,
	"pyclbr":       true,
	"pydoc":        true,
	"queue":        true,
	"quopri":       true,
	"random":       true,
	"re":           true,
	"readline":     true,
	"reprlib":      true,
	"resource":     true,
	"rlcompleter":  true,
	"runpy":        true,
	"sched":        true,
	"secrets":      true,
	"select":       true,
	"selectors":    true,
	"shelve":       true,
	"shlex":        true,
	"shutil":       true,
	"signal":       true,
	"site":         true,
	"smtpd":        true,
	"smtplib":      true,
	"sndhdr":       true,
	"socket":       true,
	"socketserver": true,
	"spwd":         true,
	"sqlite3":      true,
	"ssl":          true,
	"stat":         true,
	"statistics":   true,
	"string":       true,
	"stringprep":   true,
	"struct":       true,
	"subprocess":   true,
	"sunau":        true,
	"symtable":     true,
	"sys":          true,
	"sysconfig":    true,
	"syslog":       true,
	"tabnanny":     true,
	"tarfile":      true,
	"telnetlib":    true,
	"tempfile":     true,
	"termios":      true,
	"test":         true,
	"textwrap":     true,
	"threading":    true,
	"time":         true,
	"timeit":       true,
	"tkinter":      true,
	"token":        true,
	"tokenize":     true,
	"tomllib":      true,
	"trace":        true,
	"traceback":    true,
	"tracemalloc":  true,
	"tty":          true,
	"turtle":       true,
	"turtledemo":   true,
	"types":        true,
	"typing":       true,
	"unicodedata":  true,
	"unittest":     true,
	"urllib":       true,
	"uu":           true,
	"uuid":         true,
	"venv":         true,
	"warnings":     true,
	"wave":         true,
	"weakref":      true,
	"webbrowser":   true,
	"winreg":       true,
	"winsound":     true,
	"wsgiref":      true,
	"xdrlib":       true,
	"xml":          true,
	"xmlrpc":       true,
	"zipapp":       true,
	"zipfile":      true,
	"zipimport":    true,
	"zlib":         true,
	"zoneinfo":     true,
}
