# piziStar

## Content
- [Abstract](#Abstract)
- [Algorithm](#Algorithm)
- [Implementation](#Implementation)
- [Summary](#summary)

# Abstract
Path planning lets an autonomous vehicle or a robot find the shortest and most obstacle-free path from a start to goal state. The path can be a set of states (position and orientation) or waypoints. Path planning requires a map of the environment along with start and goal states as input.
# Algorithm
The algorithm works by iteratively selecting a node from the OPEN list and evaluating it. It then propagates the node's changes to all of the neighboring nodes and places them on the OPEN list. This propagation process is termed "expansion". 

# Implementation
for implementation we used golang and c++ programming languages,in this `README.md` I'll explain c++ code.
the code consists of the following classes:
  - land
  - node
  - pizistar
  
# summary
