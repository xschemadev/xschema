import { xschema } from "./xschema";

// Mock adapter (normally from @xschema/adapter-zod)
const zodAdapter = { name: "zod", __brand: "xschema-adapter" } as const;

// From URL
xschema.fromURL("User", "https://cdn.my/user.json", zodAdapter);
console.log("User:", xschema.User);

// From file
xschema.fromFile("Post", "./schemas/post.json", zodAdapter);
console.log("Post:", xschema.Post);

// If haven't run the CLI, here is the DX
xschema.fromURL("Unknown", "https://cdn.my/unknown.json", zodAdapter);
