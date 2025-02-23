# BRMS
Library to parse a specific syntax created to improve readability BRMS pronounced "Brems" stands for "Block Relation Mapping Syntax"

![BRMS Logo](assets/img/SDG_logo.png)

# BRMS Parser Documentation

## Introduction
The **BRMS Parser** is a command-line tool and library designed to parse and interpret files written in the **Block Relation Mapping Syntax (BRMS)**. It extracts and structures relationships between blocks and entities, and supports exclusions for fine-grained configuration of mappings.

---

## Features
- Parse BRMS files to extract:
  - Blocks: Mappings between source and destination blocks.
  - Entities: Mappings between source and destination entities.
  - Exclusions: Explicitly excluded blocks or entities.
- Dynamic path resolution for Windows and Unix-like systems.
- Robust error handling with informative log levels (`INFO`, `WARN`, `ERROR`).
- Multiplatform support for Windows and Unix.

---

## Installation

### Prerequisites
- **Go 1.20+**
- A valid BRMS configuration file (e.g., `subgroups.brms`).

### Clone the Repository
```bash
git clone https://github.com/Caezarr-OSS/brms-parser.git
cd brms-parser
```

### Build the Executable
#### Windows:
```bash
go build -o brms_parser.exe ./cmd/main.go
```

#### Unix:
```bash
go build -o brms_parser ./cmd/main.go
```

### Run the Program
```bash
./brms_parser
```

---

## Usage

### Command-Line Execution
1. Place your BRMS file in the appropriate directory (e.g., `config/examples/`).
2. Update the path in the `main.go` file or pass it dynamically.
3. Run the program:
   ```bash
   go run ./cmd/main.go
   ```

### Example Output
#### Input: `subgroups.brms`
```brms
[block_a/sub_block_1|block_b/sub_block_2]
entity_a|entity_b
entity_c|entity_d

[block_c|block_d]
sub_block_x/entity_e|sub_block_y/entity_f
```

#### Output:
```plaintext
Using file path: C:\path\to\config\examples\subgroups.brms
[INFO] Line 1: Parsed block 'block_a/sub_block_1|block_b/sub_block_2'
[INFO] Line 2: Parsed entity 'entity_a|entity_b'
[INFO] Line 3: Parsed entity 'entity_c|entity_d'
[INFO] Line 5: Parsed block 'block_c|block_d'
[INFO] Line 6: Parsed entity 'sub_block_x/entity_e|sub_block_y/entity_f'

Blocks:
  block_a/sub_block_1 -> block_b/sub_block_2
  block_c -> block_d

Entities:
  entity_a -> entity_b
  entity_c -> entity_d
  sub_block_x/entity_e -> sub_block_y/entity_f

Exclusions:
```

---

## Data Structure Explanation

### Blocks
Blocks represent mappings between source and destination block names. A **block** is defined in BRMS as:
```brms
[source_block:destination_block]
```
- **Source Block:** The name of the block in the source.
- **Destination Block:** The name of the block in the destination.
- **Example:**
  ```plaintext
  block_a/sub_block_1 -> block_b/sub_block_2
  ```
  This means `block_a/sub_block_1` in the source maps to `block_b/sub_block_2` in the destination.

### Entities
Entities are mappings within a block that specify the source and destination of individual items.
- An **entity** is defined under a block mapping as:
  ```brms
  source_entity:destination_entity
  ```
- **Source Entity:** The name of the item in the source.
- **Destination Entity:** The name of the item in the destination.
- **Example:**
  ```plaintext
  entity_a -> entity_b
  ```
  This means `entity_a` in the source maps to `entity_b` in the destination.

### Exclusions
Exclusions define blocks or entities that should be ignored during processing.
- A **block exclusion** is defined as:
  ```brms
  [source_block:]
  ```
- An **entity exclusion** is defined as:
  ```brms
  source_entity:
  ```
- **Example:**
  ```plaintext
  [block_a:]
  entity_a:
  ```
  This means `block_a` and `entity_a` will be excluded from processing.

---

## Logging and Error Handling
- **Log Levels:**
  - `INFO`: Displays parsing progress and details.
  - `WARN`: Highlights potential issues like excessive indentation.
  - `ERROR`: Reports critical parsing failures.

- **Error Handling:**
  - Invalid BRMS syntax is detected and logged with line numbers.
  - Missing files or inaccessible paths trigger descriptive errors.

---

## Best Practices
- Keep your BRMS files well-indented for better readability
- Use comments (#) to document complex mappings
- Avoid unnecessary spaces at the beginning and end of lines
- Organize your blocks in a logical and coherent manner
- Use descriptive names for your blocks and entities

## Configuration File Format
BRMS files use the `.brms` extension and follow a specific syntax:

- Lines starting with `#` are comments
- Blocks are defined between brackets `[]`
- Relations use the `|` character as a separator (configurable via `SetSeparator`)
- Hierarchical paths use `/` as a separator

### Example
```brms
[block_a/sub_block_1|block_b/sub_block_2]
entity_a|entity_b
entity_c|entity_d

[block_c|block_d]
sub_block_x/entity_e|sub_block_y/entity_f
```

### Custom Separator
You can change the default separator ("|") using the `SetSeparator` method:
```go
parser := brms.NewParser("config.brms", brms.LogLevelInfo)
parser.SetSeparator("->") // Change separator to ->
```

## Troubleshooting

### Common Errors
1. **Invalid Syntax**: Check that your blocks are properly defined with brackets
2. **Invalid Characters**: Avoid special characters in names
3. **Incorrect Indentation**: Remove spaces/tabs at the beginning of lines

### Error Messages
- `Error at line X: line outside of block section`: The line is not in a valid block
- `Error at line X: invalid format with multiple ':'`: Too many separators on one line

## Performance Considerations
- The parser loads the entire file into memory
- Parsing complexity is linear O(n) where n is the number of lines
- For large files, ensure you have sufficient memory available

## Integration Examples

### Library Usage
```go
import "brms_parser/brms"

parser := brms.NewParser("config.brms", brms.LogLevelInfo)
parsed, err := parser.Parse()
if err != nil {
    log.Fatal(err)
}
```

### Command Line Usage
```bash
# Parse with detailed logging
./brms_parser -log=info -file=config.brms

# Parse without logging
./brms_parser -log=error -file=config.brms
```

---

## Contributions
We welcome contributions to improve the BRMS Parser. Feel free to submit issues or pull requests on the repository.

---

## License
This project is licensed under the Apache 2.0. See the `LICENSE` file for details.
