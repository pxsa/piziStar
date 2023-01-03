#ifndef NODE_H
#define NODE_H

#include <string>
#include <vector>

using namespace std;

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

#endif // NODE_H
