// Basic xschema calls - all should be parsed
import { createXSchemaClient } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";

const xschema = createXSchemaClient({});

// fromURL with double quotes
xschema.fromURL("User", "https://api.example.com/user.json", zodAdapter);

// fromFile with double quotes
xschema.fromFile("Post", "./schemas/post.json", zodAdapter);
