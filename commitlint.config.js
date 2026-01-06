export default {
	extends: ["@commitlint/config-conventional"],
	rules: {
		"scope-enum": [
			2,
			"always",
			[
				"cli",     // Go CLI
				"ts",      // TypeScript package (@xschema)
				"py",      // Python package (future)
				"deps",    // Dependency updates
				"release", // Release commits
				"web", // Web / Docs / Server
				"ci" // CI
			],
		],
		"scope-empty": [0],
	},
};
