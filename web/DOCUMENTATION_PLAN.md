# xschema Documentation Plan

## Vision & Problem Statement

xschema brings **Cross Language Type Safety** to the JSON Schema ecosystem.

### The Problem (Why xschema Exists)

JSON Schema is already the standard for cross-language validation, BUT the tooling around it is a **jungle** (and we're the gorilla that tames it! 🦍). Current problems:

1. **Fragmented Ecosystem**: Each library does its own thing with no consistency
2. **Hidden Limitations**: Libraries don't document which JSON Schema features they support
3. **Poor DX**: You have to build your own tooling, commands, and workflows
4. **Non-Compliance**: Most famous libraries don't even follow the JSON Schema spec properly
5. **Complex Features Fail**: As soon as you use refs, remote refs, circular refs, etc., they break
6. **No Trust**: Without compliance testing, you're never sure if a library will work for your schema

xschema solves this by:
- Generating native validators (Zod, Pydantic, etc.) from JSON Schema
- Running comprehensive compliance tests (98%+ coverage)
- Providing excellent DX with CLI, clients, and libraries
- Transparently documenting what works and what doesn't

## Documentation Structure

Following the Fumadocs pattern of organizing by **products/interaction points**:

### Main Sections (Products)

```
web/content/docs/
├── index.mdx                    # Ecosystem overview & navigation
├── introduction.mdx             # Cross-language type safety vision
│
├── cli/                         # Product 1: Command-line tool
│   ├── index.mdx               # Overview: What is the CLI
│   ├── quickstart.mdx          # 5-minute getting started
│   ├── installation.mdx        
│   ├── configuration.mdx       # Config file format
│   ├── commands/
│   │   ├── generate.mdx        # Main command
│   │   ├── compliance.mdx      # Test adapters
│   │   └── watch.mdx           # Dev mode
│   └── guides/
│       ├── ci-cd.mdx           # GitHub Actions, etc.
│       ├── monorepos.mdx       # Monorepo setups
│       └── headers.mdx         # Auth headers
│
├── client/                      # Product 2: Runtime libraries
│   ├── index.mdx               # Overview: What are clients
│   ├── quickstart.mdx          # Dynamic schema loading
│   ├── typescript.mdx          # @xschemadev/client
│   ├── python.mdx              # xschema-client [FUTURE]
│   └── guides/
│       ├── dynamic-schemas.mdx # Loading from APIs
│       └── caching.mdx         # Performance
│
├── libraries/                   # Product 3: Programmatic APIs
│   ├── index.mdx               # Overview: Build tool integration
│   ├── quickstart.mdx          # Using xschema in your tools
│   ├── node.mdx                # Node.js API
│   ├── go.mdx                  # Go library
│   └── guides/
│       ├── webpack.mdx         # Webpack plugin [FUTURE]
│       ├── vite.mdx            # Vite plugin [FUTURE]
│       └── custom-tooling.mdx  # Build your own
│
├── adapters/                    # Reference: Available adapters
│   ├── index.mdx               # Adapter ecosystem
│   ├── typescript/             # [EXISTING]
│   │   ├── index.mdx          
│   │   ├── zod/
│   │   │   ├── index.mdx      # Features, examples
│   │   │   └── compliance.mdx # Auto-generated report
│   │   ├── valibot/...
│   │   └── ...
│   └── python/                 # [FUTURE]
│       └── pydantic/...
│
├── compliance/                  # Trust: Proof it works
│   ├── index.mdx               # What is compliance testing
│   └── unsupported-features.mdx# Static generation limits
│
├── schemas/                     # Knowledge: JSON Schema guide
│   ├── index.mdx               # JSON Schema with xschema
│   ├── basics.mdx              # Schema fundamentals
│   ├── sources.mdx             # File, URL, inline
│   ├── refs.mdx                # $ref, remote refs
│   └── advanced.mdx            # Complex patterns
│
└── recipes/                     # Cookbook: Real-world patterns
    ├── form-validation.mdx     
    ├── api-validation.mdx      
    ├── openapi.mdx             
    └── migration.mdx           # From other tools
```

## Key Pages Content Strategy

### 1. Landing Page (`index.mdx`)

```mdx
---
title: xschema - Cross Language Type Safety
description: Tame the JSON Schema jungle with native validators
---

# Cross Language Type Safety

JSON Schema everywhere. Native validators for every language.

<QuickLinks>
  <Link to="/cli/quickstart" icon="Terminal">
    Start with CLI →
  </Link>
  <Link to="/introduction" icon="Book">
    Why xschema? →
  </Link>
</QuickLinks>

## The xschema Ecosystem

<Cards>
  <Card title="CLI" href="/docs/cli">
    Generate validators at build time
  </Card>
  <Card title="Client" href="/docs/client">
    Load schemas dynamically at runtime
  </Card>
  <Card title="Libraries" href="/docs/libraries">
    Integrate into your build tools
  </Card>
</Cards>

## Trusted & Tested

<ComplianceOverview />
<!-- Shows 98%+ compliance across adapters -->
```

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

## Navigation Structure

In Fumadocs, this would appear as:

```
Main Navigation (dropdown/select):
- CLI
- Client  
- Libraries

Secondary Navigation (always visible):
- Adapters
- Compliance
- Schemas
- Recipes
```

## Writing Guidelines

1. **Lead with value**: Every page should explain WHY before HOW
2. **Show compliance**: Always link to compliance reports when mentioning adapters
3. **Be transparent**: Document what doesn't work and why
4. **Examples first**: Show code before explaining
5. **Cross-language**: Always show TypeScript + Python examples (when available)

## Implementation Phases

### Phase 1: Core Structure
- Set up main product sections (CLI, Client, Libraries)
- Create landing and introduction pages
- Move existing compliance content

### Phase 2: CLI Documentation
- Quickstart guide
- Command reference
- Configuration guide
- Common workflows

### Phase 3: Supporting Content
- Adapter guides with examples
- JSON Schema education
- Recipes for common use cases

### Phase 4: Future Products
- Python client documentation
- Build tool plugins
- Advanced library usage

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