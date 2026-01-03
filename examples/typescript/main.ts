import { schemas } from "./.xschema/xschema.gen";
import { createXSchemaClient, XSchemaType } from "@xschema/client";

export const xschema = createXSchemaClient({ schemas, defaultNamespace: "another" })

// Type extraction using XSchemaType helper
export type AnotherType = XSchemaType<"another:TestUrl">


const anotherUrl = xschema("another:TestUrl")
const userUrl = xschema("user:TestUrl")
//same as anotherUrl
const anotherUrl = xschema("TestUrl")
