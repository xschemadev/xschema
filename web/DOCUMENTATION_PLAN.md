# xschema Documentation Plan

## Vision & Problem Statement

xschema brings **Cross Language Type Safety** to the JSON Schema ecosystem.

### The Problem (Why xschema Exists)

JSON Schema is already the standard for cross-language validation, BUT the tooling around it is a **jungle** (favicon). Current problems:

1. **Fragmented Ecosystem**: Each library does its own thing with no consistency
2. **Hidden Limitations**: Libraries don't document which JSON Schema features they support
3. **Poor DX**: You have to build your own tooling, commands, and workflows
4. **Non-Compliance**: Most famous libraries don't even follow the JSON Schema spec properly
5. **Complex Features Fail**: As soon as you use refs, remote refs, circular refs, etc., they break
6. **No Trust**: Without compliance testing, you're never sure if a library will work for your schema

xschema solves this by:

- Generating native validators (Zod, Pydantic, etc.) from JSON Schema
- Running comprehensive compliance tests
- Providing excellent DX with CLI, clients, and libraries
- Transparently documenting what works and what doesn't

### 2. Introduction (`introduction.mdx`)

Tell the story of WHY xschema exists:

- The JSON Schema jungle problem
- How xschema tames it (gorilla reference!)
- Cross-language type safety vision
- What makes xschema different

### 3. Product Quickstarts

Each product section has its own quickstart:

**CLI Quickstart**: Install → Config → Generate → Use  
**Client Quickstart**: Install → Load schema → Validate  
**Libraries Quickstart**: Install → Integrate → Generate

### 4. Compliance Section

This is KEY for trust. Show:

- Live compliance scores
- What each percentage means
- Link to detailed reports per adapter
- Transparency about limitations

## Success Metrics

Good documentation will:

1. Make it obvious which xschema product to use
2. Get users to working code in <5 minutes
3. Build trust through transparency (compliance)
4. Reduce support questions
5. Showcase the "jungle taming" aspect

## Key Differentiators to Emphasize

1. **Compliance-First**: We test everything against official test suite
2. **Multi-Language**: True cross-language type safety
3. **Production-Ready**: Static generation, no runtime overhead
4. **Transparent**: We document exactly what works and what doesn't
5. **Modern DX**: Watch mode, great errors, TypeScript-first

This positions xschema as the mature, reliable solution in a chaotic ecosystem.
