// Edge cases - whitespace, comments, multiline
const adapter = { name: "zod" };

// Extra whitespace
xschema.fromURL(   "ExtraSpaces"   ,   "https://example.com/spaces.json"   ,   adapter   );

// Multiline call
xschema.fromURL(
  "Multiline",
  "https://example.com/multiline.json",
  adapter
);

// Very long URL
xschema.fromURL("LongURL", "https://very-long-domain-name.example.com/api/v1/schemas/users/definitions/extended-profile.json", adapter);

// Unicode in name
xschema.fromURL("Schéma", "https://example.com/unicode.json", adapter);

// NOTE: Comments between args DON'T work - parser skips them
// xschema.fromFile(
//   "WithComments", // schema name
//   "./schemas/commented.json", // file path
//   adapter // // adapter
// );

