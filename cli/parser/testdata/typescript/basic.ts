// Basic xschema calls - all should be parsed
const zodAdapter = { name: "zod", __brand: "xschema-adapter" } as const;

// fromURL with double quotes
xschema.fromURL("User", "https://api.example.com/user.json", zodAdapter);

// fromFile with double quotes
xschema.fromFile("Post", "./schemas/post.json", zodAdapter);
