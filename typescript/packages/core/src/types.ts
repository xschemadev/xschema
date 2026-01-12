// Core types for xschema adapters
export interface ConvertInput {
  namespace: string;
  id: string;
  varName: string;
  schema: object;
  vocabulary?: Record<string, boolean>; // $vocabulary from custom metaschema (undefined = all enabled)
}

export interface ConvertResult {
  namespace: string;
  id: string;
  varName: string;
  imports: string[];
  schema?: string;
  type?: string;
  validate?: string;
  validateImports?: string[];
}
