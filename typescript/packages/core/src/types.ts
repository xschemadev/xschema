// Core types for xschema adapters
export interface ConvertInput {
  namespace: string;
  id: string;
  varName: string;
  schema: object;
}

export interface ConvertResult {
  namespace: string;
  id: string;
  varName: string;
  imports: string[];
  schema?: string;
  type?: string;
  validate?: string;
  validationImports?: string[];
}
