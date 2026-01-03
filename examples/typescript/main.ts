import { schemas } from "./.xschema/xschema.gen";
import { createXSchemaClient, XSchemaType } from "@xschema/client";
import { zodAdapter } from "@xschema/zod";

export const xschema = createXSchemaClient({ schemas })

// Dual purpose: declaration (parsed by CLI) + returns schema for immediate use
const appleAppSite = xschema.fromURL("AppleAppSiteAssociation", "https://www.schemastore.org/apple-app-site-association.json", zodAdapter)

const calendarJsonSchema = xschema.fromFile("Calendar", "calendar.json", zodAdapter)

// Type extraction using XSchemaType helper
export type CalendarType = XSchemaType<"Calendar">

// Access schemas by ID (convenient when you don't want to repeat URL/path)
const calendar = xschema.getFromId("Calendar")
const appleAppSiteById = xschema.getFromId("AppleAppSiteAssociation")

const someData = {}

appleAppSite.parse(someData)
calendar.parse(someData)
