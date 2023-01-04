#ifndef LAND_H
#define LAND_H

#include "node.h"
#include <vector>

using namespace std;

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

#endif // LAND_H
