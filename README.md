# Lem-in
![](https://img.freepik.com/free-photo/3d-cartoon-wild-animal-nature_23-2151746842.jpg?ga=GA1.1.1553843546.1721645701&semt=ais_hybrid)

A Go implementation of an ant farm pathfinding simulator that finds the optimal paths for ants to traverse from start to end room while avoiding congestion.

## Overview

Lem-in reads a farm description from a file and calculates the most efficient way to move ants from the start room to the end room. It uses the Edmonds-Karp algorithm for finding multiple paths and implements an optimal ant distribution strategy.

## Features

- Parses and validates ant farm descriptions from input files
- Finds multiple valid paths using Edmonds-Karp algorithm
- Optimizes ant distribution across available paths
- Handles various error cases and invalid inputs
- Provides detailed move-by-move output

## Installation

### Clone the repository
```bash
git clone https://learn.zone01kisumu.ke/git/oumouma/lem-in.git
```
### Usage

```
cd lem-in
cd cmd
go run main.go farm.txt
```

## Input File Format

The input file should follow this format:
```
number_of_ants
##start
start_room x y
room1 x y
room2 x y
##end
end_room x y
room1-room2
start_room-room1
```


Example:
```
4
##start
0 0 0
1 1 1
2 2 2
##end
3 3 3
0-1
1-2
2-3
```

## Output Format
The program outputs:

The input file content

A blank line

The ant movements in the format: Lx-y where x is the ant number and y is the destination room

Example output:
```
4
##start
0 0 0
1 1 1
2 2 2
##end
3 3 3
0-1
1-2
2-3

L1-1
L1-2 L2-1
L1-3 L2-2 L3-1
L2-3 L3-2 L4-1
L3-3 L4-2
L4-3
```
## Implementation Details

This document outlines the key components, algorithms, error handling, validation rules, and performance considerations for the Ant Farm simulation project.

### Key Components

#### Room Structure
- **Stores Room Information**: Each room contains attributes such as name and coordinates.
- **Tracks Start/End Status**: Indicates whether a room is a starting or ending point for ants.
- **Maintains Connections**: Keeps track of connections to other rooms (tunnels).

#### AntFarm Structure
- **Manages the Entire Colony**: Oversees all rooms and their relationships within the ant colony.
- **Stores Rooms and Relationships**: Maintains a list of rooms and how they are interconnected.
- **Handles Path Finding and Ant Movement Simulation**: Responsible for simulating ant movements through the colony based on available paths.

### Algorithms

### Edmonds-Karp Algorithm: 
Utilized for finding multiple paths between rooms, ensuring efficient movement of ants through the colony.
```mermaid
graph TD
    A[Start Edmonds-Karp] --> B[Create Residual Graph]
    B --> C[Find Augmenting Path using BFS]
    C --> D{Path Found?}
    D -->|Yes| E[Update Residual Graph]
    E --> C
    D -->|No| F[End: All Paths Found]
    
```

### Breadth-First Search (BFS):
 Employed for path finding to identify the shortest routes between rooms.
- #### Algorithm
```mermaid
graph TD
    A[Start] --> B[Initialize:<br/>visited map<br/>parent map<br/>queue]
    B --> C[Take first room<br/>from queue]
    C --> D{Queue empty?}
    D -->|No| E[Get next unvisited<br/>connected room]
    E --> F{Room has<br/>capacity > 0?}
    F -->|Yes| G[Mark room as visited<br/>Set parent<br/>Add to queue]
    G --> H{Is it end room?}
    H -->|Yes| I[Construct and<br/>return path]
    H -->|No| C
    F -->|No| E
    D -->|Yes| J[Return empty path]

    style A fill:#f9f,stroke:#333,stroke-width:4px
    style I fill:#9f9,stroke:#333,stroke-width:4px
    style J fill:#f99,stroke:#333,stroke-width:4px
```
- #### Room exploration using BFS
```mermaid
graph LR
    Start((Start))
    A((A))
    B((B))
    C((C))
    D((D))
    End((End))
    
    Start --> A
    Start --> B
    A --> C
    B --> C
    B --> D
    C --> End
    D --> End

    style Start fill:#f96,stroke:#333,stroke-width:4px
    style End fill:#9f6,stroke:#333,stroke-width:4px
    
    classDef level1 fill:#ffb366
    classDef level2 fill:#99ff99
    classDef level3 fill:#ff99cc
    
    class A,B level1
    class C,D level2
    class End level3
```
### Optimal Ant Distribution:
Implements strategies to distribute ants optimally across available paths to maximize efficiency.
- #### Main simulation function(simulateAnts)
```mermaid
graph TD
    A[Start SimulateAnts] --> B[Sort paths by length]
    B --> C[Calculate paths info]
    C --> D[Find optimal turns & distribution]
    D --> E[Generate moves]
    E --> F[Return moves sequence]
```

- #### Finding Optimal Distribution(findOptimalTurns)
```mermaid
graph LR
    A[Binary Search] --> B[Try mid turns]
    B --> C{Can all ants finish?}
    C -->|Yes| D[Store distribution<br>Try fewer turns]
    C -->|No| E[Try more turns]
    D --> A
    E --> A
```

- #### Move Generation(generateMoves)
```mermaid
graph TD
    A[Start turn] --> B[Move existing ants]
    B --> C[Start new ants]
    C --> D[Format moves]
    D --> E{More turns?}
    E -->|Yes| A
    E -->|No| F[End]
```

- #### Collision Avoidance
```mermaid
graph TD
    A[Check room] --> B{Is room occupied?}
    B -->|No| C[Move ant]
    B -->|Yes| D[Wait]
    C --> E[Mark room occupied]
```



### Error Handling

The program is designed to handle various error cases effectively:

- **Invalid Number of Ants**: Checks if the specified number of ants is valid.
- **Missing Start/End Rooms**: Validates that both start and end rooms are defined.
- **Invalid Room Names**: Ensures room names adhere to specified rules.
- **Duplicate Rooms/Links**: Prevents the creation of duplicate rooms or tunnels between the same rooms.
- **Invalid Coordinates**: Verifies that coordinates are integers.
- **Invalid File Format**: Checks for correctness in file input formats.

### Validation Rules

To maintain the integrity of the simulation, the following validation rules are enforced:

1. **Room Names**:
   - Cannot start with 'L' or '#'.
   - Cannot contain spaces.

2. **Tunnel Connections**:
   - Each tunnel must connect exactly two rooms.
   - No duplicate tunnels between the same pair of rooms.

3. **Ant Placement**:
   - Only one ant is allowed per room, except in start and end rooms.

4. **Coordinates**:
   - All coordinates must be integers.

### Performance

The implementation focuses on efficiency through the following strategies:

- **Edmonds-Karp Algorithm**: Efficiently finds multiple paths for ant movement, optimizing flow through the colony.
- **Binary Search**: Utilized for optimal turn calculation, enhancing performance in pathfinding scenarios.
- **Map-Based Data Structures**: Implemented for quick lookups of room connections and attributes, improving overall access times.

### Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
