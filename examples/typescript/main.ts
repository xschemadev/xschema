import { createXSchemaClient } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";
import { schemas } from "./.xschema/xschema.gen";
import z from "zod";

export const xschema = createXSchemaClient({ schemas })

const appleAppSite = xschema.fromURL("AppleAppSiteAssociation", "https://www.schemastore.org/apple-app-site-association.json", zodAdapter)

const opencode = xschema.fromURL("OpenCode", "https://opencode.ai/config.json", zodAdapter)

const calendarJsonSchema = xschema.fromFile("Calendar", "calendar.json", zodAdapter)

export type CalendarType = z.infer<typeof calendarJsonSchema>

const someData = {}

appleAppSite.parse(someData)

console.log(opencode)
