#include "node.h"

Node::Node()
{

}

int Node::getValue() const
{
    return value;
}

void Node::setValue(int newValue)
{
    value = value;
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
