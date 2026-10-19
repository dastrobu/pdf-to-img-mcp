# Contributing to pdf2img

Thanks for your interest in contributing! This document explains how to do it.

## Code of Conduct

- Be respectful toward other contributors
- Use inclusive language
- Accept constructive criticism
- Focus on the code, not on the people

## How to contribute

### Reporting bugs

1. Check that the bug is not already reported
2. Create an issue with:
    - A clear description of the problem
    - Steps to reproduce
    - Expected vs actual behavior
    - Operating system and Go version
    - A minimal reproducible example

```
**Description:**
The converter fails with PDFs that contain JPG images.

**Steps to reproduce:**
1. Download document.pdf
2. Run: pdf2img -i document.pdf -o ./output
3. See the error

**Error:**
```
failed to render page: invalid image format
```

**System:**
- OS: Windows 11
- Go: 1.21
```

### Suggesting improvements

1. Open an issue with the `enhancement` label
2. Describe the improvement in detail
3. Explain the use case
4. Suggest a possible solution

### Submitting a Pull Request

#### Preparation

```bash
# Fork the repository
git clone https://github.com/scopweb/mcp-go-pdf-to-img.git
cd pdf2img

# Create a branch for your feature
git checkout -b feature/tu-feature-name
```

#### Development

1. **Keep the code clean**
    ```bash
    # Format the code
    make fmt

    # Lint
    make lint
    ```

2. **Write tests**
    ```bash
    # If you added new functions
    go test -v ./...

    # Check coverage
    go test -cover ./...
    ```

3. **Meaningful commits**
    ```bash
    # Clear and atomic commits
    git commit -m "Add feature: description"

    # Avoid
    git commit -m "Fix stuff"
    ```

#### Submit the PR

```bash
# Push to your fork
git push origin feature/tu-feature-name

# Create a Pull Request on GitHub
# - Clear title
# - Detailed description
# - Reference to related issues
# - Screenshots if it is UI
```

## Structure of a good PR

```markdown
## Description
What your PR does (2-3 lines)

## Type of change
- [ ] Bug fix
- [ ] New feature
- [ ] Existing improvement
- [ ] Documentation
- [ ] Refactoring

## Changes made
- Change 1
- Change 2
- Change 3

## Testing
How to test:
1. Step 1
2. Step 2

## Checklist
- [ ] Code formatted (make fmt)
- [ ] Tests passing (make test)
- [ ] Lint passing (make lint)
- [ ] Documentation updated
- [ ] Atomic and clear commits
```

## Style guidelines

### Go

```go
// ✓ Good
func convertPDF(path string) error {
    if err := validatePath(path); err != nil {
        return fmt.Errorf("validation failed: %w", err)
     }
     // ...
    return nil
}

// ✗ Avoid
func ConvertPDF(path string) (err error) {
    err = validatePath(path)
    if err != nil {
        panic(err)
     }
     // ...
}
```

**Rules:**
- Names in English
- 80 characters max per line (flexible)
- Explicit error handling
- Defer for cleanup
- Document public functions

### Commit Messages

```
Format:
<type>(<scope>): <subject>

<body>

<footer>

Examples:
feat(converter): add support for encrypted PDFs
fix(cli): correct DPI calculation
docs(readme): update examples
refactor(mcp): simplify server initialization
test(converter): add edge case tests

Types:
- feat: New feature
- fix: Bug fix
- docs: Documentation changes
- style: Formatting changes (no logic)
- refactor: Code refactoring
- test: Add or update tests
- chore: Dependency, tooling, etc. changes
```

### Documentation

- README: Public, clear, and concise
- Comments: Explain the "why", not the "what"
- Names: Self-explanatory
- Examples: Functional and up-to-date code

```go
// ✓ Good
// renderPage renders a PDF page to an image.
// If DPI is 0, defaults to 150.
func (c *Converter) renderPage(page int, dpi float64) error {
     // ...
}

// ✗ Avoid
// render page
func render(p int, d float64) error {
     // ...
}
```

## Review process

1. We check:
    - Clean and well-structured code
    - Tests passing
    - Documentation updated
    - Meaningful commits

2. We may request:
    - Changes to the code
    - More tests
    - Documentation updates
    - Refactoring

3. Once approved:
    - Automatic merge
    - Thanks in CHANGELOG (upcoming releases)

## Local development

### Setup

```bash
git clone https://github.com/scopweb/mcp-go-pdf-to-img.git
cd pdf2img
go mod download
```

### Build

```bash
make build           # Build the CLI
make build-dev       # With debug symbols
make test           # Run tests
make fmt            # Format the code
make lint           # Lint
make clean          # Clean
```

### Debug

```bash
# With logs
./pdf2img -i document.pdf -o ./output -v

# Profiling
go test -cpuprofile=cpu.prof ./...
go tool pprof cpu.prof
```

## Areas where help is appreciated

### High impact
- Performance improvements
- New output formats
- Better error handling
- Documentation

### Medium impact
- Additional tests
- Improved examples
- Refactoring
- CI/CD improvements

### Low priority
- Cosmetic changes
- Comments without value
- Name changes without context

## Frequently asked questions

**Can I work on a feature that is on the roadmap?**
Open an issue first to coordinate.

**How long does a review take?**
Usually 3-5 days, at most 2 weeks.

**My PR was rejected. What do I do?**
Read the comments and learn from them. You can adjust and resubmit.

**Do I need to sign anything?**
No. Just keep the Apache 2.0 license.

## Acknowledgments

We value every contribution, big or small. If your PR is merged, you will appear in:
- Commit history
- CONTRIBUTORS.md (coming soon)
- Release notes

---

**Questions?** Open an issue or contact the maintainers.

Thanks for contributing to pdf2img!