export interface XSchemaAdapter {
  readonly __brand: "xschema-adapter";
  readonly name: string;
  readonly language: string;
}

export interface ConvertInput {
  namespace: string;
  id: string;
  schema: object;
}

export interface ConvertResult {
  namespace: string;
  id: string;
  imports: string[];
  schema: string;
  type: string;
}
