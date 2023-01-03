#include "node.h"

Node::Node()
{

}

int Node::getCostToGoal() const
{
    return costToGoal;
}

void Node::setCostToGoal(int value)
{
    costToGoal = value;
}

string Node::getTag() const
{
    return tag;
}

void Node::setTag(const string &value)
{
    tag = value;
}

Node *Node::getNext() const
{
    return next;
}

void Node::setNext(Node *value)
{
    next = value;
}

int Node::getRow() const
{
    return row;
}

void Node::setRow(int newRow)
{
    row = newRow;
}

int Node::getCol() const
{
    return col;
}

void Node::setCol(int newCol)
{
    col = newCol;
}

int Node::getKey() const
{
    return key;
}

void Node::setKey(int newKey)
{
    key = newKey;
}

bool Node::getIsObstacle() const
{
    return isObstacle;
}

void Node::setIsObstacle(bool newIsObstacle)
{
    isObstacle = newIsObstacle;
}

vector<Node *> Node::getNeighbors() const
{
    return neighbors;
}

void Node::setNeighbors(const vector<Node *> &newNeighbors)
{
    neighbors = newNeighbors;
}

void Node::appendNeighbor(Node *node)
{
    neighbors.push_back(node);
}
