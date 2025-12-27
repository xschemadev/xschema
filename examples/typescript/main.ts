import { createXSchemaClient } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";
// After `xschema generate`, CLI injects: import { schemas } from "./.xschema/xschema.gen";
// and adds `schemas` to the config below

export const xschema = createXSchemaClient({
  outputDir: ".xschema",
  // schemas, // <- injected by CLI after generation
});

// Register schemas - these calls are parsed by CLI to know what to generate
const User = xschema.fromURL("User", "https://api.example.com/schemas/user.json", zodAdapter);
const Post = xschema.fromFile("Post", "./schemas/post.json", zodAdapter);

// Before generation: User is PleaseRunXSchemaGenerate
// After generation: User is z.ZodObject<...>

// Use the schemas - full Zod API works (after generation)
// const userData = User.parse({
//   id: "123e4567-e89b-12d3-a456-426614174000",
//   name: "Alice",
//   email: "alice@example.com",
// });

// Type inference works (after generation)
// type UserType = typeof User._type;

console.log("User schema:", User);
console.log("Post schema:", Post);
