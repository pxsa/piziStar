#include "pizistar.h"

PiziStar::PiziStar(int row, int col)
{
    land = new Land(row, col);
}

PiziStar::~PiziStar()
{
    clearList(this->pathStorage);
    clearList(this->openList);
}

void PiziStar::clearList(vector<Node *> list)
{
    while (!list.empty()) {
        delete list.front();
        list.pop_back();
    }
    list.clear();
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
