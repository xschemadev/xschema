import { Schema as S } from "effect";

// {{GENERATED_CODE}}
const schema = S.Unknown;

const testCases: Array<{ id: string; data: unknown }> = []; // {{TEST_CASES_STRING}}

for (const tc of testCases) {
	const result = S.decodeUnknownEither(schema)(tc.data);
	console.log(
		JSON.stringify({
			id: tc.id,
			valid: result._tag === "Right",
		})
	);
}
