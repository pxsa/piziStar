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
