// Core types for xschema adapters
export interface ConvertInput {
  namespace: string;
  id: string;
  schema: object;
}

export interface ConvertResult {
  namespace: string;
  id: string;
  imports: string[];
  schema?: string;
  type?: string;
  validate?: string;
  validateImports?: string[];
}
