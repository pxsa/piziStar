#include "pizistar.h"

PiziStar::PiziStar()
{
    for (int i=0; i<numberOfNode; i++) {
        Node* node = new Node();
        node->setCostToGoal(0);
        node->setTag("new");
        node->setNext(nullptr);

        nodes.push_front(node);
    }
}

PiziStar::~PiziStar()
{
    clearList(this->insertedList);
    clearList(this->pathStorage);
    clearList(this->closedNodes);
    clearList(this->openNodes);
    clearList(this->openList);
    clearList(this->nodes);
}

void PiziStar::clearList(list<Node *> list)
{
    while (!list.empty()) {
        delete list.front();
        list.pop_front();
    }
    list.clear();
}

list<Node *> PiziStar::getOpenList() const
{
    return openList;
}

void PiziStar::setOpenList(const list<Node *> &newOpenList)
{
    openList = newOpenList;
}

list<Node *> PiziStar::getPathStorage() const
{
    return pathStorage;
}

void PiziStar::setPathStorage(const list<Node *> &newPathStorage)
{
    pathStorage = newPathStorage;
}

Node *PiziStar::getStart() const
{
    return start;
}

void PiziStar::setStart(Node *newStart)
{
    start = newStart;
}

Node *PiziStar::getGoal() const
{
    return goal;
}

void PiziStar::setGoal(Node *newGoal)
{
    goal = newGoal;
}

vector<Node *> PiziStar::getOpenList() const
{
    return openList;
}

void PiziStar::setOpenList(const vector<Node *> &newOpenList)
{
    openList = newOpenList;
}

vector<Node *> PiziStar::getPathStorage() const
{
    return pathStorage;
}

void PiziStar::setPathStorage(const vector<Node *> &newPathStorage)
{
    pathStorage = newPathStorage;
}
