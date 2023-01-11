# piziStar

## Content
- [Abstract](#Abstract)
- [Algorithm](#Algorithm)
  - [Land](#Land)
  - [Node](#Node)
  - [PiziStar](#PiziStar)
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
  
## Land
```c++
class Land
{
public:
    Land(int rows, int cols);

    void configureObstacle(Node *node, bool status);
    Node* createNode(int i, int j);
    void detectNeighbors(Node* node);

    int getRows() const;
    void setRows(int newRows);

    int getCols() const;
    void setCols(int newCols);

    vector<vector<Node*> > getNodes() const;
    void setNodes(const vector<vector<Node*> > &newNodes);

private:
    int rows;
    int cols;
    vector<vector<Node*>> nodes;
};
```
## Node
```c++
class Node
{
public:
    Node();

    int getRow() const;
    void setRow(int newRow);

    int getCol() const;
    void setCol(int newCol);

    string getTag() const;
    void setTag(const string &value);

    int getCostToGoal() const;
    void setCostToGoal(int value);

    Node *getNext() const;
    void setNext(Node *value);

    int getKey() const;
    void setKey(int newKey);

    bool getIsObstacle() const;
    void setIsObstacle(bool newIsObstacle);

    vector<Node *> getNeighbors() const;
    void setNeighbors(const vector<Node *> &newNeighbors);

    void appendNeighbor(Node* node);

private:
    int row;
    int col;
    string tag;
    int costToGoal;
    int key;
    bool isObstacle;
    Node* next;
    vector<Node*> neighbors;
};
```
## PiziStar
```c++
class PiziStar
{
public:
    PiziStar(int row, int col);
    ~PiziStar();

    void init();
    vector<Node*> reconstruction();
    int process_state();
    int modify_state();
    Node* min_state();
    int get_kmin();
    void deleteNode(Node* node);
    void insert(Node* node, int h_new);

    void clearList(vector<Node*> list);

    Node *getStart() const;
    void setStart(Node *newStart);

    Node *getGoal() const;
    void setGoal(Node *newGoal);

    vector<Node *> getOpenList() const;
    void setOpenList(const vector<Node *> &newOpenList);

    vector<Node *> getPathStorage() const;
    void setPathStorage(const vector<Node *> &newPathStorage);

private:
    vector<Node*> openList;
    vector<Node*> pathStorage;
    Land* land  = nullptr;
    Node* start = nullptr;
    Node* goal  = nullptr;
};
```
# summary
