# Commands Package
This package provides the store of commands to send.

# Contents
1. [Command Type](#command)
2. [Repository Type](#repository)

<a name="command"></a>

## Command Type
Contains the attributes:
- **robotID**: *int*
- **linearVelocity**: *float64*
- **angularVelocity**: *float64*
And the exported getters and setters to data handling.

<a name="repository"></a>

## Repository Type
Contains the attribute:
- **Repository**: *map[string]Command*
And the exported getter and setter to data handlig.