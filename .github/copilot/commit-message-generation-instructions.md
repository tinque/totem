

# Commit Message Generation Instructions (Totem project)

Generate commit messages STRICTLY following the commitlint configuration for Conventional Commits format: `<type>(<scope>): <subject>`

## REQUIRED FORMAT:

```
type(scope): lowercase imperative verb followed by concise description
```

## ALLOWED TYPES (lowercase only):

- `feat`: new feature
- `fix`: bug fix
- `docs`: documentation changes
- `style`: formatting, whitespace, etc.
- `refactor`: code changes that neither fix bugs nor add features
- `perf`: performance improvements
- `test`: adding or correcting tests
- `build`: build system or external dependencies changes
- `ci`: CI configuration changes
- `chore`: other changes that don't modify src or test files
- `revert`: reverts a previous commit

## SCOPE (optional but recommended):

Use project/module names:
- `address`, `contact`, `parser`, `main`, `infra`, etc.
- Always lowercase, e.g. `(address)`, `(contact/email)`

## SUBJECT:

- Start with a lowercase imperative verb (e.g., "add", "fix", "update", "remove")
- Be concise and specific
- No punctuation at the end

## COMMITLINT VALIDATION RULES:

- [ ] **Header length**: 10-50 characters (including type, scope, and subject)
- [ ] **Type**: Must be one of the allowed types above
- [ ] **Format**: Exact format `type(scope): subject` or `type: subject`
- [ ] **Subject**: Imperative, lowercase, no period at end
- [ ] **Body lines**: Maximum 72 characters per line (if body is used)
- [ ] **Footer lines**: Maximum 72 characters per line (if footer is used)

## EXAMPLES FOR THIS PROJECT:

```
feat(address): add postal validation        # 33 chars
fix(contact): fix phone display             # 31 chars
refactor(parser): simplify csv logic        # 37 chars
docs: update commit instructions            # 32 chars
test(contact): add email tests              # 30 chars
chore: update dependencies                  # 26 chars
```

## IMPORTANT: CHARACTER COUNT

**CRITICAL**: The entire header (type + scope + subject) MUST be ≤ 50 characters.

Examples of headers that are TOO LONG:
- ❌ `docs: update commit message instructions` (42 chars - OK)
- ❌ `docs: update commit message generation instructions` (51 chars - TOO LONG!)

Examples of correct length:
- ✅ `docs: update commit instructions` (32 chars)
- ✅ `docs: update commit rules` (26 chars)

## BODY AND FOOTER (optional):

If additional context is needed, use body and footer with max 72 characters per line:

```
feat(contact): add phone validation

Implement comprehensive phone number validation
for international formats including country codes.

Closes #123
```
