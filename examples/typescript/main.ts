import { createXSchemaClient } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";
import { schemas } from "./.xschema/xschema.gen";

export const xschema = createXSchemaClient(schemas);

// Register schemas - these calls are parsed by CLI to know what to generate
const User = xschema.fromURL("User", "https://api.example.com/schemas/user.json", zodAdapter);
const Post = xschema.fromFile("Post", "./schemas/post.json", zodAdapter);

// Use the schemas - full Zod API works
const userData = User.parse({
  id: "123e4567-e89b-12d3-a456-426614174000",
  name: "Alice",
  email: "alice@example.com",
});

// Type inference works
type UserType = typeof User._type;
//   ^? { id: string; name: string; email: string; age?: number }

console.log("Parsed user:", userData);

// Validation errors work
try {
  User.parse({ id: "not-a-uuid", name: "", email: "invalid" });
} catch (e) {
  console.log("Validation failed as expected");
}
