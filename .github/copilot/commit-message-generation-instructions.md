Generate commit messages that STRICTLY follow the Conventional Commits format: <type>(<scope>): <subject>

## MANDATORY OUTPUT FORMAT:

**YOU MUST GENERATE EXACTLY THIS PATTERN:**

```
type(scope): lowercase imperative verb followed by description
```

**NEVER GENERATE:**

- Sentences starting with capital letters
- Messages without type(scope): prefix
- Past tense verbs (added, updated, fixed, removed)
- Periods at the end

## CRITICAL REQUIREMENTS:

### 1. HEADER FORMAT (MANDATORY):

- **EXACT FORMAT**: `<type>(<scope>): <subject>`
- **Maximum 100 characters total**
- **Must be all lowercase** (no capitals anywhere)
- **No period at the end**
- **Subject in imperative mood** (e.g., "add", "fix", "update", not "added", "fixed", "updated")

### 2. ALLOWED TYPES (lowercase only):

- `feat`: new feature
- `fix`: bug fix
- `docs`: documentation changes
- `style`: formatting, white-space, etc.
- `refactor`: code changes that neither fix bugs nor add features
- `perf`: performance improvements
- `test`: adding or correcting tests
- `build`: build system or external dependencies changes
- `ci`: CI configuration changes
- `chore`: other changes that don't modify src or test files
- `revert`: reverts a previous commit

### 3. SCOPE (optional but recommended):

- Use project/module names: `infra`, `api`, `webapp`, `core`, `shared`, etc.
- Lowercase only: `(infra)`, `(api)`, `(core/event-store)`

### 4. SUBJECT RULES:

- **Start with lowercase letter**
- **Imperative mood**: "replace env_name with resource_prefix", NOT "replaces" or "replaced"
- **Be concise and specific**
- **No punctuation at the end**

### 5. EXAMPLES OF CORRECT FORMAT:

```
feat(api): add user authentication endpoint
fix(infra): resolve terraform provider version conflicts
refactor(infra): replace env_name with resource_prefix and environment flags
docs(readme): update installation instructions
chore(deps): bump terraform aws provider to 6.12.0
```

**FOR YOUR CURRENT CHANGES, USE:**

```
refactor(infra): replace env_name with resource_prefix and environment flags
```

### 6. BODY (optional):

- Leave blank line after header
- Explain **what** and **why**, not how
- Max 100 characters per line
- Use imperative mood

### 7. FOOTER (optional):

- For breaking changes: `BREAKING CHANGE: <description>`
- For issue references: `closes #123`

## VALIDATION CHECKLIST:

- [ ] Starts with allowed type in lowercase
- [ ] Uses exact format `type(scope): subject`
- [ ] Subject starts with lowercase verb in imperative mood
- [ ] No period at end of subject
- [ ] Total header ≤ 100 characters
- [ ] No capitals except in breaking change footer

## COMMON MISTAKES TO AVOID:

❌ **WRONG**:

```
Refactor infrastructure variables and resource naming conventions
Update Terraform configuration to introduce resource prefix
Added new boolean variables for environment detection
```

✅ **CORRECT**:

```
refactor(infra): replace env_name with resource_prefix and environment flags
fix(terraform): resolve provider version conflicts
feat(infra): add environment detection boolean variables
```

## ENFORCEMENT RULES:

**CRITICAL**: You MUST generate EXACTLY in this format:

1. **NEVER start with a capital letter**
2. **ALWAYS include type and colon**: `type(scope): `
3. **ALWAYS use imperative mood in present tense**
4. **NEVER use past tense verbs** (renamed, introduced, updated, adjusted, removed)
5. **USE present tense imperatives**: replace, introduce, update, adjust, remove

**OUTPUT FORMAT**: Generate ONLY the header `type(scope): subject` unless body/footer is specifically requested.

**VALIDATION**: Every commit message MUST pass this regex: `^(feat|fix|docs|style|refactor|perf|test|build|ci|chore|revert)(\([a-z0-9/-]+\))?: [a-z][^.]*[^.]$`

## FINAL CHECK:

Before generating any commit message, verify:

1. ✅ Starts with lowercase type from allowed list
2. ✅ Has format `type(scope): `
3. ✅ Subject starts with lowercase imperative verb
4. ✅ No period at end
5. ✅ Under 100 characters total

**EXAMPLE VALIDATION:**

- ❌ `Refactor infrastructure variables` → Wrong (capital R, no type)
- ✅ `refactor(infra): replace env_name with resource_prefix` → Correct
