#ifndef PIZISTAR_H
#define PIZISTAR_H

#include "node.h"
#include "land.h"
#include <vector>

using namespace std;

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

#endif // PIZISTAR_H
