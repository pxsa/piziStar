#include "land.h"

Land::Land(int rows, int cols)
{
    this->rows = rows;
    this->cols = cols;

    nodes = vector<vector<Node*>>(rows);
    for (int i=0; i<rows; i++) {
        nodes[i] = vector<Node*>(cols);
        for (int j=0; j<cols; j++) {
            nodes[i][j] = createNode(i, j);
        }
    }

    // bad code
    for (int i=0; i<rows; i++) {
        for (int j=0; j<cols; j++) {
            detectNeighbors(nodes[i][j]);
        }
    }
}

void Land::configureObstacle(Node* node, bool status)
{
    node->setIsObstacle(status);
}

Node *Land::createNode(int i, int j)
{
    Node* node = new Node();

    node->setRow(i);
    node->setCol(j);
    node->setIsObstacle(false);
    node->setTag("new");
    node->setNext(nullptr);
    node->setCostToGoal(-1);
    node->setKey(0);

    return node;
}

void Land::detectNeighbors(Node *node)
{
    vector<Node*> neighbors;

    // left neighbor
    if (node->getRow() - 1 >= 0) {
        neighbors.push_back(nodes[node->getRow()-1][node->getCol()]);
    }

    // right neighbor
    if (node->getRow() + 1 < rows ){
        neighbors.push_back(nodes[node->getRow()+1][node->getCol()]);
    }

    // upper neighbor
    if (node->getCol() - 1 >= 0) {
        neighbors.push_back(nodes[node->getRow()][node->getCol()-1]);
    }

    // lower neighbor
    if (node->getCol() + 1 < cols) {
        neighbors.push_back(nodes[node->getRow()][node->getCol()+1]);
    }

    node->setNeighbors(neighbors);
}

int Land::getRows() const
{
    return rows;
}

void Land::setRows(int newRows)
{
    rows = newRows;
}

int Land::getCols() const
{
    return cols;
}

void Land::setCols(int newCols)
{
    cols = newCols;
}

vector<vector<Node *> > Land::getNodes() const
{
    return nodes;
}

void Land::setNodes(const vector<vector<Node *> > &newNodes)
{
    nodes = newNodes;
}
