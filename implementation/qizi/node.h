#ifndef NODE_H
#define NODE_H

#include <string>

using namespace std;

class Node
{
public:
    Node();

    int getValue() const;
    void setValue(int newValue);

    int getCostToGoal() const;
    void setCostToGoal(int value);

    string getTag() const;
    void setTag(const string &value);

    Node *getNext() const;
    void setNext(Node *value);

private:
    int value;
    int costToGoal;
    string tag;
    Node* next;
};

#endif // NODE_H
