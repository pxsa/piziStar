#ifndef PIZISTAR_H
#define PIZISTAR_H

#include "node.h"
#include <list>

using namespace std;

class PiziStar
{
public:
    PiziStar();
    ~PiziStar();

    void clearList(list<Node*> list);

private:
    list<Node*> pathStorage;
    list<Node*> openList;
    list<Node*> insertedList;
    //
    list<Node*> openNodes;
    list<Node*> closedNodes;

    list<Node*> nodes;
    //
    Node start;
    Node goal;
    int numberOfNode;
};

#endif // PIZISTAR_H
